package internal

import (
	"net/http"

	"github.com/gin-contrib/graceful"
	"github.com/gin-gonic/gin"
)

func setupRoutes(r *graceful.Graceful) {
	// Insert Routes Here
	r.GET("/ready", standardLogger(), healthCheck)
}

func healthCheck(c *gin.Context) {
	c.String(http.StatusOK, "ready")
}
