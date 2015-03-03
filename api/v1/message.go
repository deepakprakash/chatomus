package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST or Create a new message.
func CreateMessage(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implemented yet. Check back later.")
}
