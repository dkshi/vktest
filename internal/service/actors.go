package service

import (
	"github.com/dkshi/vktest"
	"github.com/dkshi/vktest/internal/repository"
)

type ActorsService struct {
	repo *repository.Repository
}

func NewActorsService(repo *repository.Repository) *ActorsService {
	return &ActorsService{
		repo: repo,
	}
}

func (s *ActorsService) GetActors() (*[]vktest.Actor, error) {
	return s.repo.GetActors()
}

func (s *ActorsService) DeleteActor(actorID int) error {
	return s.repo.DeleteActor(actorID)
}

func (s *ActorsService) InsertActor(actor *vktest.Actor) (int, error) {
	return s.repo.InsertActor(actor)
}

func (s *ActorsService) UpdateActor(actor *map[string]any) error {
	return s.repo.UpdateActor(actor)
}
