package controllers

import (
	"net/http"

	_ "github.com/Arkar27/gin-bulletinboard/backend/docs"
	"github.com/Arkar27/gin-bulletinboard/backend/helper"
	_ "github.com/Arkar27/gin-bulletinboard/backend/models"
	services "github.com/Arkar27/gin-bulletinboard/backend/services"
	request "github.com/Arkar27/gin-bulletinboard/backend/types/request"
	response "github.com/Arkar27/gin-bulletinboard/backend/types/response"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	PostServiceInterface services.PostServiceInterface
}

func NewPostController(PostServiceInterface services.PostServiceInterface) *PostController {
	return &PostController{
		PostServiceInterface: PostServiceInterface,
	}
}

// Create a post
// @Summary Create a new post
// @Description Creates a new post
// @Tags POST
// @Accept json
// @Produce json
// @Param PostRequest body request.PostRequest true "Post Request Body"
// @Success 200 {object} response.Response{}
// @Router /api/posts [post]
// @Security ApiKeyAuth
func (controller *PostController) Create(ctx *gin.Context) {
	PostRequest := request.PostRequest{}
	err := ctx.ShouldBindJSON(&PostRequest)
	helper.ErrorPanic(err, ctx)

	controller.PostServiceInterface.Create(PostRequest, ctx)
	response := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	ctx.Header("Content-Type", "applicaton/json")
	ctx.JSON(http.StatusOK, response)
}

// Update a post
// @Summary Update a post
// @Description Update a post
// @Tags POST
// @Accept json
// @Produce json
// @Param PostRequest body request.PostRequest true "Post Request Body"
// @Param id  path string  true  "Update post by id"
// @Success 200 {object} response.Response{}
// @Router /api/posts/{id} [put]
// @Security ApiKeyAuth
func (controller *PostController) Update(ctx *gin.Context) {
	postId := ctx.Param("id")
	PostRequest := request.PostRequest{}
	err := ctx.ShouldBindJSON(&PostRequest)
	helper.ErrorPanic(err, ctx)

	retData := controller.PostServiceInterface.Update(PostRequest, postId , ctx)
	response := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   retData,
	}

	ctx.Header("Content-Type", "applicaton/json")
	ctx.JSON(http.StatusOK, response)
}

// Delete a post
// @Summary Delete a post
// @Description Returns nil
// @Tags POST
// @Produce plain
// @Param id  path string  true  "Delete post by id"
// @Success 200 {object} response.Response{}
// @Router /api/posts/{id} [delete]
// @Security ApiKeyAuth
func (controller *PostController) Delete(ctx *gin.Context) {
	postId := ctx.Param("id")
	controller.PostServiceInterface.Delete(postId, ctx)
	response := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	ctx.Header("Content-Type", "applicaton/json")
	ctx.JSON(http.StatusOK, response)
}

// Get post list
// @Summary Get post list
// @Description Returns post list
// @Tags POST
// @Produce plain
// @Success 200 {object} object "OK"
// @Router /api/posts [get]
// @Security ApiKeyAuth
func (controller *PostController) PostList(ctx *gin.Context) {
	data := controller.PostServiceInterface.FindAll(ctx)
	response := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	}

	ctx.Header("Content-Type", "applicaton/json")
	ctx.JSON(http.StatusOK, response)
}

// Find a post by id
// @Summary Find a post by id
// @Description Returns Found post
// @Tags POST
// @Produce plain
// @Param id  path string  true  "Find post by id"
// @Success 200 {object} response.Response{}
// @Router /api/posts/{id} [get]
// @Security ApiKeyAuth
func (controller *PostController) PostById(ctx *gin.Context) {
	postId := ctx.Param("id")
	data := controller.PostServiceInterface.FindOne(postId, ctx)
	response := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	}

	ctx.Header("Content-Type", "applicaton/json")
	ctx.JSON(http.StatusOK, response)
}
