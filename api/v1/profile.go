package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get the profile of currently logged in user
func GetProfile(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implemented yet. Check back later.")
}
