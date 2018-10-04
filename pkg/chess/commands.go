package chess

import (
	"fmt"
	"github.com/urfave/cli"
	"strconv"
	"time"
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

func CmdPerftDivide(c *cli.Context) error {
	if depth, err := strconv.Atoi(c.Args().First()); err == nil {
		result := DebugPerftDivide(depth)
		fmt.Println(result)
	}
	return nil
}

func CmdPieces(c *cli.Context) error {
	fmt.Println("King", KING)
	fmt.Println("Queen", QUEEN)
	fmt.Println("Rook", ROOK)
	fmt.Println("Bishop", BISHOP)
	fmt.Println("Knight", KNIGHT)
	fmt.Println("Pawn", PAWN)
	return nil
}

func CmdExperiment(c *cli.Context) error {
	fmt.Println("Combined", KING|QUEEN|BISHOP)
	return nil
}
