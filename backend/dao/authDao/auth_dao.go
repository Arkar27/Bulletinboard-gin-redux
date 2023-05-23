package authDao

import (
	"github.com/Arkar27/gin-bulletinboard/backend/helper"
	"github.com/Arkar27/gin-bulletinboard/backend/initializers"
	"github.com/Arkar27/gin-bulletinboard/backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthDao struct {
	DB *gorm.DB
}

func NewAuthDao(DB *gorm.DB) AuthDaoInterface {
	return &AuthDao{DB: DB}
}

func (authDao *AuthDao) Login(email string, password string, ctx *gin.Context) models.User {
	var user models.User
	result := initializers.DB.Where("email = ? AND password = ?", email, password).First(&user)
	helper.ErrorPanic(result.Error, ctx)
	return user
}
