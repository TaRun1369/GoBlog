package main

import (
	"log"
	"os"

	"github.com/TaRun1369/GoBlog/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect() // connect to database
	err := godotenv.Load() // loads the env file
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT") // get port from env
	app := fiber.New() // using fiber framework which is similar to express, fiber used for routing , routing means which url will go to which function
	app.Listen(":"+port) // Listen function is used to listen to the port
}
