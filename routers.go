package main

import (
	"libs/middlewares"
	"libs/modules/common"
	"libs/modules/users"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

func setupRouter() *gin.Engine {

	gin.SetMode(gin.DebugMode)

	r := gin.New()
	r.Use(gin.Recovery(), middlewares.Logger(), gindump.Dump())
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.HandleMethodNotAllowed = true

	v1 := r.Group("/api")
	v2 := r.Group("/")

	// Basic APIs
	common.CommonApis(v2.Group("/"))

	// Users APIs
	users.UserApis(v1.Group("/user"))
	// settings.SettingsApis(v1.Group("/settings"))

	return r
}
