package main

import (
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"

	"github.com/deepakprakash/chattic/api/v1"
)

func setupMiddleware(e *gin.Engine) {
	// Setup all the global middleware that we need

	// Basic middlewares
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
}

func setupRoutes(e *gin.Engine) {
	// Setup our routes and route specific middleware

	e.GET("/", homePage)

	// Group routes for v1 API
	apiV1 := e.Group("/api/v1")

	apiV1.GET("/profile", v1.GetProfile)

	apiV1.POST("/auth/login", v1.Login)
	apiV1.POST("/auth/logout", v1.Logout)

	apiV1.POST("/messages", v1.CreateMessage)
}

func main() {
	app := cli.NewApp()

	app.Name = "Chattic"
	app.Usage = "Standalone, self-hosted instant messaging."
	app.Version = "0.1exp"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "bind, b",
			Value: ":4000",
			Usage: "Host:Port address for the server binding. Eg: `:4000`, `localhost:4000`, `0.0.0.0:4000`",
		},
	}

	app.Action = func(c *cli.Context) {
		// Action when the app is run.

		// Init the minimal gin engine
		engine := gin.New()

		// Setup middleware
		setupMiddleware(engine)

		// Setup routes
		setupRoutes(engine)

		// Run the server
		if err := http.ListenAndServe(c.String("bind"), engine); err != nil {
			// Error starting server. Log and exit.
			log.WithFields(log.Fields{
				"error": err,
			}).Fatal("Failed to start server.")
		}
	}

	app.Run(os.Args)
}
