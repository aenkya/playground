package main

import (
	"os"

	"enkya.org/playground/playground"
	"github.com/urfave/cli/v2"
)

func main() {
	// initialize commandline based application
	app := cli.NewApp()
	app.Name = "Snakes & Ladders"
	app.Usage = "A simple snakes and ladders game"
	app.Action = func(c *cli.Context) error {
		game := playground.NewGame()
		game.Start()
		return nil
	}

	// run the application
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
