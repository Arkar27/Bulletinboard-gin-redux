package userDao

import (
	helper "github.com/Arkar27/gin-bulletinboard/backend/helper"
	"github.com/Arkar27/gin-bulletinboard/backend/initializers"
	"github.com/Arkar27/gin-bulletinboard/backend/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserDao struct {
	DB *gorm.DB
}

func NewUserDao(DB *gorm.DB) UserDaoInterface {
	return &UserDao{DB: DB}
}

// Create implements UserDaoInterface.
func (userDao *UserDao) Create(user models.User, ctx *gin.Context) {

	result := initializers.DB.Create(&user)
	helper.ErrorPanic(result.Error, ctx)
}

// Update implements UserDaoInterface.
func (userDao *UserDao) Update(user models.User, userId string, ctx *gin.Context) models.User {

	result := initializers.DB.Model(&user).Where("id = ?", userId).Updates(user)
	helper.ErrorPanic(result.Error, ctx)
	result = initializers.DB.First(&user, userId)
	helper.ErrorPanic(result.Error, ctx)
	return user
}

// Delete implements UserDaoInterface.
func (userDao *UserDao) Delete(userId string, ctx *gin.Context) {
	var user models.User

	// soft delete
	result := userDao.DB.Where("created_user_id = ?", userId).Delete(&models.Post{})
	helper.ErrorPanic(result.Error, ctx)
	result = initializers.DB.Delete(&user, userId)
	helper.ErrorPanic(result.Error, ctx)

	// hard delete
	// result := userDao.DB.Unscoped().Where("created_user_id = ?", userId).Delete(&models.Post{})
	// helper.ErrorPanic(result.Error)
	// result = initializers.DB.Unscoped().Delete(&user, userId)
	// helper.ErrorPanic(result.Error)

}

// FindAll implements UserDaoInterface.
func (*UserDao) FindAll(ctx *gin.Context) []models.User {
	var users []models.User

	// Get userId and userType from session
	session := sessions.Default(ctx)
	userId := session.Get("userId")
	userType := session.Get("userType")

	// Admin will see all users and member only see its created users
	if userType == "0" {

		result := initializers.DB.Model(&users).Preload("Posts").Find(&users)
		helper.ErrorPanic(result.Error, ctx)
	} else {

		result := initializers.DB.Model(&users).Where("created_user_id = ?", userId).Preload("Posts").Find(&users)
		helper.ErrorPanic(result.Error, ctx)
	}

	return users
}

// FindOne implements UserDaoInterface.
func (*UserDao) FindOne(userId string, ctx *gin.Context) models.User {
	var user models.User
	result := initializers.DB.Preload("Posts").First(&user, userId)
	helper.ErrorPanic(result.Error, ctx)
	return user
}
