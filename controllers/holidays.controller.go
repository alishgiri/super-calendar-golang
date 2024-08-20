package controllers

import (
	"super_calendar/models"
	"super_calendar/services"
	"super_calendar/util"

	"github.com/gofiber/fiber/v2"
)

func GetHolidays(c *fiber.Ctx) error {
	var payload models.HolidayPayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	validationErr := util.Validate.Struct(payload)
	if validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": validationErr})
	}

	holidays, err := services.FetchHolidays(payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	return c.JSON(fiber.Map{"holidays": holidays})
}
