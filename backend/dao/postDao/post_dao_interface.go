package postDao

import (
	"github.com/Arkar27/gin-bulletinboard/backend/models"
	"github.com/gin-gonic/gin"
)

type PostDaoInterface interface {
	FindAll(ctx *gin.Context) []models.Post
	FindOne(postId string, ctx *gin.Context) models.Post
	Create(post models.Post, ctx *gin.Context)
	Update(post models.Post, postId string, ctx *gin.Context) models.Post
	Delete(postId string, ctx *gin.Context)
}
