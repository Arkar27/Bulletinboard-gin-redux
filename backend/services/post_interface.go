package services

import (
	"github.com/Arkar27/gin-bulletinboard/backend/models"
	"github.com/Arkar27/gin-bulletinboard/backend/types/request"
	"github.com/gin-gonic/gin"
)

type PostServiceInterface interface {
	FindAll(ctx *gin.Context) []models.Post
	FindOne(postId string, ctx *gin.Context) models.Post
	Create(post request.PostRequest, ctx *gin.Context)
	Update(post request.PostRequest, postId string, ctx *gin.Context) models.Post
	Delete(postId string, ctx *gin.Context)
}
