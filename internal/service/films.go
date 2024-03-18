package service

import (
	"github.com/dkshi/vktest"
	"github.com/dkshi/vktest/internal/repository"
)

type FilmsService struct {
	repo *repository.Repository
}

func NewFilmsService(repo *repository.Repository) *FilmsService {
	return &FilmsService{
		repo: repo,
	}
}

func (s *FilmsService) InsertFilm(film *vktest.Film) (int, error) {
	return s.repo.InsertFilm(film)
}

func (s *FilmsService) UpdateFilm(film *map[string]any) error {
	return s.repo.UpdateFilm(film)
}

func (s *FilmsService) DeleteFilm(filmID int) error {
	return s.repo.DeleteFilm(filmID)
}

func (s *FilmsService) GetFilms(attrs *map[string]any) (*[]vktest.Film, error) {
	return s.repo.GetFilms(attrs)
}
