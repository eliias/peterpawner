package main

import (
	"fmt"
	"github.com/eliias/peterpawner/chess"
)

func main() {
	fmt.Println("\n-----------------\npeterpawner 1.0\nHannes Moser 2016\n-----------------\n")

	var board []uint8

	board = chess.Start
	fmt.Println(chess.DebugBoard(board))

	var depth = 3
	fmt.Println(chess.DebugPerft(depth))
	//fmt.Println(chess.DebugPerftDivide(depth))

	//var game = chess.Start
	//fmt.Println(chess.Save(game))
	//fmt.Println(chess.DebugGame(game))
	//fmt.Println(chess.DebugMoves(game.Board, 1, chess.COLOR_WHITE))
}
