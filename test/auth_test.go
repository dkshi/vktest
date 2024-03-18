package test

import (
	"testing"

	"github.com/dkshi/vktest"
	"github.com/dkshi/vktest/internal/repository"
)

func TestAuthPostgres_CreateAdmin(t *testing.T) {
	setup()
	defer teardown()

	authRepo := repository.NewAuthPostgres(testDB)

	admin := &vktest.Admin{
		Adminname: "admin",
		Password:  "password123",
	}

	id, err := authRepo.CreateAdmin(admin)
	if err != nil {
		t.Fatalf("error creating admin: %v", err)
	}

	if id == 0 {
		t.Fatalf("expected non-zero admin ID, got %d", id)
	}
}

func TestAuthPostgres_GetAdmin(t *testing.T) {
	setup()
	defer teardown()

	authRepo := repository.NewAuthPostgres(testDB)

	adminname := "admin"
	password := "password123"

	admin, err := authRepo.GetAdmin(adminname, password)
	if err != nil {
		t.Fatalf("error getting admin: %v", err)
	}

	if admin.Adminname != adminname {
		t.Fatalf("expected admin name %s, got %s", adminname, admin.Adminname)
	}

	if admin.Password != password {
		t.Fatalf("expected admin password %s, got %s", password, admin.Password)
	}
}
