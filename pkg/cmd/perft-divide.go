package cmd

import (
	"fmt"
	"github.com/eliias/peterpawner/pkg/debug"
	"github.com/urfave/cli"
	"strconv"
)

func PerftDivide(c *cli.Context) error {
	if depth, err := strconv.Atoi(c.Args().First()); err == nil {
		result := debug.DebugPerftDivide(depth)
		fmt.Println(result)
	}
	return nil
}
