package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string    `gorm:"type:varchar;unique;not null"`
	Email           string    `gorm:"unique;not null"`
	Password        string    `gorm:"type:text;not null"`
	Profile         string    `gorm:"type:varchar"`
	Type            string    `gorm:"type:varchar(1);not null"`
	Phone           string    `gorm:"type:varchar(20)"`
	Address         string    `gorm:"type:varchar(255)"`
	Dob             time.Time `gorm:"type:date"`
	Created_user_id uint      `gorm:"type:int;not null"`
	Updated_user_id uint      `gorm:"type:int;not null"`
	Deleted_user_id uint      `gorm:"type:int"`
	Posts           []Post    `gorm:"foreignKey:Created_user_id"`
}

type AuthUser struct {
	User  User
	Token string
}
