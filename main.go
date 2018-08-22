package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.New()

	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	app.Run(":8080")
}
