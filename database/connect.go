package database

import (
	"log"
	"os"

	"github.com/TaRun1369/GoBlog/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")

	}
	dsn := os.Getenv("DSN")                                     // dsn from env file
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // open connection
	// &gorm.config defines the config for gorm connection
	// config means the way we want to connect to the database
	if err != nil {
		panic("could not connect to db")
	} else {
		log.Println("connected successfully to db")
	}
	DB = database

	database.AutoMigrate(
		&models.User(),
	)
}
