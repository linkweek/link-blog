package admin

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "This is admin index page.",
	})
}
