package cmd

import (
	"fmt"
	"github.com/eliias/peterpawner/pkg/debug"
	"github.com/urfave/cli"
	"strconv"
	"time"
)

func Perft(c *cli.Context) error {
	if depth, err := strconv.Atoi(c.Args().First()); err == nil {
		now := time.Now()
		result := debug.DebugPerft(depth)
		time := time.Now().Sub(now).Nanoseconds() / 1000000
		fmt.Println(result)
		fmt.Println("Calculated in", time, "milliseconds")
	}
	return nil
}
