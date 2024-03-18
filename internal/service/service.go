package service

import (
	"github.com/dkshi/vktest"
	"github.com/dkshi/vktest/internal/repository"
)

type Actors interface {
	DeleteActor(actorID int) error
	InsertActor(actor *vktest.Actor) (int, error)
	UpdateActor(actor *map[string]any) error
	GetActors() (*[]vktest.Actor, error)
}

type Films interface {
	InsertFilm(film *vktest.Film) (int, error)
	UpdateFilm(film *map[string]any) error
	DeleteFilm(filmID int) error
	GetFilms(attrs *map[string]any) (*[]vktest.Film, error)
}

type Authorization interface {
	CreateAdmin(user *vktest.Admin) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
	Films
	Actors
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		Films: NewFilmsService(repo),
		Actors: NewActorsService(repo),
	}
}
