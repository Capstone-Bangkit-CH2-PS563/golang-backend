package main

import (
	"capstone/config"
	"capstone/middleware"
	"capstone/routes"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env")
	}

	db := config.InitDB()
	e := echo.New()
	routes.Routes(e, db)
	middleware.Logmiddleware(e)

	e.Logger.Fatal(e.Start(":8000"))

}
