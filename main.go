package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strconv"
	"github.com/eliias/peterpawner/chess"
)

func main() {
	app := cli.NewApp()
	app.Name = "peterpawner"
	app.Usage = "A basic and very raw chess engine, written in GO."
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name: "perft",
			Aliases: []string{"p"},
			Usage: "Run perft",
			Action: func(c *cli.Context) error {
				if depth, err := strconv.Atoi(c.Args().First()); err == nil {
					fmt.Printf(chess.DebugPerft(depth))
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
