package main

import (
	"log"

	"example/biblioteca/config"
	"example/biblioteca/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDb()
}

func main() {
	route := gin.Default()

	routes.HomeRoute(route)

	log.Fatal(route.Run())
}
