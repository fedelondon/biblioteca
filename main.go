package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	log.Fatal(route.Run())
}
