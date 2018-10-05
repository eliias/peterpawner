package chess

var Start = []uint8{
	Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
	Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
	Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
	WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn,
	WhiteRook, WhiteKnight, WhiteBishop, WhiteQueen, WhiteKing, WhiteBishop, WhiteKnight, WhiteRook}

var Test = []uint8{
	BlackRook, BlackKnight, BlackBishop, BlackQueen, BlackKing, BlackBishop, BlackKnight, BlackRook,
	BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn,
	Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
	Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
	Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
	Empty, Empty, Empty, WhitePawn, Empty, Empty, Empty, Empty,
	WhitePawn, WhitePawn, WhitePawn, Empty, WhitePawn, WhitePawn, WhitePawn, WhitePawn,
	WhiteRook, WhiteKnight, WhiteBishop, WhiteQueen, WhiteKing, WhiteBishop, WhiteKnight, WhiteRook}

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
	board[move.From] = Empty
	board[move.To] = move.Piece
	return board
}

func UndoMove(board []uint8, move *Move) []uint8 {
	board[move.To] = move.Prev
	board[move.From] = move.Piece
	return board
}

type Move struct {
	Piece     uint8
	Prev      uint8
	From      uint8
	To        uint8
	EnPassant uint8
}
