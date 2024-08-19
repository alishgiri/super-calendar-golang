package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"super_calendar/models"
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

	holidays, err := fetchHolidays(payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	return c.JSON(fiber.Map{"holidays": holidays})
}

func fetchHolidays(payload models.HolidayPayload) ([]map[string]interface{}, error) {
	HOLIDAY_API_KEY := os.Getenv("HOLIDAYS_API_KEY")
	url := fmt.Sprintf("https://api.api-ninjas.com/v1/holidays?country=%s&year=%s&type=major_holiday", payload.Country, payload.Year)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Api-Key", HOLIDAY_API_KEY)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var holidays []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&holidays); err != nil {
		return nil, err
	}

	return holidays, nil
}
