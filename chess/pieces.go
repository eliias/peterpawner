package chess

import "strings"

const EMPTY uint8 = 0

const KING uint8 = 1 << 0
const QUEEN uint8 = 1 << 1
const ROOK uint8 = 1 << 2
const BISHOP uint8 = 1 << 3
const KNIGHT uint8 = 1 << 4
const PAWN uint8 = 1 << 5

const COLOR_WHITE uint8 = 1 << 6
const COLOR_BLACK uint8 = 1 << 7

const W_KING uint8 = KING | COLOR_WHITE
const W_QUEEN uint8 = QUEEN | COLOR_WHITE
const W_ROOK uint8 = ROOK | COLOR_WHITE
const W_BISHOP uint8 = BISHOP | COLOR_WHITE
const W_KNIGHT uint8 = KNIGHT | COLOR_WHITE
const W_PAWN uint8 = PAWN | COLOR_WHITE

const B_KING uint8 = KING | COLOR_BLACK
const B_QUEEN uint8 = QUEEN | COLOR_BLACK
const B_ROOK uint8 = ROOK | COLOR_BLACK
const B_BISHOP uint8 = BISHOP | COLOR_BLACK
const B_KNIGHT uint8 = KNIGHT | COLOR_BLACK
const B_PAWN uint8 = PAWN | COLOR_BLACK

func PieceCode(piece string) uint8 {
  switch piece {
  // white
  case "K": return W_KING
  case "Q": return W_QUEEN
  case "R": return W_ROOK
  case "B": return W_BISHOP
  case "N": return W_KNIGHT
  case "P": return W_PAWN
  // black
  case "k": return B_KING
  case "q": return B_QUEEN
  case "r": return B_ROOK
  case "b": return B_BISHOP
  case "n": return B_KNIGHT
  case "p": return B_PAWN
  }
  return EMPTY
}

func PosCode(pos string) uint8 {
  var parts = strings.Split(pos, "")
  var file = parts[0]
  var rank = parts[1]
  var r = strings.Index(Rank, rank)
  var f = strings.Index(File, file)
  var i = idx(uint8(7 - r), uint8(f))
  return i
}

func ColorCode(color string) uint8 {
  if color == "w" {
    return COLOR_WHITE
  } else {
    return COLOR_BLACK
  }
}
