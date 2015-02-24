package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
)

func main() {
	app := cli.NewApp()

	app.Name = "Chattic"
	app.Usage = "Standalone, self-hosted instant messaging."
	app.Version = "0.1exp"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "bind, b",
			Value: ":4000",
			Usage: "Host:Port for the server to bind to. Eg: `:4000`, `localhost:4000`, `0.0.0.0:4000`",
		},
	}

	app.Action = func(c *cli.Context) {
		// Action when the app is run.

		// Setup the minimal gin router
		router := gin.Default()

		// Run the server
		router.Run(c.String("bind"))
	}

	app.Run(os.Args)
}
