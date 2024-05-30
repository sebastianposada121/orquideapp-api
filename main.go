package main

import (
	"log"
	"orquideapp/src/infrastructure/database"
	"orquideapp/src/interfaces/controllers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error load env: %v", err)
	}

	db, err := database.OrquideaDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Iniciar Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	controllers.InitControllers(e, db)

	// Iniciar el servidor
	e.Logger.Fatal(e.Start(":5000"))
}
