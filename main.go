package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TaRun1369/GoBlog/database"
	"github.com/TaRun1369/GoBlog/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()     // connect to database
	err := godotenv.Load() // loads the env file
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	defer fmt.Println("hello")
	port := os.Getenv("PORT") // get port from env
	app := fiber.New()        // using fiber framework which is similar to express, fiber used for routing , routing means which url will go to which function
	app.Use(cors.New(cors.Config{

		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000",
	}))

	routes.Setup(app)
	app.Listen(":" + port) // Listen function is used to listen to the port
}
