package debug

import (
  "github.com/eliias/peterpawner/moves"
)

var rank = "12345678"
var file = "abcdefgh"

func Board(board [64]int16) string {
  var str = ""
  var col int
  var row int = 0
  for i := 0; i < 100; i += 1 {
    if i % 10 == 0 && i > 0 {
      row += 1
      str += "\n"
    }
    col = i - row * 10

    if col == 0 && row == 0 || col == 9 && row == 0 || col == 0 && row == 9 || col == 9 && row == 9 {
      str += "+"
    } else if row == 0 && col > 0 && col < 9 {
      str += "-"
    } else if row == 9 && col > 0 && col < 9 {
      str += "-"
    } else if col == 0 || col == 9 {
      str += "|"
    } else {
      str += PieceName(board[col + 8 * (row - 1) - 1])
    }
  }

  return str
}

func Moves(board [64]int16, depth int, color int16) string {
  var list = moves.Generate(board, color)

  if color == moves.COLOR_WHITE {
    color = moves.COLOR_BLACK
  } else {
    color = moves.COLOR_WHITE
  }

  var str = ""
  for _, move := range list {
    board = moves.Add(board, move)
    str += Board(board) + "\n"
    str += Move(move) + "\n"
    // again?
    if (depth > 1) {
      str += Moves(board, depth - 1, color)
    }
    board = moves.Remove(board, move)
  }

  return str
}

func Pos(i int16) string {
  var row = i / 8
  var col = i - row * 8
  return string(file[col]) + string(rank[7 - row])
}

func Move(move moves.Move) string {
  if move.Capture {
    return PieceName(move.Piece) + "x" + PieceName(move.Prev) + Pos(move.To)
  }
  return PieceName(move.Piece) + Pos(move.To)
}

func Hash(board [64]int16) int16 {
  var h int16 = 0
  for _, v := range board {
    h += v
  }
  return h
}
