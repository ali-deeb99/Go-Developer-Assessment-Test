package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name                string
	Phone_number        string
	Otp                 string
	Otp_expiration_time time.Time
}
