package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
}
