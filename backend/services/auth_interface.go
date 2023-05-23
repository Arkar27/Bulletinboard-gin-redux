package services

import (
	"github.com/Arkar27/gin-bulletinboard/backend/types/request"
	"github.com/gin-gonic/gin"
)

type AuthServiceInterface interface {
	Authenticate(user request.LoginRequest, ctx *gin.Context) interface{}
}
