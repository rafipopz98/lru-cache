package services

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// handler for set  key value pair
func PostCache() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		keyStr := ctx.Query("key")
		value := ctx.Query("value")
		durationStr := ctx.Query("duration")

		key, err := strconv.Atoi(keyStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid key"})
			return
		}

		duration, err := strconv.Atoi(durationStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid duration"})
			return
		}

		cacheInstance.Set(key, value, time.Duration(duration)*time.Second)
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}

// handler for get one key value pair
func GetCache() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		keyStr := ctx.Query("key")
		key, err := strconv.Atoi(keyStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid key"})
			return
		}
		var expiry time.Time
		value, expiry, found := cacheInstance.Get(key)
		if !found {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Key not found or expired"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"key": key, "value": value, "expiry": expiry})
	}
}

// handler for get all key value pair
func GetAllCache() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cacheItems := cacheInstance.GetAll()
		ctx.JSON(http.StatusOK, gin.H{"cacheItems": cacheItems})

	}
}

func DeleteOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		keyStr := ctx.Query("key")
		keyInt, err := strconv.Atoi(keyStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid key"})
		}
		cacheItems := cacheInstance.DeleteOne(keyInt)
		if !cacheItems {
			ctx.JSON(http.StatusOK, gin.H{"message": "Such key doesnot exist or it is expired"})

		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
		}

	}
}
