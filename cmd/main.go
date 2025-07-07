package main

import (
	"log"
	"os"

	"github.com/bhushan-aruto/justarquest/internal/handler"
	"github.com/bhushan-aruto/justarquest/internal/infrastructure/database"
	"github.com/bhushan-aruto/justarquest/internal/usecase"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	databaseName := os.Getenv("DATABASE_NAME")

	if databaseURL == "" || databaseName == "" {
		panic("DATABASE_URL or DATABASE_NAME is not set")
	}

	dbConn := database.NewDatabase(databaseURL)
	repo := dbConn

	userUsecase := usecase.NewUserCreate(repo)
	userHandler := handler.NewUserHandler(userUsecase)

	e := echo.New()
	e.POST("/register", userHandler.RegisterUser)
	e.Logger.Fatal(e.Start(":8080"))
}
