package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Api v1 Opening Endpoint",
			})
		})
		v1.POST("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Post Opening",
			})
		})
		v1.DELETE("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "delete opening",
			})
		})
		v1.PUT("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "update opening",
			})
		})
		v1.GET("/openings", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "all the openings",
			})
		})
	}
}
