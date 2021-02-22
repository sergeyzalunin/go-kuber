package handlers

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

const isReadyDuration = 0 * time.Second

// Router register necessary routes and returns an instance of a router.
func Router(buildTime, commit, release string) *gin.Engine {
	isReady := &atomic.Value{}
	isReady.Store(false)

	go func() {
		log.Printf("Readyz probe is negative by default...")
		time.Sleep(isReadyDuration)
		isReady.Store(true)
		log.Printf("Readyz probe is positive.")
	}()

	r := gin.Default()
	r.GET("/home", home(buildTime, commit, release))
	r.GET("/healthz", healthz)
	r.GET("/readyz", readyz(isReady))

	return r
}
