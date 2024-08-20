package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"super_calendar/models"
)

func FetchHolidays(payload models.HolidayPayload) ([]map[string]interface{}, error) {
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
