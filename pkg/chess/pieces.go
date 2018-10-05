package chess

const Rank = "12345678"
const File = "abcdefgh"

func PieceCode(piece string) uint8 {
	switch piece {
	// white
	case "K":
		return WhiteKing
	case "Q":
		return WhiteQueen
	case "R":
		return WhiteRook
	case "B":
		return WhiteBishop
	case "N":
		return WhiteKnight
	case "P":
		return WhitePawn
		// black
	case "k":
		return BlackKing
	case "q":
		return BlackQueen
	case "r":
		return BlackRook
	case "b":
		return BlackBishop
	case "n":
		return BlackKnight
	case "p":
		return BlackPawn
	}
	return Empty
}
