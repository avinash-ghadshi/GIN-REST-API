package main

import (
	"libs/modules/common"
	"libs/modules/users"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	v1 := r.Group("/api")
	v2 := r.Group("/")

	// Basic APIs
	common.CommonApis(v2.Group("/"))
	// r.GET("/", common.Welcome)
	// r.GET("/test", common.Test)
	// r.GET("/ping", common.Ping)
	// r.GET("/health", common.Health)

	// User APIs
	users.UserApis(v1.Group("/user"))
	// r.GET("/get", users.GetUsers)
	// r.POST("/add", users.AddUser)

	return r
}
