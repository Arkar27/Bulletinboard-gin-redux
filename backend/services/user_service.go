package services

import (
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
		Name:            user.Name,
		Email:           user.Email,
		Password:        user.Password,
		Profile:         user.Profile,
		Type:            user.Type,
		Phone:           user.Phone,
		Address:         user.Address,
		Dob:             dob,
		Created_user_id: uint(user.Created_user_id),
	}

	service.UserDaoInterface.Create(userModel, ctx)
}

// Update implements UserServiceInterface.
func (service *UserService) Update(user request.UserRequest, userId string, ctx *gin.Context) models.User {
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
		Created_user_id: uint(user.Created_user_id),
		Updated_user_id: uint(user.Updated_user_id),
	}
	data := service.UserDaoInterface.Update(userModel, userId, ctx)

	return data
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
