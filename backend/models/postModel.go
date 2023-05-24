package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title           string `gorm:"type:varchar;unique;not null"`
	Description     string `gorm:"not null"`
	Status          *int   `gorm:"not null"`
	Created_user_id uint   `gorm:"type:int;not null"`
	Updated_user_id uint   `gorm:"type:int;not null"`
	Deleted_user_id uint   `gorm:"type:int"`
}
