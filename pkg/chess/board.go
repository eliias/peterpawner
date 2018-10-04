package chess

import (
	"strconv"
	"strings"
)

var Start = []uint8{
	B_ROOK, B_KNIGHT, B_BISHOP, B_QUEEN, B_KING, B_BISHOP, B_KNIGHT, B_ROOK,
	B_PAWN, B_PAWN, B_PAWN, B_PAWN, B_PAWN, B_PAWN, B_PAWN, B_PAWN,
	EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY,
	EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY,
	EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY,
	EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY,
	W_PAWN, W_PAWN, W_PAWN, W_PAWN, W_PAWN, W_PAWN, W_PAWN, W_PAWN,
	W_ROOK, W_KNIGHT, W_BISHOP, W_QUEEN, W_KING, W_BISHOP, W_KNIGHT, W_ROOK}

var Test = []uint8{
	B_ROOK, B_KNIGHT, B_BISHOP, B_QUEEN, B_KING, B_BISHOP, B_KNIGHT, B_ROOK,
	B_PAWN, B_PAWN, B_PAWN, B_PAWN, B_PAWN, B_PAWN, B_PAWN, B_PAWN,
	EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY,
	EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY,
	EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY,
	EMPTY, EMPTY, EMPTY, W_PAWN, EMPTY, EMPTY, EMPTY, EMPTY,
	W_PAWN, W_PAWN, W_PAWN, EMPTY, W_PAWN, W_PAWN, W_PAWN, W_PAWN,
	W_ROOK, W_KNIGHT, W_BISHOP, W_QUEEN, W_KING, W_BISHOP, W_KNIGHT, W_ROOK}

type Game struct {
	Board                []uint8
	Color                uint8
	WhiteCastleKingSide  bool
	WhiteCastleQueenSide bool
	BlackCastleKingSide  bool
	BlackCastleQueenSide bool
	EnPassant            uint8
	Halfmove             int
	Fullmove             int
}

func MakeMove(board []uint8, move *Move) []uint8 {
	board[move.From] = EMPTY
	board[move.To] = move.Piece
	return board
}

func UndoMove(board []uint8, move *Move) []uint8 {
	board[move.To] = move.Prev
	board[move.From] = move.Piece
	return board
}

func Load(fen string) Game {
	// board
	var board []uint8

	// parts
	var parts []string = strings.Split(fen, " ")

	// pieces
	var ranks []string = strings.Split(parts[0], "/")
	for _, rank := range ranks {
		for _, piece := range rank {
			// numbers are numbers of pawns in rank
			if l, err := strconv.Atoi(string(piece)); err == nil {
				for i := 0; i < l; i += 1 {
					board = append(board, EMPTY)
				}
			} else {
				board = append(board, PieceCode(string(piece)))
			}
		}
	}

	// active color
	var color = ColorCode(parts[1])

	// castling
	var castling = parts[2]
	var whiteCastleKingSide = strings.Contains(castling, "K")
	var whiteCastleQueenSide = strings.Contains(castling, "Q")
	var blackCastleKingSide = strings.Contains(castling, "k")
	var blackCastleQueenSide = strings.Contains(castling, "q")

	// en passant
	var enPassant uint8
	if parts[3] != "-" {
		enPassant = PosCode(parts[3])
	} else {
		enPassant = INVALID_MOVE
	}

	// half move
	var halfmove int
	if hm, err := strconv.Atoi(string(parts[4])); err == nil {
		halfmove = hm
	}

	// full move
	var fullmove int
	if fm, err := strconv.Atoi(string(parts[5])); err == nil {
		fullmove = fm
	}

	// game
	return Game{
		Board:                board,
		Color:                color,
		WhiteCastleKingSide:  whiteCastleKingSide,
		WhiteCastleQueenSide: whiteCastleQueenSide,
		BlackCastleKingSide:  blackCastleKingSide,
		BlackCastleQueenSide: blackCastleQueenSide,
		EnPassant:            enPassant,
		Halfmove:             halfmove,
		Fullmove:             fullmove}
}

func Save(game Game) string {
	var str = ""
	var cnt = 0

	// board
	for i := 0; i < 64; i += 1 {
		// piece
		var piece = game.Board[uint8(i)]

		// empty
		if piece == EMPTY {
			// found an empty field
			cnt += 1
		}

		// regular piece
		if piece != EMPTY {
			if cnt > 0 {
				str += strconv.FormatInt(int64(cnt), 10)
			}
			str += PieceName(piece)
			cnt = 0
		}

		// new rank
		if (i+1)%8 == 0 {
			if cnt > 0 {
				str += strconv.FormatInt(int64(cnt), 10)
			}
			if i < 63 {
				str += "/"
			}
			cnt = 0
		}
	}

	// active color
	str += " " + ColorName(game.Color)

	// castling
	var castling = ""
	if game.WhiteCastleKingSide {
		castling += "K"
	}
	if game.WhiteCastleQueenSide {
		castling += "Q"
	}
	if game.BlackCastleKingSide {
		castling += "k"
	}
	if game.BlackCastleQueenSide {
		castling += "q"
	}
	if len(castling) == 0 {
		castling = "-"
	}
	str += " " + castling

	// en passant
	var enPassant = ""
	if game.EnPassant != INVALID_MOVE {
		enPassant = DebugPos(game.EnPassant)
	} else {
		enPassant = "-"
	}
	str += " " + enPassant

	// halfmove
	str += " " + strconv.FormatInt(int64(game.Halfmove), 10)

	// fullmove
	str += " " + strconv.FormatInt(int64(game.Fullmove), 10)

	return str
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
	Move  *Move
	Nodes int
}

func stats(moves []*Move) PerftResult {
	var result = PerftResult{Nodes: len(moves)}
	for _, move := range moves {
		if move.Prev != EMPTY {
			result.Captures += 1
		}
	}
	return result
}

func perft(board []uint8, depth int, color uint8, enPassantFields []uint8) PerftResult {
	var nodes = 0
	var captures = 0
	var result PerftResult
	var enPassant []uint8
	var moves = Generate(board, color, enPassantFields)
	result = stats(moves)

	if color == COLOR_WHITE {
		color = COLOR_BLACK
	} else {
		color = COLOR_WHITE
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
		MakeMove(board, move)

		// next level
		result = perft(board, depth-1, color, enPassant)

		// stats
		nodes += result.Nodes
		captures += result.Captures

		// undo move
		UndoMove(board, move)
	}

	return PerftResult{Nodes: nodes, Captures: captures}
}

func Perft(depth int) PerftResult {
	return perft(Start, depth, COLOR_WHITE, []uint8{})
}

func perftDivide(board []uint8, depth int, color uint8, enPassantFields []uint8) []PerftDivideResult {
	var moves = Generate(board, color, enPassantFields)
	var result PerftResult
	var divides []PerftDivideResult
	var enPassant []uint8

	if color == COLOR_WHITE {
		color = COLOR_BLACK
	} else {
		color = COLOR_WHITE
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
		MakeMove(board, move)

		// next level
		result = perft(board, depth-1, color, enPassant)

		// stats
		divide = PerftDivideResult{Move: move, Nodes: result.Nodes}
		divides = append(divides, divide)

		// undo move
		UndoMove(board, move)
	}

	return divides
}

func PerftDivide(depth int) []PerftDivideResult {
	return perftDivide(Start, depth, COLOR_WHITE, []uint8{})
}
