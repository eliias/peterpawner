package chess

import (
	"github.com/urfave/cli"
	"strconv"
	"time"
	"fmt"
)

func CmdPerft(c *cli.Context) error {
	if depth, err := strconv.Atoi(c.Args().First()); err == nil {
		now := time.Now()
		result := DebugPerft(depth)
		time := time.Now().Sub(now).Nanoseconds() / 1000000
		fmt.Println(result)
		fmt.Println("Calculated in", time, "milliseconds")
	}
	return nil
}
