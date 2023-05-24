package services

import (
	"github.com/Arkar27/gin-bulletinboard/backend/dao/postDao"
	"github.com/Arkar27/gin-bulletinboard/backend/models"
	"github.com/Arkar27/gin-bulletinboard/backend/types/request"
	"github.com/gin-gonic/gin"
)

type PostService struct {
	PostDaoInterface postDao.PostDaoInterface
}

// Create implements PostServiceInterface.
func (service *PostService) Create(post request.PostRequest, ctx *gin.Context) {
	postModel := models.Post{
		Title:           post.Title,
		Description:     post.Description,
		Status:          &post.Status,
		Created_user_id: post.Created_user_id,
		Updated_user_id: post.Updated_user_id,
	}
	service.PostDaoInterface.Create(postModel, ctx)
}

// Update implements PostServiceInterface.
func (service *PostService) Update(post request.PostRequest, postId string, ctx *gin.Context) models.Post {
	postModel := models.Post{
		Title:           post.Title,
		Description:     post.Description,
		Status:          &post.Status,
		Updated_user_id: post.Updated_user_id,
	}
	data := service.PostDaoInterface.Update(postModel, postId, ctx)

	return data
}

// Delete implements PostServiceInterface.
func (service *PostService) Delete(postId string, ctx *gin.Context) {
	service.PostDaoInterface.Delete(postId, ctx)
}

// FinAll implements PostServiceInterface.
func (service *PostService) FindAll(ctx *gin.Context) []models.Post {
	data := service.PostDaoInterface.FindAll(ctx)

	return data
}

// FindOne implements PostServiceInterface.
func (service *PostService) FindOne(postId string, ctx *gin.Context) models.Post {
	data := service.PostDaoInterface.FindOne(postId, ctx)

	return data
}

func NewPostService(PostDaoInterface postDao.PostDaoInterface) PostServiceInterface {
	return &PostService{
		PostDaoInterface: PostDaoInterface,
	}
}
