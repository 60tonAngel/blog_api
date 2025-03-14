package routes

import (
	"blog_api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) {
	api := engine.Group("/api/v1")
	{
		posts := api.Group("/posts")
		{
			posts.GET("", controllers.GetAllPosts)
			posts.POST("", controllers.CreatePost)
			posts.GET("/:id", controllers.GetPost)
			posts.PUT("/:id", controllers.UpdatePost)
			posts.DELETE("/:id", controllers.DeletePost)
		}
	}

	// 健康检查端点
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})
}
