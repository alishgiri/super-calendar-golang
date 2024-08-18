package routes

import (
	"super_calendar/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/calendar-events", controllers.GetCalEvents)
	app.Get("/api/calendar-events/:id", controllers.GetCalEventWithId)
	app.Post("/api/calendar-event", controllers.AddCalEvent)
	app.Put("/api/calendar-event/:id", controllers.UpdateCalEvent)
	app.Delete("/api/calendar-event/:id", controllers.DeleteCalEvent)
}
