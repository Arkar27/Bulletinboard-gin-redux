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

type LoginController struct {
	AuthServiceInterface services.AuthServiceInterface
}

func NewLoginController(AuthServiceInterface services.AuthServiceInterface) *LoginController {
	return &LoginController{
		AuthServiceInterface: AuthServiceInterface,
	}
}

// Login
// @Summary User Login
// @Description Authenticates user login
// @Tags LOGIN
// @Accept json
// @Produce json
// @Param LoginRequest body request.LoginRequest true "Login Request Body"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /api/login [post]
func (controller *LoginController) Login(ctx *gin.Context) {
	LoginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&LoginRequest)
	helper.ErrorPanic(err, ctx)

	data := controller.AuthServiceInterface.Authenticate(LoginRequest, ctx)
	if data != struct{}{} {
		response := response.Response{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   data,
		}

		ctx.Header("Content-Type", "applicaton/json")
		ctx.JSON(http.StatusOK, response)
	} else {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login data"})
		return
	}
}
