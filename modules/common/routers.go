package common

import "github.com/gin-gonic/gin"

func CommonApis(r *gin.RouterGroup) {
	r.GET("/", Welcome)
	r.GET("/test", Test)
	r.GET("/ping", Ping)
	r.GET("/health", Health)
}
