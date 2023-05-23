package authDao

import (
	"github.com/Arkar27/gin-bulletinboard/backend/models"
	"github.com/gin-gonic/gin"
)

type AuthDaoInterface interface {
	Login(email string, password string, ctx *gin.Context) models.User
}
