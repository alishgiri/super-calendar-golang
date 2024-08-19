package initialize

import (
	"log"

	database "super_calendar/db"
	"super_calendar/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:5173",
	}))

	routes.SetupRoutes(app)

	return app
}
