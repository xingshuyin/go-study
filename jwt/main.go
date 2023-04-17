package main

import (
	"os"
	routes "project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.Auth_router(router)
	routes.User_router(router)
	router.GET("/api-1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for api-1"})
	})
	router.GET("api-2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": " Access api-2"})
	})
	router.Run(":" + port)
}
