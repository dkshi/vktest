package test

import (
	"testing"

	"github.com/dkshi/vktest"
	"github.com/dkshi/vktest/internal/repository"
	_ "github.com/lib/pq"
)

func TestFilmsPostgres_InsertFilm(t *testing.T) {
	setup()
	defer teardown()

	filmRepo := repository.NewFilmsPostgres(testDB)

	film := &vktest.Film{
		Title:       "Test Film",
		Description: "Test Description",
		ReleaseDate: "2023-01-01",
		Rating:      8.5,
		Actors:      []vktest.Actor{},
	}

	id, err := filmRepo.InsertFilm(film)
	if err != nil {
		t.Fatalf("error inserting film: %v", err)
	}

	if id == 0 {
		t.Fatalf("expected non-zero film ID, got %d", id)
	}
}

func TestFilmsPostgres_GetFilms(t *testing.T) {
	setup()
	defer teardown()

	filmRepo := repository.NewFilmsPostgres(testDB)

	attrs := map[string]interface{}{
		"title": "Test",
	}

	films, err := filmRepo.GetFilms(&attrs)
	if err != nil {
		t.Fatalf("error getting films: %v", err)
	}

	if films == nil {
		t.Fatalf("expected non-nil films slice")
	}
}

func TestFilmsPostgres_UpdateFilm(t *testing.T) {
	setup()
	defer teardown()

	filmRepo := repository.NewFilmsPostgres(testDB)

	film := map[string]interface{}{
		"film_id":      1,
		"title":        "Updated Film Title",
		"release_date": "2024-01-01",
	}

	err := filmRepo.UpdateFilm(&film)
	if err != nil {
		t.Fatalf("error updating film: %v", err)
	}
}

func TestFilmsPostgres_DeleteFilm(t *testing.T) {
	setup()
	defer teardown()

	filmRepo := repository.NewFilmsPostgres(testDB)

	err := filmRepo.DeleteFilm(1)
	if err != nil {
		t.Fatalf("error deleting film: %v", err)
	}
}
