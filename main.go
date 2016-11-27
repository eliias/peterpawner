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

	var game = chess.Load("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	fmt.Println(chess.Save(game))
	//fmt.Println(chess.DebugGame(game))
	//fmt.Println(debug.Moves(board, depth, chess.COLOR_WHITE))
}
