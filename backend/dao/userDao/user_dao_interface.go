package userDao

import (
	"github.com/Arkar27/gin-bulletinboard/backend/models"
	"github.com/gin-gonic/gin"
)

type UserDaoInterface interface {
	FindAll(ctx *gin.Context) []models.User
	FindOne(userId string, ctx *gin.Context) models.User
	Create(user models.User, ctx *gin.Context)
	Update(user models.User, userId string, ctx *gin.Context) models.User
	Delete(userId string, ctx *gin.Context)
}
