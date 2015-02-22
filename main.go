package main

import (
  "os"
  "fmt"

  "github.com/codegangsta/cli"
)

func main() {
  app := cli.NewApp()

  app.Name = "Chattic"
  app.Usage = "Standalone, self-hosted instant messaging."
  app.Version = "0.1exp"

  app.Flags = []cli.Flag {
    cli.IntFlag{
      Name: "port, p",
      Value: 4000,
      Usage: "Port to be used for the server.",
    },
  }

  app.Action = func(c *cli.Context) {
    fmt.Println("Using port:", c.Int("port"))
  }

  app.Run(os.Args)
}
