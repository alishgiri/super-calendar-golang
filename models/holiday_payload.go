package models

type HolidayPayload struct {
	Country string `json:"country" validate:"required,min=2,max=10"`
	Year    string `json:"year" validate:"required,min=4,max=4"`
}
