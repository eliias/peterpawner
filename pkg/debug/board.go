package debug

import "github.com/eliias/peterpawner/pkg/chess"

const Rank = "12345678"
const File = "abcdefgh"

func PieceName(piece uint8) string {
	switch piece {
	// white
	case chess.WhiteKing:
		return "K"
	case chess.WhiteQueen:
		return "Q"
	case chess.WhiteRook:
		return "R"
	case chess.WhiteBishop:
		return "B"
	case chess.WhiteKnight:
		return "N"
	case chess.WhitePawn:
		return "P"
		// black
	case chess.BlackKing:
		return "k"
	case chess.BlackQueen:
		return "q"
	case chess.BlackRook:
		return "r"
	case chess.BlackBishop:
		return "b"
	case chess.BlackKnight:
		return "n"
	case chess.BlackPawn:
		return "p"
	}
	return "."
}

func ColorName(color uint8) string {
	if color == chess.ColorWhite {
		return "w"
	}
	return "b"
}
