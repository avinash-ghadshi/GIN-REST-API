package users

import (
	"github.com/gin-gonic/gin"
)

func UserApis(r *gin.RouterGroup) {
	r.GET("/get", GetUsers)
	r.POST("/add", AddUser)
}
