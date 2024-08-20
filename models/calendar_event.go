package models

import (
	"time"
)

type CalendarEvent struct {
	Id         uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title      string     `json:"title" gorm:"not null" validate:"required,min=3,max=100"`
	Color      string     `json:"color" gorm:"not null" validate:"required,min=3,max=7"`
	Email      string     `json:"email" gorm:"not null" validate:"required,email"`
	StartDate  time.Time  `json:"start" gorm:"not null" validate:"required"`
	EndDate    *time.Time `json:"end" gorm:"default:null" validate:"omitempty"`
	AllDay     *bool      `json:"all_day" gorm:"default:true" validate:"omitempty,boolean"`
	Display    *string    `json:"display" gorm:"default:null" validate:"omitempty,min=3,max=60"`
	ResourceId *int64     `json:"resource_Id" gorm:"type:uint;default:null" validate:"omitempty,number,min=0,max=1000000"`
}
