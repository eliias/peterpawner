package main

import (
	"github.com/urfave/cli"
	"os"
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
			Action: chess.CmdPerft,
		},
	}

	app.Run(os.Args)
}
