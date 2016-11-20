package moves

const INVALID_MOVE int16 = 0

type Move struct {
  piece int16
  from  int16
  to    int16
}

func pos(i int16) (row int16, col int16) {
  row = i / 8
  col = i - row * 8
  return
}

func idx(row int16, col int16) int16 {
  return row * 8 + col
}

func field(board [64]int16, row int16, col int16) int16 {
  if (row < 0 || row > 7) {
    return INVALID_MOVE
  }
  if (col < 0 || col > 7) {
    return INVALID_MOVE
  }
  var i = idx(row, col)
  return board[i]
}

func Generate(board [64]int16, color int16) []Move {
  var dir int16
  var col int16
  var row int16
  var from int16
  var to int16
  var move int16
  var enemy int16
  var moves []Move

  if (color == COLOR_WHITE) {
    dir = -1
    enemy = COLOR_BLACK
  } else {
    dir = +1
    enemy = COLOR_WHITE
  }
  for i, piece := range board {
    // check if empty
    if piece == EMPTY {
      continue
    }

    // TODO exclude enemy pieces from moves

    // from
    from = int16(i)

    // col & row
    row, col = pos(int16(i))

    ////////////
    /// PAWN ///
    ////////////
    if piece == PAWN | color {
      // move one up, must be clear
      move = field(board, row + 1 * dir, col)

      if move == EMPTY {
        to = idx(row + 1 * dir, col)
        moves = append(moves, Move{piece:PAWN | color, from: from, to: to})
      }
      // move two up, if in starting position, next two fields are clear
      if move == EMPTY && (color == COLOR_WHITE && row == 6 || color == COLOR_BLACK && row == 1) {
        move = field(board, row + 2 * dir, col)
        if move == EMPTY {
          to = idx(row + 2 * dir, col)
          moves = append(moves, Move{piece:PAWN | color, from: from, to: to})
        }
      }
      // attack left/up if enemy piece is there
      move = field(board, row + 1 * dir, col - 1 * dir)
      if move & enemy == enemy {
        to = idx(row + 1 * dir, col - 1)
        moves = append(moves, Move{piece:PAWN | color, from: from, to: to})
      }
      // attack right/up if enemy piece is there
      move = field(board, row + 1 * dir, col + 1 * dir)
      if move & enemy == enemy {
        to = idx(row + 1 * dir, col + 1)
        moves = append(moves, Move{piece:PAWN | color, from: from, to: to})
      }
    }

    //////////////
    /// KNIGHT ///
    //////////////
    if piece == KNIGHT | color {
      //
    }
  }

  return moves
}
