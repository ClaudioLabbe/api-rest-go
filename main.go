package main

import (
	"github.com/gin-gonic/gin"

	"api-rest-go/src/routes"

	"api-rest-go/src/config"
)

func main() {

	config.InitDb()

	router := gin.Default()

	routes.RegisterRouter(router)

	router.Run("localhost:8080")
}
