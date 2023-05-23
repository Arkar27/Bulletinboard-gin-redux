package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID              uint   `gorm:"primaryKey;autoIncrement;not null"`
	Title           string `gorm:"type:varchar;unique;not null"`
	Description     string `gorm:"not null"`
	Status          *int    `gorm:"not null"`
	Create_user_id  int    `gorm:"not null"`
	Updated_user_id int    `gorm:"not null"`
	Deleted_user_id int
}
