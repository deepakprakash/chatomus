package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login using a username and password
func Login(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implemented yet. Check back later.")
}

// Logout the currently logged in user
func Logout(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implemented yet. Check back later.")
}
