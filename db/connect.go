package database

import (
	"os"

	"super_calendar/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	dbURL := "postgres://" + user + ":" + password + "@localhost:" + port + "/" + dbName

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database!")
	}

	DB = db

	db.AutoMigrate(&models.CalendarEvent{})
}
