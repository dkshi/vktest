package main

import (
	"os"

	"github.com/dkshi/vktest/internal/handler"
	"github.com/dkshi/vktest/internal/repository"
	"github.com/dkshi/vktest/internal/service"
	"github.com/sirupsen/logrus"
)

// @title REST API for VK
// @version 1.0
// @description API Server for VK Intern test task

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		DBName:   os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	if err != nil {
		logrus.Fatalf("error connecting database: %s", err)
	}
	repo := repository.NewRepository(db)
	s := service.NewService(repo)

	h := handler.NewHandler("8080", s)
	h.InitRoutes()
}
