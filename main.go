package main

import (
	"fmt"
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
		// initialize the game
		game := playground.NewGame()
		// start the game
		game.Start()
		fmt.Println("Game over!")
		return nil
	}

	// run the application
	app.Run(os.Args)
}
