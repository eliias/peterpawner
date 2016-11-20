package debug

import "github.com/eliias/peterpawner/moves"

func PieceName(piece int16) string {
  switch piece {
  // white
  case moves.W_KING: return "K"
  case moves.W_QUEEN: return "Q"
  case moves.W_ROOK: return "R"
  case moves.W_BISHOP: return "B"
  case moves.W_KNIGHT: return "N"
  case moves.W_PAWN: return "P"
  // black
  case moves.B_KING: return "k"
  case moves.B_QUEEN: return "q"
  case moves.B_ROOK: return "r"
  case moves.B_BISHOP: return "b"
  case moves.B_KNIGHT: return "n"
  case moves.B_PAWN: return "p"
  }
  return "."
}
