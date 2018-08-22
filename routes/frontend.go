package routes

import (
	"github.com/gin-gonic/gin"
	handler "linkweek.cn/link-blog/handlers/blog"
)

func Frontend(router *gin.Engine) *gin.Engine {

	router.GET("/", handler.Index)

	return router
}
