package debug

import (
	"github.com/eliias/peterpawner/pkg/chess"
	"github.com/eliias/peterpawner/pkg/generator"
)

func DebugHash(board []uint8) int {
	var h int = 0
	for _, v := range board {
		h += int(v)
	}
	return h
}

func perft(board []uint8, depth int, color uint8, enPassantFields []uint8) PerftResult {
	var nodes = 0
	var captures = 0
	var result PerftResult
	var enPassant []uint8
	var moves = generator.Generate(board, color, enPassantFields)
	result = stats(moves)

	if color == chess.ColorWhite {
		color = chess.ColorBlack
	} else {
		color = chess.ColorWhite
	}

	if depth == 1 {
		return result
	}

	for _, move := range moves {
		if move.EnPassant > 0 {
			enPassant = append(enPassant, move.EnPassant)
		}
	}

	for _, move := range moves {
		// stats
		result = stats(moves)

		// make move
		chess.MakeMove(board, move)

		// next level
		result = perft(board, depth-1, color, enPassant)

		// stats
		nodes += result.Nodes
		captures += result.Captures

		// undo move
		chess.UndoMove(board, move)
	}

	return PerftResult{Nodes: nodes, Captures: captures}
}

func Perft(depth int) PerftResult {
	return perft(chess.Start, depth, chess.ColorWhite, []uint8{})
}

func perftDivide(board []uint8, depth int, color uint8, enPassantFields []uint8) []PerftDivideResult {
	var moves = generator.Generate(board, color, enPassantFields)
	var result PerftResult
	var divides []PerftDivideResult
	var enPassant []uint8

	if color == chess.ColorWhite {
		color = chess.ColorBlack
	} else {
		color = chess.ColorWhite
	}

	for _, move := range moves {
		if move.EnPassant > 0 {
			enPassant = append(enPassant, move.EnPassant)
		}
	}

	for _, move := range moves {
		// divide
		var divide PerftDivideResult

		// stats
		result = stats(moves)

		// make move
		chess.MakeMove(board, move)

		// next level
		result = perft(board, depth-1, color, enPassant)

		// stats
		divide = PerftDivideResult{Move: move, Nodes: result.Nodes}
		divides = append(divides, divide)

		// undo move
		chess.UndoMove(board, move)
	}

	return divides
}

type PerftResult struct {
	Nodes      int
	Captures   int
	EnPassant  int
	Castles    int
	Promotions int
	Checks     int
	Checkmates int
}

type PerftDivideResult struct {
	Move  *chess.Move
	Nodes int
}

func stats(moves []*chess.Move) PerftResult {
	var result = PerftResult{Nodes: len(moves)}
	for _, move := range moves {
		if move.Prev != chess.Empty {
			result.Captures += 1
		}
	}
	return result
}

func PerftDivide(depth int) []PerftDivideResult {
	return perftDivide(chess.Start, depth, chess.ColorWhite, []uint8{})
}
