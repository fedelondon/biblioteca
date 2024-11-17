package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeRoute(route *gin.Engine) {
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
}
