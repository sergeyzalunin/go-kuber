package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Information about the service
// swagger:response InfoResponse
type InfoResponse struct {
	BuildTime string `json:"buildTime"`
	Commit    string `json:"commit"`
	Release   string `json:"release"`
}

// swagger:route GET /home home
// Returns a simple HTTP handler function which writes a response.
// responses:
//		200: InfoResponse

// home returns a simple HTTP handler function which writes a response.
func home(buildTime, commit, release string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// The result of home request
		// out: body
		info := InfoResponse{
			buildTime, commit, release,
		}

		c.JSON(http.StatusOK, info)
	}
}
