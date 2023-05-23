package postDao

import (
	helper "github.com/Arkar27/gin-bulletinboard/backend/helper"
	"github.com/Arkar27/gin-bulletinboard/backend/initializers"
	"github.com/Arkar27/gin-bulletinboard/backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostDao struct {
	DB *gorm.DB
}

func NewPostDao(DB *gorm.DB) PostDaoInterface {
	return &PostDao{DB: DB}
}

// Create implements PostDaoInterface.
func (PostDao *PostDao) Create(post models.Post, ctx *gin.Context) {

	result := initializers.DB.Create(&post)
	helper.ErrorPanic(result.Error, ctx)
}

// Update implements PostDaoInterface.
func (postDao *PostDao) Update(post models.Post, postId string, ctx *gin.Context) models.Post {

	result := initializers.DB.Model(&post).Where("id = ?", postId).Updates(post)
	helper.ErrorPanic(result.Error, ctx)
	result = initializers.DB.First(&post, postId)
	helper.ErrorPanic(result.Error, ctx)
	return post
}

// Delete implements PostDaoInterface.
func (postDao *PostDao) Delete(postId string, ctx *gin.Context) {
	var post models.Post
	result := initializers.DB.Delete(&post, postId)
	helper.ErrorPanic(result.Error, ctx)
}

// FindAll implements PostDaoInterface.
func (*PostDao) FindAll(ctx *gin.Context) []models.Post {
	var posts []models.Post
	result := initializers.DB.Find(&posts)
	helper.ErrorPanic(result.Error, ctx)

	return posts
}

// FindOne implements PostDaoInterface.
func (*PostDao) FindOne(postId string, ctx *gin.Context) models.Post {
	var post models.Post
	result := initializers.DB.First(&post, postId)
	helper.ErrorPanic(result.Error, ctx)
	return post
}
