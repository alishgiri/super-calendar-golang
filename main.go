package main

import (
	"log"
	"os"
	"super_calendar/controllers"
	database "super_calendar/db"
	"super_calendar/initialize"
	"super_calendar/services"
	"super_calendar/util"

	"github.com/go-co-op/gocron/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env load failed", err)
	}

	database.Connect()
	util.InitializePayloadValidator()

	app := initialize.App()
	services.ScheduleCronJob(gocron.NewTask(controllers.SendEventNotificationToUserEmail))

	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
