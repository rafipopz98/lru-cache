package routes

import (
	"lru-cache-backend/services"

	"github.com/gin-gonic/gin"
)

func CacheRouter(incomingRoute *gin.Engine) {
	incomingRoute.GET("/get", services.GetCache())
	incomingRoute.POST("/add", services.PostCache())
	incomingRoute.GET("/all", services.GetAllCache())
	incomingRoute.POST("/delete", services.DeleteOne())
}
