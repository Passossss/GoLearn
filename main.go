package main

import (
	"PostGo/config"
	_ "PostGo/docs"
	"PostGo/routes"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db := config.ConnectDB()
	router := gin.Default()

	// postController := controllers.NewPostController(db)

	// router.GET("/api/post", postController.GetPosts)
	// router.GET("/api/post/:id", postController.GetPostByID)
	// router.POST("/api/post", postController.CreatePost)
	// router.PUT("/api/post/:id", postController.UpdatePost)
	// router.DELETE("/api/post/:id", postController.DeletePost)

	routes.AppRoutes(router, db)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":3001")
}
