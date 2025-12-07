package main

import (
	"github.com/Umair1231/goMicroserviceGateway/internal/proxy"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/api/users/*path", proxy.HandleUserServices)
	router.Run("localhost:8080")
}
