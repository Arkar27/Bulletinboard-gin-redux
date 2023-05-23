package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID              uint      `gorm:"primaryKey;autoIncrement;not null;"`
	Name            string    `gorm:"type:varchar;unique;not null"`
	Email           string    `gorm:"unique;not null"`
	Password        string    `gorm:"type:text;not null"`
	Profile         string    `gorm:"type:varchar"`
	Type            string    `gorm:"type:varchar(1);not null"`
	Phone           string    `gorm:"type:varchar(20)"`
	Address         string    `gorm:"type:varchar(255)"`
	Dob             time.Time `gorm:"type:date"`
	Create_user_id  uint      `gorm:"not null"`
	Updated_user_id uint      `gorm:"not null"`
	Deleted_user_id uint
	Posts           []Post `gorm:"foreignKey:Create_user_id"`
}

type AuthUser struct {
	User  User
	Token string
}
