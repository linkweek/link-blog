package main

import (
	"github.com/gin-gonic/gin"
	"linkweek.cn/link-blog/routes"
)

func main() {
	app := gin.New()

	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	app.Static("/assets", "./assets")
	app.StaticFile("/favicon.ico", "./resources/favicon.ico")

	routes.Backend(app)
	routes.Frontend(app)

	app.Run(":8080")
}
