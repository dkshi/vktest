package repository

import (
	"github.com/dkshi/vktest"
	"github.com/jmoiron/sqlx"
)

type Actors interface {
	InsertActor(actor *vktest.Actor) (int, error)
	UpdateActor(actor *map[string]any) error
	DeleteActor(id int) error
	DeleteActorFilms(actorID int) error
	GetActors() (*[]vktest.Actor, error)
}

type Films interface {
	InsertFilm(film *vktest.Film) (int, error)
	UpdateFilm(film *map[string]any) error
	DeleteFilm(filmID int) error
	GetFilms(attrs *map[string]any) (*[]vktest.Film, error)
	InsertFilmActors(filmID int, actors []vktest.Actor) error
	DeleteFilmActors(filmID int) error
}

type Authorization interface {
	CreateAdmin(admin *vktest.Admin) (int, error)
	GetAdmin(adminname, password string) (*vktest.Admin, error)
}

type Repository struct {
	Actors
	Films
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Actors:        NewActorsPostgres(db),
		Films:         NewFilmsPostgres(db),
		Authorization: NewAuthPostgres(db),
	}
}
