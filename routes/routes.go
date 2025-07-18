package routes

import (
	"PostGo/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AppRoutes(router *gin.Engine, db *gorm.DB) {
	postController := controllers.NewPostController(db)

	api := router.Group("/api/post")
	{
		api.GET("/", postController.GetPosts)
		api.GET("/:id", postController.GetPostByID)
		api.POST("/", postController.CreatePost)
		api.PUT("/:id", postController.UpdatePost)
		api.DELETE("/:id", postController.DeletePost)
	}
}
