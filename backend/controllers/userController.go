package controllers

import (
	"net/http"

	_ "github.com/Arkar27/gin-bulletinboard/backend/docs"
	"github.com/Arkar27/gin-bulletinboard/backend/helper"
	_ "github.com/Arkar27/gin-bulletinboard/backend/models"
	services "github.com/Arkar27/gin-bulletinboard/backend/services"
	request "github.com/Arkar27/gin-bulletinboard/backend/types/request"
	"github.com/Arkar27/gin-bulletinboard/backend/types/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserServiceInterface services.UserServiceInterface
}

func NewUserController(UserServiceInterface services.UserServiceInterface) *UserController {
	return &UserController{
		UserServiceInterface: UserServiceInterface,
	}
}

// Create a user
// @Summary Create a new user
// @Description Creates a new user
// @Tags USER
// @Accept json
// @Produce json
// @Param UserRequest body request.UserRequest true "User Request Body"
// @Success 200 {object} response.Response{}
// @Router /api/users [post]
// @Security ApiKeyAuth
func (controller *UserController) CreateUser(ctx *gin.Context) {
	UserRequest := request.UserRequest{}
	err := ctx.ShouldBindJSON(&UserRequest)
	helper.ErrorPanic(err, ctx)

	controller.UserServiceInterface.Create(UserRequest, ctx)
	response := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	ctx.Header("Content-Type", "applicaton/json")
	ctx.JSON(http.StatusOK, response)
}

// Update a user
// @Summary Update a user
// @Description Update a user
// @Tags USER
// @Accept json
// @Produce json
// @Param UserRequest body request.UserRequest true "User Request Body"
// @Param id  path string  true  "Update user by id"
// @Success 200 {object} response.Response{}
// @Router /api/users/{id} [put]
// @Security ApiKeyAuth
func (controller *UserController) UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	UserRequest := request.UserRequest{}
	err := ctx.ShouldBindJSON(&UserRequest)
	helper.ErrorPanic(err, ctx)

	retData := controller.UserServiceInterface.Update(UserRequest, userId, ctx)
	response := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   retData,
	}

	ctx.Header("Content-Type", "applicaton/json")
	ctx.JSON(http.StatusOK, response)
}

// Delete a user
// @Summary Delete a user
// @Description Returns nil
// @Tags USER
// @Produce plain
// @Param id  path string true  "Delete user by id"
// @Success 200 {object} response.Response{}
// @Router /api/users/{id} [delete]
// @Security ApiKeyAuth
func (controller *UserController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	controller.UserServiceInterface.Delete(userId, ctx)
	response := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	ctx.Header("Content-Type", "applicaton/json")
	ctx.JSON(http.StatusOK, response)
}

// Get user list
// @Summary Get user list
// @Description Returns user list
// @Tags USER
// @Produce plain
// @Success 200 {object} object "OK"
// @Router /api/users [get]
// @Security ApiKeyAuth
func (controller *UserController) GetUserList(ctx *gin.Context) {
	data := controller.UserServiceInterface.FindAll(ctx)
	response := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	}

	ctx.Header("Content-Type", "applicaton/json")
	ctx.JSON(http.StatusOK, response)
}

// Find a user by id
// @Summary Find a user by id
// @Description Returns Found user
// @Tags USER
// @Produce json
// @Param id  path string  true  "Find user by id"
// @Success 200 {object} response.Response{}
// @Router /api/users/{id} [get]
// @Security ApiKeyAuth
func (controller *UserController) GetUserById(ctx *gin.Context) {
	userId := ctx.Param("id")
	data := controller.UserServiceInterface.FindOne(userId, ctx)
	response := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	}

	ctx.Header("Content-Type", "applicaton/json")
	ctx.JSON(http.StatusOK, response)
}
