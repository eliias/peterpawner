package moves

var Start = [64]int16{
  B_ROOK, B_KNIGHT, B_BISHOP, B_QUEEN, B_KING, B_BISHOP, B_KNIGHT, B_ROOK,
  B_PAWN, B_PAWN,   B_PAWN,   B_PAWN,  B_PAWN, B_PAWN,   B_PAWN,   B_PAWN,
  EMPTY,  EMPTY,    EMPTY,    EMPTY,   EMPTY,  EMPTY,    EMPTY,    EMPTY,
  EMPTY,  EMPTY,    EMPTY,    EMPTY,   EMPTY,  EMPTY,    EMPTY,    EMPTY,
  EMPTY,  EMPTY,    EMPTY,    EMPTY,   EMPTY,  EMPTY,    EMPTY,    EMPTY,
  EMPTY,  EMPTY,    EMPTY,    EMPTY,   EMPTY,  EMPTY,    EMPTY,    EMPTY,
  W_PAWN, W_PAWN,   W_PAWN,   W_PAWN,  W_PAWN, W_PAWN,   W_PAWN,   W_PAWN,
  W_ROOK, W_KNIGHT, W_BISHOP, W_QUEEN, W_KING, W_BISHOP, W_KNIGHT, W_ROOK}

func Add(board [64]int16, move Move) [64]int16 {
  board[move.from] = EMPTY
  board[move.to] = move.piece
  return board
}

func Remove(board [64]int16, move Move) [64]int16 {
  board[move.to] = EMPTY
  board[move.from] = move.piece
  return board
}
/*
func Load(fen string) [64]int16 {
  return []int16{}
}

func Save(board [64]int16) string {
  return ""
}
*/
func Perft(depth int) int {
  var board = Start
  var nodes int = 0

  var moves = Generate(board, COLOR_WHITE)

  if depth == 1 {
    return len(moves)
  }

  for _, move := range moves {
    Add(board, move)
    nodes += Perft(depth - 1)
    Remove(board, move)
  }

  return nodes
}
