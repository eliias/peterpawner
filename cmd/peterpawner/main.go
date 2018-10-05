package main

import (
	"github.com/eliias/peterpawner/pkg/cmd"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "peterpawner"
	app.Usage = "A basic and very raw chess engine, written in GO."
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:    "perft",
			Aliases: []string{"p"},
			Usage:   "Run perft",
			Action:  cmd.Perft,
		},
		{
			Name:    "perft-divide",
			Aliases: []string{"d"},
			Usage:   "Run perft divide",
			Action:  cmd.PerftDivide,
		},
		{
			Name:   "pieces",
			Usage:  "Piece codes",
			Action: cmd.Pieces,
		},
		{
			Name:    "experimental",
			Aliases: []string{"e"},
			Usage:   "Experimental",
			Action:  cmd.Experiment,
		},
	}

	app.Run(os.Args)
}
