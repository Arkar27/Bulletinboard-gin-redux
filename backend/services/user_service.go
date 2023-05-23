package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Arkar27/gin-bulletinboard/backend/dao/userDao"
	"github.com/Arkar27/gin-bulletinboard/backend/models"
	"github.com/Arkar27/gin-bulletinboard/backend/types/request"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	UserDaoInterface userDao.UserDaoInterface
}

// Create implements UserServiceInterface.
func (service *UserService) Create(user request.UserRequest, ctx *gin.Context) {
	dob, _ := time.Parse("2006-01-02", user.Dob)
	userModel := models.User{
		Name:           user.Name,
		Email:          user.Email,
		Password:       user.Password,
		Profile:        user.Profile,
		Type:           user.Type,
		Phone:          user.Phone,
		Address:        user.Address,
		Dob:            dob,
		Create_user_id: uint(user.Create_user_id),
	}

	service.UserDaoInterface.Create(userModel, ctx)
}

// Update implements UserServiceInterface.
func (service *UserService) Update(user request.UserRequest, userId string, ctx *gin.Context) models.AuthUser {
	dob, _ := time.Parse("2006-01-02", user.Dob)
	userModel := models.User{
		Name:            user.Name,
		Email:           user.Email,
		Password:        user.Password,
		Profile:         user.Profile,
		Type:            user.Type,
		Phone:           user.Phone,
		Address:         user.Address,
		Dob:             dob,
		Create_user_id:  uint(user.Create_user_id),
		Updated_user_id: uint(user.Updated_user_id),
	}
	data := service.UserDaoInterface.Update(userModel, userId, ctx)
	var token string

	// check self updating
	fmt.Println("userId param-=-=> ", userId)
	fmt.Println("userId updateId-=-=> ", strconv.Itoa(int(user.Updated_user_id)))
	if userId == strconv.Itoa(int(user.Updated_user_id)) {
		token, _ = GenerateToken(data.Email, data.Name)
	}
	retData := models.AuthUser{
		User:  data,
		Token: token,
	}
	return retData
}

// Delete implements UserServiceInterface.
func (service *UserService) Delete(userId string, ctx *gin.Context) {
	service.UserDaoInterface.Delete(userId, ctx)
}

// FindAll implements UserServiceInterface.
func (service *UserService) FindAll(ctx *gin.Context) []models.User {
	data := service.UserDaoInterface.FindAll(ctx)

	return data
}

// FindOne implements UserServiceInterface.
func (service *UserService) FindOne(userId string, ctx *gin.Context) models.User {
	data := service.UserDaoInterface.FindOne(userId, ctx)

	return data
}

func NewUserService(UserDaoInterface userDao.UserDaoInterface) UserServiceInterface {
	return &UserService{
		UserDaoInterface: UserDaoInterface,
	}
}
