package main

import (
	"PostGo/config"
	_ "PostGo/docs"
	"PostGo/routes"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title PostGo API
// @version 1.0
// @description Uma API de exemplo para gerenciamento de posts em Go com Gin e GORM.
// @host localhost:3001
// @BasePath /api
func main() {
	db := config.ConnectDB()
	router := gin.Default()

	routes.AppRoutes(router, db)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":3001")
}
