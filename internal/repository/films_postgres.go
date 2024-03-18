package repository

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dkshi/vktest"
	"github.com/dkshi/vktest/scripts"
	"github.com/jmoiron/sqlx"
)

type FilmsPostgres struct {
	db *sqlx.DB
}

func NewFilmsPostgres(db *sqlx.DB) *FilmsPostgres {
	return &FilmsPostgres{db: db}
}

func (f *FilmsPostgres) InsertFilm(film *vktest.Film) (int, error) {
	insertQuery := `INSERT INTO films (title, description, release_date, rating) VALUES ($1, $2, $3, $4) RETURNING film_id;`
	releaseDate, err := scripts.StringToDate(film.ReleaseDate)
	if err != nil {
		return 0, err
	}
	res := f.db.QueryRow(insertQuery, film.Title, film.Description, releaseDate, film.Rating)
	var id int
	err = res.Scan(&id)
	if err != nil {
		return 0, err
	}
	err = f.InsertFilmActors(id, film.Actors)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (f *FilmsPostgres) GetFilms(attrs *map[string]any) (*[]vktest.Film, error) {
	films := make([]vktest.Film, 0)
	selectQueryB := strings.Builder{}
	selectQueryB.WriteString(`SELECT DISTINCT f.*
		FROM films f
		LEFT JOIN actor_film af ON f.film_id = af.film_id
		LEFT JOIN actors a ON af.actor_id = a.actor_id `)

	conds := false
	if title, exists := (*attrs)["title"]; exists {
		strTitle, ok := title.(string)
		if !ok {
			return &[]vktest.Film{}, fmt.Errorf("incorrect type of title fragment")
		}
		selectQueryB.WriteString("WHERE f.title ILIKE '%" + strTitle + "%' ")
		conds = true
	}

	if _, exists := (*attrs)["actor"]; exists {
		if actorName, ok := (*attrs)["actor"].(string); ok {
			if conds {
				selectQueryB.WriteString("AND ")
			} else {
				selectQueryB.WriteString("WHERE ")
			}
			selectQueryB.WriteString("a.name ILIKE '%" + actorName + "%' ")
		} else {
			return &[]vktest.Film{}, fmt.Errorf("incorrect type of actor name fragment")
		}
	}
	selectQueryB.WriteString("ORDER BY ")
	lenOrder := 0
	if order, exists := (*attrs)["order"]; exists {
		orderSlice, ok := order.([]string)
		if !ok {
			return &[]vktest.Film{}, fmt.Errorf("incorrect order format")
		}
		lenOrder = len(orderSlice)
		for _, param := range orderSlice {
			valid := regexp.MustCompile("^[A-Za-z0-9_]+$")
			if !valid.MatchString(param) {
				return &[]vktest.Film{}, fmt.Errorf("invalid order attribute, might be sql injection")
			}
			selectQueryB.WriteString("f." + param + ", ")
		}
	}
	if lenOrder == 0 {
		selectQueryB.WriteString("f.rating, ")
	}

	selectQuery := selectQueryB.String()[:selectQueryB.Len()-2]
	selectQueryB.Reset()
	selectQueryB.WriteString(selectQuery)

	if asc, exists := (*attrs)["ascending"]; exists {
		boolAsc, ok := asc.(bool)
		if !ok {
			return &[]vktest.Film{}, fmt.Errorf("incorrect type of ascending attribute")
		}
		if boolAsc {
			selectQueryB.WriteString(" ASC;")
		} else {
			selectQueryB.WriteString(" DESC;")
		}
	} else {
		selectQueryB.WriteString(" DESC;")
	}

	res, err := f.db.Query(selectQueryB.String())
	if err != nil {
		return &[]vktest.Film{}, err
	}
	defer res.Close()

	for res.Next() {
		var newFilm vktest.Film
		err = res.Scan(&newFilm.FilmID, &newFilm.Title, &newFilm.Description, &newFilm.ReleaseDate, &newFilm.Rating)
		if err != nil {
			return &[]vktest.Film{}, err
		}
		actors := make([]vktest.Actor, 0)
		actorsSelectQuery := `SELECT DISTINCT a.*
			FROM actors a
			JOIN actor_film af ON a.actor_id = af.actor_id
			JOIN films f ON af.film_id = f.film_id
			WHERE f.film_id = $1;`
		resActors, err := f.db.Query(actorsSelectQuery, newFilm.FilmID)
		if err != nil {
			return &[]vktest.Film{}, err
		}
		for resActors.Next() {
			var newActor vktest.Actor
			err = resActors.Scan(&newActor.ActorID, &newActor.Name, &newActor.Gender, &newActor.DateOfBirth)
			if err != nil {
				return &[]vktest.Film{}, err
			}
			actors = append(actors, newActor)
		}
		newFilm.Actors = actors
		films = append(films, newFilm)
	}
	return &films, nil
}

func (f *FilmsPostgres) UpdateFilm(film *map[string]any) error {
	updateQueryB := strings.Builder{}
	updateQueryB.WriteString("UPDATE films SET ")
	params := make([]any, 0, len(*film))

	for k, v := range *film {
		if k == "release_date" {
			vStr, ok := v.(string)
			if !ok {
				return fmt.Errorf("incorrect type of release_date")
			}
			_, err := scripts.StringToDate(vStr)
			if err != nil {
				return err
			}
		}

		if k == "actors" {
			filmID, ok := (*film)["film_id"].(int)
			if !ok {
				return fmt.Errorf("incorrect type of film_id")
			}
			f.DeleteFilmActors(filmID)
			actorsAny, ok := v.([]any)
			if !ok {
				return fmt.Errorf("incorrect type of films list")
			}
			actors := make([]vktest.Actor, 0, len(actorsAny))
			for _, actor := range v.([]any) {
				actorMap, ok := actor.(map[string]any)
				if !ok {
					return fmt.Errorf("incorrect type of actors")
				}
				actorIDFloat64, ok := actorMap["actor_id"].(float64)
				if !ok {
					return fmt.Errorf("incorrect type of actor_id")
				}
				actorID, err := scripts.СonvertFloat64ToInt(actorIDFloat64)
				if err != nil {
					return err
				}
				newActor := vktest.Actor{ActorID: actorID}
				actors = append(actors, newActor)
			}
			f.InsertFilmActors(filmID, actors)
			continue
		}

		params = append(params, v)
		updateQueryB.WriteString(k + "=$" + strconv.Itoa(len(params)) + ", ")
	}

	if len(params) == 0 {
		return nil
	}

	updateQuery := updateQueryB.String()[:updateQueryB.Len()-2]
	updateQueryB.Reset()
	updateQueryB.WriteString(updateQuery)
	params = append(params, (*film)["film_id"])
	updateQueryB.WriteString(" WHERE film_id = $" + strconv.Itoa(len(params)) + ";")

	_, err := f.db.Exec(updateQueryB.String(), params...)
	if err != nil {
		return err
	}
	return nil
}

func (f *FilmsPostgres) DeleteFilm(filmID int) error {
	err := f.DeleteFilmActors(filmID)
	if err != nil {
		return err
	}
	deleteQuery := `DELETE FROM films WHERE film_id=$1;`
	_, err = f.db.Exec(deleteQuery, filmID)
	if err != nil {
		return err
	}
	return nil
}

func (f *FilmsPostgres) InsertFilmActors(filmID int, actors []vktest.Actor) error {
	insertQuery := `INSERT INTO actor_film (actor_id, film_id) VALUES ($1, $2);`

	for _, actor := range actors {
		_, err := f.db.Exec(insertQuery, actor.ActorID, filmID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *FilmsPostgres) DeleteFilmActors(filmID int) error {
	deleteQuery := `DELETE FROM actor_film WHERE film_id=$1`
	_, err := f.db.Exec(deleteQuery, filmID)
	if err != nil {
		return err
	}
	return nil
}
