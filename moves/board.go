package moves

import "fmt"

var Start = [64]int16{
  B_ROOK, B_KNIGHT, B_BISHOP, B_QUEEN, B_KING, B_BISHOP, B_KNIGHT, B_ROOK,
  B_PAWN, B_PAWN,   B_PAWN,   B_PAWN,  B_PAWN, B_PAWN,   B_PAWN,   B_PAWN,
  EMPTY,  EMPTY,    EMPTY,    EMPTY,   EMPTY,  EMPTY,    EMPTY,    EMPTY,
  EMPTY,  EMPTY,    EMPTY,    EMPTY,   EMPTY,  EMPTY,    EMPTY,    EMPTY,
  EMPTY,  EMPTY,    EMPTY,    EMPTY,   EMPTY,  EMPTY,    EMPTY,    EMPTY,
  EMPTY,  EMPTY,    EMPTY,    EMPTY,   EMPTY,  EMPTY,    EMPTY,    EMPTY,
  W_PAWN, W_PAWN,   W_PAWN,   W_PAWN,  W_PAWN, W_PAWN,   W_PAWN,   W_PAWN,
  W_ROOK, W_KNIGHT, W_BISHOP, W_QUEEN, W_KING, W_BISHOP, W_KNIGHT, W_ROOK}

var Test = [64]int16{
  B_ROOK, B_KNIGHT, B_BISHOP, B_QUEEN, B_KING, B_BISHOP, B_KNIGHT, B_ROOK,
  B_PAWN, B_PAWN,   B_PAWN,   B_PAWN,  B_PAWN, B_PAWN,   B_PAWN,   B_PAWN,
  EMPTY,  EMPTY,    EMPTY,    EMPTY,   EMPTY,  EMPTY,    EMPTY,    EMPTY,
  EMPTY,  EMPTY,    EMPTY,    EMPTY,   EMPTY,  EMPTY,    EMPTY,    EMPTY,
  EMPTY,  EMPTY,    EMPTY,    EMPTY,   EMPTY,  EMPTY,    EMPTY,    EMPTY,
  EMPTY,  EMPTY,    EMPTY,    W_PAWN,  EMPTY,  EMPTY,    EMPTY,    EMPTY,
  W_PAWN, W_PAWN,   W_PAWN,   EMPTY,   W_PAWN, W_PAWN,   W_PAWN,   W_PAWN,
  W_ROOK, W_KNIGHT, W_BISHOP, W_QUEEN, W_KING, W_BISHOP, W_KNIGHT, W_ROOK}

func Add(board [64]int16, move Move) [64]int16 {
  board[move.From] = EMPTY
  board[move.To] = move.Piece
  return board
}

func Remove(board [64]int16, move Move) [64]int16 {
  board[move.To] = move.Prev
  board[move.From] = move.Piece
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
func  Perft(board [64]int16, depth int, color int16) int {
  var nodes int = 0
  var captures int = 0

  var moves = Generate(board, color)

  // stats
  for _, move := range moves {
    if move.Capture {
      fmt.Println("capture", move.Capture)
      captures += 1
    }
  }

  if color == COLOR_WHITE {
    color = COLOR_BLACK
  } else {
    color = COLOR_WHITE
  }

  if depth == 1 {
    return len(moves)
  }

  for _, move := range moves {
    // make move
    Add(board, move)

    // look deeper
    nodes += Perft(board, depth - 1, color)

    // undo move
    Remove(board, move)
  }

  return nodes
}
