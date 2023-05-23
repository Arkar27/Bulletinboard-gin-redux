package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorPanic(err interface{}, ctx *gin.Context) {
	if err != nil {
		switch e := err.(type) {
		case error:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		case string:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": e})
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unknown error"})
		}
		ctx.Abort()
		panic(err)
	}
}
