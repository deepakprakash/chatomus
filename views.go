package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handle the home page
func homePage(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to Chattic!\nYou'll need to use the API endpoints to do anything meaningful.")
}
