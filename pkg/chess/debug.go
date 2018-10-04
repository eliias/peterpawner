package chess

import (
	"strconv"
)

const Rank = "12345678"
const File = "abcdefgh"

func PieceName(piece uint8) string {
	switch piece {
	// white
	case W_KING:
		return "K"
	case W_QUEEN:
		return "Q"
	case W_ROOK:
		return "R"
	case W_BISHOP:
		return "B"
	case W_KNIGHT:
		return "N"
	case W_PAWN:
		return "P"
	// black
	case B_KING:
		return "k"
	case B_QUEEN:
		return "q"
	case B_ROOK:
		return "r"
	case B_BISHOP:
		return "b"
	case B_KNIGHT:
		return "n"
	case B_PAWN:
		return "p"
	}
	return "."
}

func ColorName(color uint8) string {
	if color == COLOR_WHITE {
		return "w"
	}
	return "b"
}

func DebugAttacks(list []uint8) string {
	var str = ""
	for i := 0; i < 64; i += 1 {
		if i%8 == 0 && i > 0 {
			str += "\n"
		}
		if Contains(list, uint8(i)) {
			str += "1"
		} else {
			str += "0"
		}
	}
	return str
}

func DebugBoard(board []uint8) string {
	var str = ""
	var col int
	var row int = 0
	for i := 0; i < 100; i += 1 {
		if i%10 == 0 && i > 0 {
			row += 1
			str += "\n"
		}
		col = i - row*10

		if col == 0 && row == 0 || col == 9 && row == 0 || col == 0 && row == 9 || col == 9 && row == 9 {
			str += "+"
		} else if row == 0 && col > 0 && col < 9 {
			str += "-"
		} else if row == 9 && col > 0 && col < 9 {
			str += "-"
		} else if col == 0 || col == 9 {
			str += "|"
		} else {
			str += PieceName(board[col+8*(row-1)-1])
		}
	}

	return str
}

func DebugGame(game Game) string {
	var str = "Game:\n"
	str += DebugBoard(game.Board) + "\n"
	str += "Active Color: " + ColorName(game.Color) + "\n"
	str += "White - Castle King Side: " + strconv.FormatBool(game.WhiteCastleKingSide) + "\n"
	str += "        Castle Queen Side: " + strconv.FormatBool(game.WhiteCastleQueenSide) + "\n"
	str += "Black - Castle King Side: " + strconv.FormatBool(game.BlackCastleKingSide) + "\n"
	str += "        Castle Queen Side: " + strconv.FormatBool(game.BlackCastleQueenSide) + "\n"
	str += "En Passant: " + DebugPos(game.EnPassant) + "\n"
	str += "Halfmove: " + strconv.FormatInt(int64(game.Halfmove), 10) + "\n"
	str += "Fullmove: " + strconv.FormatInt(int64(game.Fullmove), 10)
	return str
}

func DebugMoves(board []uint8, depth int, color uint8, enpassant []uint8) string {
	var list = Generate(board, color, enpassant)

	if color == COLOR_WHITE {
		color = COLOR_BLACK
	} else {
		color = COLOR_WHITE
	}

	var str = ""
	for _, move := range list {
		board = MakeMove(board, move)
		str += DebugBoard(board) + "\n"
		str += DebugMove(move) + "\n"
		// again?
		if depth > 1 {
			str += DebugMoves(board, depth-1, color, enpassant)
		}
		board = UndoMove(board, move)
	}

	return str
}

func DebugPos(i uint8) string {
	if i > 63 {
		return "-"
	}
	var row = i / 8
	var col = i - row*8
	return string(File[col]) + string(Rank[7-row])
}

func DebugMove(move Move) string {
	if move.Prev != EMPTY {
		return PieceName(move.Piece) + "x" + PieceName(move.Prev) + DebugPos(move.To)
	}
	return PieceName(move.Piece) + DebugPos(move.To)
}

func DebugHash(board []uint8) int {
	var h int = 0
	for _, v := range board {
		h += int(v)
	}
	return h
}

func DebugPerft(depth int) string {
	var result PerftResult = Perft(depth)
	var str = ""
	str += "Perft(" + strconv.Itoa(depth) + "):\n"
	str += "     Nodes: " + strconv.Itoa(result.Nodes) + "\n"
	str += "  Captures: " + strconv.Itoa(result.Captures) + "\n"
	return str
}

func DebugPerftDivide(depth int) string {
	var results []PerftDivideResult = PerftDivide(depth)
	var total int = 0
	var str = ""
	str += "Perft Divide(" + strconv.Itoa(depth) + "):\n"
	for _, result := range results {
		total += result.Nodes
		str += DebugMove(result.Move) + " : " + strconv.FormatInt(int64(result.Nodes), 10) + "\n"
	}
	str += "Total: " + strconv.FormatInt(int64(total), 10) + "\n"
	return str
}
