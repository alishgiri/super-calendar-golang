package initialize

import (
	"log"

	database "super_calendar/db"
	"super_calendar/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Database(envPath string) {
	err := godotenv.Load(envPath)

	if err != nil {
		log.Fatal(".env load failed", err)
	}

	database.Connect()
}

func App() *fiber.App {
	app := fiber.New()

	routes.SetupRoutes(app)

	return app
}
