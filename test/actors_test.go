package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/dkshi/vktest"
	"github.com/dkshi/vktest/internal/repository"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	testDBDriver   = "postgres"
	testDBUser     = "postgres"
	testDBPassword = "qwerty"
	testDBName     = "postgres"
)

var testDB *sqlx.DB

func setup() {
	db, err := sqlx.Open(testDBDriver, fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", testDBUser, testDBPassword, testDBName))
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}

	testDB = db
}

func teardown() {
	if testDB != nil {
		testDB.Close()
	}
}

func TestActorsPostgres_InsertActor(t *testing.T) {
	setup()
	defer teardown()

	actorRepo := repository.NewActorsPostgres(testDB)

	actor := &vktest.Actor{
		Name:        "Test Actor",
		Gender:      "Male",
		DateOfBirth: "1990-01-01",
	}

	id, err := actorRepo.InsertActor(actor)
	if err != nil {
		t.Fatalf("error inserting actor: %v", err)
	}

	if id == 0 {
		t.Fatalf("expected non-zero actor ID, got %d", id)
	}
}

func TestActorsPostgres_DeleteActor(t *testing.T) {
	setup()
	defer teardown()

	actorRepo := repository.NewActorsPostgres(testDB)

	actor := &vktest.Actor{
		Name:        "Test Actor",
		Gender:      "Male",
		DateOfBirth: "1990-01-01",
	}

	id, err := actorRepo.InsertActor(actor)
	if err != nil {
		t.Fatalf("error inserting actor: %v", err)
	}

	err = actorRepo.DeleteActor(id)
	if err != nil {
		t.Fatalf("error deleting actor: %v", err)
	}
}

func TestActorsPostgres_GetActors(t *testing.T) {
	setup()
	defer teardown()

	actorRepo := repository.NewActorsPostgres(testDB)

	actors, err := actorRepo.GetActors()
	if err != nil {
		t.Fatalf("error getting actors: %v", err)
	}

	if actors == nil {
		t.Fatalf("expected non-nil actors slice")
	}
}

func TestActorsPostgres_UpdateActor(t *testing.T) {
	setup()
	defer teardown()

	actorRepo := repository.NewActorsPostgres(testDB)

	actor := &vktest.Actor{
		Name:        "Test Actor",
		Gender:      "Male",
		DateOfBirth: "1990-01-01",
	}

	id, err := actorRepo.InsertActor(actor)
	if err != nil {
		t.Fatalf("error inserting actor: %v", err)
	}

	newName := "Updated Actor Name"
	updateData := map[string]interface{}{
		"name":     newName,
		"actor_id": id,
	}

	err = actorRepo.UpdateActor(&updateData)
	if err != nil {
		t.Fatalf("error updating actor: %v", err)
	}

	updatedActors, err := actorRepo.GetActors()
	if err != nil {
		t.Fatalf("error getting actors: %v", err)
	}

	for _, a := range *updatedActors {
		if a.ActorID == id && a.Name != newName {
			t.Fatalf("expected actor name to be %s, got %s", newName, a.Name)
		}
	}
}
