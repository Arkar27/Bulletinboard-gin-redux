package models

import (
	"time"

	"gorm.io/gorm"
)

type Password_resets struct {
	gorm.Model
	Email      string    `gorm:"type:varchar"`
	Token      string    `gorm:"type:varchar"`
	Created_at time.Time `gorm:"type:date"`
}
