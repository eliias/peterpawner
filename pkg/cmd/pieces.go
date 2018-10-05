package cmd

import (
	"fmt"
	"github.com/eliias/peterpawner/pkg/chess"
	"github.com/urfave/cli"
)

func Pieces(c *cli.Context) error {
	fmt.Println("King", chess.King)
	fmt.Println("Queen", chess.Queen)
	fmt.Println("Rook", chess.Rook)
	fmt.Println("Bishop", chess.Bishop)
	fmt.Println("Knight", chess.Knight)
	fmt.Println("Pawn", chess.Pawn)
	return nil
}
