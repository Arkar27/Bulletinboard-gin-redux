package main

import (
	"os"

	"github.com/Arkar27/gin-bulletinboard/backend/initializers"
	"github.com/Arkar27/gin-bulletinboard/backend/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	router := gin.Default()
	routes.Routes(router)
	port := os.Getenv("PORT")
	router.Run(":" + port)
}
