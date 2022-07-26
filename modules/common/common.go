package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "This is test response",
	})
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "PONG",
	})
}

func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Gin Server is running",
	})
}

func Welcome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the GIN-REST-API",
	})
}
