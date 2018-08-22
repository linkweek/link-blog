package main

import (
	"github.com/gin-gonic/gin"
	"linkweek.cn/link-blog/routes"
)

func main() {
	app := gin.New()

	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	routes.Backend(app)
	routes.Frontend(app)

	app.Run(":8080")
}
