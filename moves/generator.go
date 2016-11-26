package moves

const INVALID_MOVE int16 = -1

type Move struct {
  Piece   int16
  Prev    int16
  From    int16
  To      int16
  Capture bool
}

func pos(i int16) (row int16, col int16) {
  row = i / 8
  col = i - row * 8
  return
}

func idx(row int16, col int16) int16 {
  if (row < 0 || row > 7) {
    return INVALID_MOVE
  }
  if (col < 0 || col > 7) {
    return INVALID_MOVE
  }
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

    // from
    from = int16(i)

    // col & row
    row, col = pos(int16(i))

    ////////////
    /// PAWN ///
    ////////////
    if piece == PAWN | color {
      var capture = false
      // move one up, must be clear
      move = field(board, row + 1 * dir, col)
      if move == EMPTY {
        to = idx(row + 1 * dir, col)
        moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
      }
      // move two up, if in starting position, next two fields are clear
      if move == EMPTY && (color == COLOR_WHITE && row == 6 || color == COLOR_BLACK && row == 1) {
        move = field(board, row + 2 * dir, col)
        if move == EMPTY {
          to = idx(row + 2 * dir, col)
          moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
        }
      }
      // attack left/up if enemy piece is there
      move = field(board, row + 1 * dir, col - 1)
      capture = move & enemy == enemy
      //fmt.Println("from row:", row, "col:", col, "to row:", row + 1 * dir, "col:", col - 1, "capture:", capture, "move:", move)
      if move > -1 && capture {
        to = idx(row + 1 * dir, col - 1)
        moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
      }
      // attack right/up if enemy piece is there
      move = field(board, row + 1 * dir, col + 1)
      capture = move & enemy == enemy
      if move > -1 && capture {
        to = idx(row + 1 * dir, col + 1)
        moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
      }
    }

    //////////////
    /// KNIGHT ///
    //////////////
    if piece == KNIGHT | color {
      var capture = false
      // Quadrant 1
      // -2, +1
      move = field(board, row - 2, col + 1)
      to = idx(row - 2, col + 1)
      capture = move & enemy == enemy
      if to > -1 && (move == EMPTY || capture) {
        moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
      }
      // -1, +2
      move = field(board, row - 1, col + 2)
      to = idx(row - 1, col + 2)
      capture = move & enemy == enemy
      if to > -1 && (move == EMPTY || capture) {
        moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
      }
      // Quadrant 2
      // +2, +1
      move = field(board, row + 2, col + 1)
      to = idx(row + 2, col + 1)
      capture = move & enemy == enemy
      if to > -1 && (move == EMPTY || capture) {
        moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
      }
      // +1, +2
      move = field(board, row + 1, col + 2)
      to = idx(row + 1, col + 2)
      capture = move & enemy == enemy
      if to > -1 && (move == EMPTY || capture) {
        moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
      }
      // Quadrant 3
      // +2, -1
      move = field(board, row + 2, col - 1)
      to = idx(row + 2, col - 1)
      capture = move & enemy == enemy
      if to > -1 && (move == EMPTY || capture) {
        moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
      }
      // +1, -2
      move = field(board, row + 1, col - 2)
      to = idx(row + 1, col - 2)
      capture = move & enemy == enemy
      if to > -1 && (move == EMPTY || capture) {
        moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
      }
      // Quadrant 4
      // -2, +1
      move = field(board, row - 2, col - 1)
      to = idx(row - 2, col - 1)
      capture = move & enemy == enemy
      if to > -1 && (move == EMPTY || capture) {
        moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
      }
      // -1, +2
      move = field(board, row - 1, col - 2)
      to = idx(row - 1, col - 2)
      capture = move & enemy == enemy
      if to > -1 && (move == EMPTY || capture) {
        moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
      }
    }

    //////////////
    /// BISHOP ///
    //////////////
    if piece == BISHOP | color {
      var capture = false
      // Diagonale 1 → -row,+col
      // TODO limit j to board limits
      for j := 1; j <= 7; j += 1 {
        var r int16 = int16(j * -1)
        var c int16 = int16(j)
        move = field(board, row + r, col + c)
        to = idx(row + r, col + c)
        capture = move & enemy == enemy
        if to > -1 && (move == EMPTY || capture) {
          moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
          if capture {
            break
          }
        } else {
          break
        }
      }
      // Diagonale 2 → +row,+col
      // TODO limit j to board limits
      for j := 1; j <= 7; j += 1 {
        var r int16 = int16(j)
        var c int16 = int16(j)
        move = field(board, row + r, col + c)
        to = idx(row + r, col + c)
        capture = move & enemy == enemy
        if to > -1 && (move == EMPTY || capture) {
          moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
          if capture {
            break
          }
        } else {
          break
        }
      }
      // Diagonale 3 → +row,-col
      // TODO limit j to board limits
      for j := 1; j <= 7; j += 1 {
        var r int16 = int16(j)
        var c int16 = int16(j * -1)
        move = field(board, row + r, col + c)
        to = idx(row + r, col + c)
        capture = move & enemy == enemy
        if to > -1 && (move == EMPTY || capture) {
          moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
          if capture {
            break
          }
        } else {
          break
        }
      }
      // Diagonale 4 → -row,-col
      // TODO limit j to board limits
      for j := 1; j <= 7; j += 1 {
        var r int16 = int16(j * -1)
        var c int16 = int16(j * -1)
        move = field(board, row + r, col + c)
        to = idx(row + r, col + c)
        capture = move & enemy == enemy
        if to > -1 && (move == EMPTY || capture) {
          moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
          if capture {
            break
          }
        } else {
          break
        }
      }
    }

    ////////////
    /// ROOK ///
    ////////////
    if piece == ROOK | color {
      var capture = false
      // -rank
      for r := row - 1; r >= 0; r -= 1 {
        move = field(board, r, col)
        to = idx(r, col)
        capture = move & enemy == enemy
        if to > -1 && (move == EMPTY || capture) {
          moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
          if capture {
            break
          }
        } else {
          break
        }
      }
      // +rank
      for r := row + 1; r <= 7; r += 1 {
        move = field(board, r, col)
        to = idx(r, col)
        capture = move & enemy == enemy
        if to > -1 && (move == EMPTY || capture) {
          moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
          if capture {
            break
          }
        } else {
          break
        }
      }
      // -file
      for c := col - 1; c >= 0; c -= 1 {
        move = field(board, row, c)
        to = idx(row, c)
        capture = move & enemy == enemy
        if to > -1 && (move == EMPTY || capture) {
          moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
          if capture {
            break
          }
        } else {
          break
        }
      }
      // +file
      for c := col + 1; c <= 7; c += 1 {
        move = field(board, row, c)
        to = idx(row, c)
        capture = move & enemy == enemy
        if to > -1 && (move == EMPTY || capture) {
          moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
          if capture {
            break
          }
        } else {
          break
        }
      }
    }

    /////////////
    /// QUEEN ///
    /////////////
    if piece == QUEEN | color {
      var capture = false
      // Diagonale 1 → -row,+col
      // TODO limit j to board limits
      for j := 1; j <= 7; j += 1 {
        var r int16 = int16(j * -1)
        var c int16 = int16(j)
        move = field(board, row + r, col + c)
        to = idx(row + r, col + c)
        capture = move & enemy == enemy
        if to > -1 && (move == EMPTY || capture) {
          moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
          if capture {
            break
          }
        } else {
          break
        }
      }
      // Diagonale 2 → +row,+col
      // TODO limit j to board limits
      for j := 1; j <= 7; j += 1 {
        var r int16 = int16(j)
        var c int16 = int16(j)
        move = field(board, row + r, col + c)
        to = idx(row + r, col + c)
        capture = move & enemy == enemy
        if to > -1 && (move == EMPTY || capture) {
          moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
          if capture {
            break
          }
        } else {
          break
        }
      }
      // Diagonale 3 → +row,-col
      // TODO limit j to board limits
      for j := 1; j <= 7; j += 1 {
        var r int16 = int16(j)
        var c int16 = int16(j * -1)
        move = field(board, row + r, col + c)
        to = idx(row + r, col + c)
        capture = move & enemy == enemy
        if to > -1 && (move == EMPTY || capture) {
          moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
          if capture {
            break
          }
        } else {
          break
        }
      }
      // Diagonale 4 → -row,-col
      // TODO limit j to board limits
      for j := 1; j <= 7; j += 1 {
        var r int16 = int16(j * -1)
        var c int16 = int16(j * -1)
        move = field(board, row + r, col + c)
        to = idx(row + r, col + c)
        capture = move & enemy == enemy
        if to > -1 && (move == EMPTY || capture) {
          moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
          if capture {
            break
          }
        } else {
          break
        }
      }
      // -rank
      for r := row - 1; r >= 0; r -= 1 {
        move = field(board, r, col)
        to = idx(r, col)
        capture = move & enemy == enemy
        if to > -1 && (move == EMPTY || capture) {
          moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
          if capture {
            break
          }
        } else {
          break
        }
      }
      // +rank
      for r := row + 1; r <= 7; r += 1 {
        move = field(board, r, col)
        to = idx(r, col)
        capture = move & enemy == enemy
        if to > -1 && (move == EMPTY || capture) {
          moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
          if capture {
            break
          }
        } else {
          break
        }
      }
      // -file
      for c := col - 1; c >= 0; c -= 1 {
        move = field(board, row, c)
        to = idx(row, c)
        capture = move & enemy == enemy
        if to > -1 && (move == EMPTY || capture) {
          moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
          if capture {
            break
          }
        } else {
          break
        }
      }
      // +file
      for c := col + 1; c <= 7; c += 1 {
        move = field(board, row, c)
        to = idx(row, c)
        capture = move & enemy == enemy
        if to > -1 && (move == EMPTY || capture) {
          moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
          if capture {
            break
          }
        } else {
          break
        }
      }
    }

    ////////////
    /// KING ///
    ////////////
    if piece == KING | color {
      var capture = false
      for r := -1; r <= 1; r += 2 {
        for c := -1; c <= 1; c += 2 {
          move = field(board, row + int16(r), col + int16(c))
          to = idx(row + int16(r), col + int16(c))
          capture = move & enemy == enemy
          if to > -1 && (move == EMPTY || capture) {
            moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to, Capture: capture})
          }
        }
      }
    }

  }

  return moves
}
