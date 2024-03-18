package repository

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dkshi/vktest"
	"github.com/dkshi/vktest/scripts"
	"github.com/jmoiron/sqlx"
)

type ActorsPostgres struct {
	db *sqlx.DB
}

func NewActorsPostgres(db *sqlx.DB) *ActorsPostgres {
	return &ActorsPostgres{db: db}
}

func (a *ActorsPostgres) InsertActor(actor *vktest.Actor) (int, error) {
	var id int
	dateOfBirth, err := scripts.StringToDate(actor.DateOfBirth)
	if err != nil {
		return 0, err
	}
	res := a.db.QueryRow("INSERT INTO actors (name, gender, birth_date) VALUES ($1, $2, $3) RETURNING actor_id;",
		actor.Name, actor.Gender, dateOfBirth)
	err = res.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (a *ActorsPostgres) DeleteActor(actorID int) error {
	err := a.DeleteActorFilms(actorID)
	if err != nil {
		return err
	}
	deleteQuery := `DELETE FROM actors WHERE actor_id=$1;`
	_, err = a.db.Exec(deleteQuery, actorID)
	if err != nil {
		return err
	}
	return nil
}

func (a *ActorsPostgres) GetActors() (*[]vktest.Actor, error) {
	actors := make([]vktest.Actor, 0)
	selectQuery := `SELECT * FROM actors`
	res, err := a.db.Query(selectQuery)
	if err != nil {
		return &[]vktest.Actor{}, err
	}
	for res.Next() {
		var newActor vktest.Actor
		err = res.Scan(&newActor.ActorID, &newActor.Name, &newActor.Gender, &newActor.DateOfBirth)
		if err != nil {
			return &[]vktest.Actor{}, err
		}
		films := make([]vktest.Film, 0)
		filmsSelectQuery := `SELECT f.*
			FROM films f
			JOIN actor_film af ON f.film_id = af.film_id
			JOIN actors a ON af.actor_id = a.actor_id
			WHERE a.actor_id = $1;`
		resFilms, err := a.db.Query(filmsSelectQuery, newActor.ActorID)
		if err != nil {
			return &[]vktest.Actor{}, err
		}
		for resFilms.Next() {
			var newFilm vktest.Film
			err = resFilms.Scan(&newFilm.FilmID, &newFilm.Title, &newFilm.Description, &newFilm.ReleaseDate, &newFilm.Rating)
			if err != nil {
				return &[]vktest.Actor{}, err
			}
			films = append(films, newFilm)
		}
		newActor.Films = films
		actors = append(actors, newActor)
	}
	return &actors, nil
}

func (a *ActorsPostgres) UpdateActor(actor *map[string]any) error {
	updateQueryB := strings.Builder{}
	updateQueryB.WriteString("UPDATE actors SET ")
	params := make([]any, 0, len(*actor))

	for k, v := range *actor {
		if k == "birth_date" {
			vStr, ok := v.(string)
			if !ok {
				return fmt.Errorf("incorrect type of date")
			}
			_, err := scripts.StringToDate(vStr)
			if err != nil {
				return err
			}
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
	params = append(params, (*actor)["actor_id"])
	updateQueryB.WriteString(" WHERE actor_id = $" + strconv.Itoa(len(params)) + ";")

	_, err := a.db.Exec(updateQueryB.String(), params...)
	if err != nil {
		return err
	}
	return nil
}

func (a *ActorsPostgres) DeleteActorFilms(actorID int) error {
	deleteQuery := `DELETE FROM actor_film WHERE actor_id=$1`
	_, err := a.db.Exec(deleteQuery, actorID)
	if err != nil {
		return err
	}
	return nil
}
