package main

import (
	"github.com/Arkar27/gin-bulletinboard/backend/initializers"
	"github.com/Arkar27/gin-bulletinboard/backend/models"
)

func main() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Password_resets{})
}
