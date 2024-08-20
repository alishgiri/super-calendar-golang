package controllers

import (
	"strconv"
	"strings"
	database "super_calendar/db"
	"super_calendar/models"
	"super_calendar/util"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetCalEvents(c *fiber.Ctx) error {
	userEmail := c.Query("email")
	var calEvents []models.CalendarEvent
	if userEmail != "" {
		if err := database.DB.Where("email like ?", userEmail).Find(&calEvents).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
		}
	}
	return c.JSON(calEvents)
}

func GetCalEventWithId(c *fiber.Ctx) error {
	calEventId, err := strconv.Atoi(c.Params("id", ""))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid event id.",
		})
	}

	var calEvent models.CalendarEvent
	if err := database.DB.Where("id =?", calEventId).First(&calEvent).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Calendar event with id not found.",
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.JSON(calEvent)
}

func AddCalEvent(c *fiber.Ctx) error {
	var payload models.CalendarEvent
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	validationErr := util.Validate.Struct(payload)
	if validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": validationErr.Error()})
	}

	payload.Email = strings.ToLower(payload.Email)
	if payload.EndDate == nil {
		payload.EndDate = &payload.StartDate
	}

	if err := database.DB.Create(&payload).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	return c.JSON(fiber.Map{"inserted_id": payload.Id})
}

func UpdateCalEvent(c *fiber.Ctx) error {
	calEventId, err := strconv.Atoi(c.Params("id", ""))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid event id.",
		})
	}

	var payload models.CalendarEvent
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	validationErr := util.Validate.Struct(payload)
	if validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": validationErr.Error()})
	}

	result := database.DB.Where("id =?", calEventId).Updates(&payload)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{"success": true})
}

func DeleteCalEvent(c *fiber.Ctx) error {
	calEventId, err := strconv.Atoi(c.Params("id", ""))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid event id.",
		})
	}

	var calEvent models.CalendarEvent
	result := database.DB.Where("id =?", calEventId).Delete(&calEvent)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{"success": true})
}
