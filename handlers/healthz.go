package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// healthz is a liveness probe.
func healthz(c *gin.Context) {
	c.Status(http.StatusOK)
}
