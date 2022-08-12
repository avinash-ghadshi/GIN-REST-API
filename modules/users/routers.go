package users

import (
	"github.com/gin-gonic/gin"
)

func UserApis(r *gin.RouterGroup) {
	r.GET("/get", GetUsers)
	r.GET("/get/:id", RetrieveUser)
	r.POST("/add", AddUser)
	r.DELETE("/delete/:id", DeleteUser)
}
