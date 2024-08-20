package controllers

import (
	"fmt"
	database "super_calendar/db"
	"super_calendar/models"
	"super_calendar/services"
	"time"
)

func SendEventNotificationToUserEmail() error {
	calEvents, err := getCalEventsForNextOneMinute()
	if err != nil {
		return err
	}
	if len(calEvents) == 0 {
		return nil
	}
	for _, ce := range calEvents {
		err := services.SendEmail(ce.Email, "REMINDER: "+ce.Title, "<h1>Calendar Event Reminder</h1>")
		if err != nil {
			// Record Error Somewhere
			fmt.Println(err)
		} else {
			RecordUserNotified(ce.Email)
		}
	}
	return nil
}

func RecordUserNotified(email string) error {
	updates := map[string]interface{}{
		"notified": true,
	}
	result := database.DB.Where("email LIKE", email).Updates(updates)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func getCalEventsForNextOneMinute() ([]models.CalendarEvent, error) {
	var calEvents []models.CalendarEvent
	startTime := time.Now().UTC()
	endTime := startTime.Add(time.Minute * 1)
	if err := database.DB.Where("start_date BETWEEN ? AND ? AND all_day = ?", startTime, endTime, false).Find(&calEvents).Error; err != nil {
		return nil, err
	}
	return calEvents, nil
}
