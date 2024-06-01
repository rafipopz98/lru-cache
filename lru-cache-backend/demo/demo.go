package demo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World")
	}
}
