package proxy

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleUserServices(c *gin.Context) {
	fullPath := c.Request.URL.Path
	query := c.Request.URL.RawQuery
	microserviceUrl, err := GetTargetURL(fullPath, query)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	proxy, err := ForwardRequest(microserviceUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create proxy"})
		return
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
