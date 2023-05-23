package services

import (
	"github.com/Arkar27/gin-bulletinboard/backend/models"
	"github.com/Arkar27/gin-bulletinboard/backend/types/request"
	"github.com/gin-gonic/gin"
)

type UserServiceInterface interface {
	FindAll(ctx *gin.Context) []models.User
	FindOne(userId string, ctx *gin.Context) models.User
	Create(user request.UserRequest, ctx *gin.Context)
	Update(user request.UserRequest, userId string, ctx *gin.Context) models.AuthUser
	Delete(userId string, ctx *gin.Context)
}
