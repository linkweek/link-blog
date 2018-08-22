package routes

import (
	"github.com/gin-gonic/gin"
	handler "linkweek.cn/link-blog/handlers/admin"
	"linkweek.cn/link-blog/middleware"
)

func Backend(router *gin.Engine) *gin.Engine {
	admin := router.Group("/admin")
	admin.Use(middleware.Auth())

	admin.GET("/", handler.Index)
	admin.GET("/login", handler.Login)

	return router
}
