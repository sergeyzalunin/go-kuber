package handlers

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

// readyz is a readiness probe.
func readyz(isReady *atomic.Value) gin.HandlerFunc {
	return func(c *gin.Context) {
		if isReady == nil || !isReady.Load().(bool) {
			c.Status(http.StatusServiceUnavailable)
			c.Error(fmt.Errorf(http.StatusText(http.StatusServiceUnavailable)))
			return
		}
		c.Status(http.StatusOK)
	}
}
