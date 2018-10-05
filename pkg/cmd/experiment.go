package cmd

import (
	"fmt"
	"github.com/eliias/peterpawner/pkg/chess"
	"github.com/urfave/cli"
)

func Experiment(c *cli.Context) error {
	fmt.Println("Combined", chess.King|chess.Queen|chess.Bishop)
	return nil
}
