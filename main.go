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
		{
			Name: "perft-divide",
			Aliases: []string{"d"},
			Usage: "Run perft divide",
			Action: chess.CmdPerftDivide,
		},
		{
			Name: "pieces",
			Usage: "Piece codes",
			Action: chess.CmdPieces,
		},
		{
			Name: "experimental",
			Aliases: []string{"e"},
			Usage: "Experimental",
			Action: chess.CmdExperiment,
		},
	}

	app.Run(os.Args)
}
