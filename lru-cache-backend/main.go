package main

import (
	"lru-cache-backend/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())
	routes.CacheRouter(router)
	router.Run(":" + port)
}
