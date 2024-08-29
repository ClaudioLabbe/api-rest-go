package main

import (
	"github.com/gin-gonic/gin"

	"api-rest-go/src/routes"
)

func main() {
	router := gin.Default()

	routes.RegisterRouter(router)

	router.Run("localhost:8080")
}
