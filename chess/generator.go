package chess

const INVALID_MOVE uint8 = 255

type Move struct {
	Piece uint8
	Prev  uint8
	From  uint8
	To    uint8
}

func pos(i uint8) (row uint8, col uint8) {
	row = i / 8
	col = i - row * 8
	return
}

func idx(row uint8, col uint8) uint8 {
	if row < 0 || row > 7 {
		return INVALID_MOVE
	}
	if col < 0 || col > 7 {
		return INVALID_MOVE
	}
	return row * 8 + col
}

func field(board []uint8, row uint8, col uint8) uint8 {
	if row < 0 || row > 7 {
		return INVALID_MOVE
	}
	if col < 0 || col > 7 {
		return INVALID_MOVE
	}
	var i = idx(row, col)
	return board[i]
}

func pawnAttacks(board []uint8, color uint8, enemy uint8, row uint8, col uint8) []uint8 {
	var fields []uint8
	var move uint8
	var trow uint8
	var to uint8
	var capture bool
	var empty bool

	// attack left/up if enemy piece is there
	if color == COLOR_BLACK {
		trow = row + 1
	} else {
		trow = row - 1
	}
	move = field(board, trow, col - 1)
	capture = move & enemy == enemy
	empty = move == EMPTY
	if move != INVALID_MOVE && (capture || empty) {
		to = idx(trow, col - 1)
		fields = append(fields, to)
	}
	// attack right/up if enemy piece is there
	move = field(board, trow, col + 1)
	capture = move & enemy == enemy
	empty = move == EMPTY
	if move != INVALID_MOVE && (capture || empty) {
		to = idx(trow, col + 1)
		fields = append(fields, to)
	}
	return fields
}

func pawn(board []uint8, piece uint8, color uint8, enemy uint8, row uint8, col uint8) []Move {
	var moves []Move
	var capture = false
	var move uint8
	var from uint8 = idx(row, col)
	var to uint8
	var trow uint8
	// move one up, target must be empty
	if color == COLOR_BLACK {
		trow = row + 1
	} else {
		trow = row - 1
	}
	move = field(board, trow, col)
	if move == EMPTY {
		to = idx(trow, col)
		moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	// move two up, if in starting position, next two fields are clear
	if move == EMPTY && (color == COLOR_WHITE && row == 6 || color == COLOR_BLACK && row == 1) {
		if color == COLOR_BLACK {
			trow = row + 2
		} else {
			trow = row - 2
		}
		move = field(board, trow, col)
		if move == EMPTY {
			to = idx(trow, col)
			moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
		}
	}
	// attack left/up if enemy piece is there
	if color == COLOR_BLACK {
		trow = row + 1
	} else {
		trow = row - 1
	}
	move = field(board, trow, col - 1)
	capture = move & enemy == enemy
	if move != INVALID_MOVE && capture {
		to = idx(trow, col - 1)
		moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	// attack right/up if enemy piece is there
	move = field(board, trow, col + 1)
	capture = move & enemy == enemy
	if move != INVALID_MOVE && capture {
		to = idx(trow, col + 1)
		moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	return moves
}

func knightAttacks(board []uint8, enemy uint8, row uint8, col uint8) []uint8 {
	var fields []uint8
	var capture bool
	var empty bool
	var move uint8
	var to uint8
	// Quadrant 1
	// -2, +1
	move = field(board, row - 2, col + 1)
	to = idx(row - 2, col + 1)
	capture = move & enemy == enemy
	empty = move == EMPTY
	if to != INVALID_MOVE && (empty || capture) {
		fields = append(fields, to)
	}
	// -1, +2
	move = field(board, row - 1, col + 2)
	to = idx(row - 1, col + 2)
	capture = move & enemy == enemy
	empty = move == EMPTY
	if to != INVALID_MOVE && (empty || capture) {
		fields = append(fields, to)
	}
	// Quadrant 2
	// +2, +1
	move = field(board, row + 2, col + 1)
	to = idx(row + 2, col + 1)
	capture = move & enemy == enemy
	empty = move == EMPTY
	if to != INVALID_MOVE && (empty || capture) {
		fields = append(fields, to)
	}
	// +1, +2
	move = field(board, row + 1, col + 2)
	to = idx(row + 1, col + 2)
	capture = move & enemy == enemy
	empty = move == EMPTY
	if to != INVALID_MOVE && (empty || capture) {
		fields = append(fields, to)
	}
	// Quadrant 3
	// +2, -1
	move = field(board, row + 2, col - 1)
	to = idx(row + 2, col - 1)
	capture = move & enemy == enemy
	empty = move == EMPTY
	if to != INVALID_MOVE && (empty || capture) {
		fields = append(fields, to)
	}
	// +1, -2
	move = field(board, row + 1, col - 2)
	to = idx(row + 1, col - 2)
	capture = move & enemy == enemy
	empty = move == EMPTY
	if to != INVALID_MOVE && (empty || capture) {
		fields = append(fields, to)
	}
	// Quadrant 4
	// -2, +1
	move = field(board, row - 2, col - 1)
	to = idx(row - 2, col - 1)
	capture = move & enemy == enemy
	empty = move == EMPTY
	if to != INVALID_MOVE && (empty || capture) {
		fields = append(fields, to)
	}
	// -1, +2
	move = field(board, row - 1, col - 2)
	to = idx(row - 1, col - 2)
	capture = move & enemy == enemy
	empty = move == EMPTY
	if to != INVALID_MOVE && (empty || capture) {
		fields = append(fields, to)
	}
	return fields
}

func knight(board []uint8, piece uint8, enemy uint8, row uint8, col uint8) []Move {
	var moves []Move
	var capture = false
	var move uint8
	var from uint8 = idx(row, col)
	var to uint8
	// Quadrant 1
	// -2, +1
	move = field(board, row - 2, col + 1)
	to = idx(row - 2, col + 1)
	capture = move & enemy == enemy
	if to != INVALID_MOVE && (move == EMPTY || capture) {
		moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	// -1, +2
	move = field(board, row - 1, col + 2)
	to = idx(row - 1, col + 2)
	capture = move & enemy == enemy
	if to != INVALID_MOVE && (move == EMPTY || capture) {
		moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	// Quadrant 2
	// +2, +1
	move = field(board, row + 2, col + 1)
	to = idx(row + 2, col + 1)
	capture = move & enemy == enemy
	if to != INVALID_MOVE && (move == EMPTY || capture) {
		moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	// +1, +2
	move = field(board, row + 1, col + 2)
	to = idx(row + 1, col + 2)
	capture = move & enemy == enemy
	if to != INVALID_MOVE && (move == EMPTY || capture) {
		moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	// Quadrant 3
	// +2, -1
	move = field(board, row + 2, col - 1)
	to = idx(row + 2, col - 1)
	capture = move & enemy == enemy
	if to != INVALID_MOVE && (move == EMPTY || capture) {
		moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	// +1, -2
	move = field(board, row + 1, col - 2)
	to = idx(row + 1, col - 2)
	capture = move & enemy == enemy
	if to != INVALID_MOVE && (move == EMPTY || capture) {
		moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	// Quadrant 4
	// -2, +1
	move = field(board, row - 2, col - 1)
	to = idx(row - 2, col - 1)
	capture = move & enemy == enemy
	if to != INVALID_MOVE && (move == EMPTY || capture) {
		moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	// -1, +2
	move = field(board, row - 1, col - 2)
	to = idx(row - 1, col - 2)
	capture = move & enemy == enemy
	if to != INVALID_MOVE && (move == EMPTY || capture) {
		moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	return moves
}

func bishopAttacks(board []uint8, enemy uint8, row uint8, col uint8) []uint8 {
	var fields []uint8
	var capture bool
	var move uint8
	var to uint8
	// Diagonale 1 → -row,+col
	// TODO limit j to board limits
	for j := 1; j <= 7; j += 1 {
		var r uint8 = uint8(j * -1)
		var c uint8 = uint8(j)
		move = field(board, row + r, col + c)
		to = idx(row + r, col + c)
		capture = move & enemy == enemy
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			fields = append(fields, to)
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
		var r uint8 = uint8(j)
		var c uint8 = uint8(j)
		move = field(board, row + r, col + c)
		to = idx(row + r, col + c)
		capture = move & enemy == enemy
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			fields = append(fields, to)
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
		var r uint8 = uint8(j)
		var c uint8 = uint8(j * -1)
		move = field(board, row + r, col + c)
		to = idx(row + r, col + c)
		capture = move & enemy == enemy
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			fields = append(fields, to)
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
		var r uint8 = uint8(j * -1)
		var c uint8 = uint8(j * -1)
		move = field(board, row + r, col + c)
		to = idx(row + r, col + c)
		capture = move & enemy == enemy
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			fields = append(fields, to)
			if capture {
				break
			}
		} else {
			break
		}
	}
	return fields
}

func bishop(board []uint8, piece uint8, enemy uint8, row uint8, col uint8) []Move {
	var moves []Move
	var capture = false
	var move uint8
	var from uint8 = idx(row, col)
	var to uint8
	// Diagonale 1 → -row,+col
	// TODO limit j to board limits
	for j := 1; j <= 7; j += 1 {
		var r uint8 = uint8(j * -1)
		var c uint8 = uint8(j)
		move = field(board, row + r, col + c)
		to = idx(row + r, col + c)
		capture = move & enemy == enemy
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
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
		var r uint8 = uint8(j)
		var c uint8 = uint8(j)
		move = field(board, row + r, col + c)
		to = idx(row + r, col + c)
		capture = move & enemy == enemy
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
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
		var r uint8 = uint8(j)
		var c uint8 = uint8(j * -1)
		move = field(board, row + r, col + c)
		to = idx(row + r, col + c)
		capture = move & enemy == enemy
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
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
		var r uint8 = uint8(j * -1)
		var c uint8 = uint8(j * -1)
		move = field(board, row + r, col + c)
		to = idx(row + r, col + c)
		capture = move & enemy == enemy
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
			if capture {
				break
			}
		} else {
			break
		}
	}
	return moves
}

func rookAttacks(board []uint8, enemy uint8, row uint8, col uint8) []uint8 {
	var fields []uint8
	var capture bool
	var move uint8
	var to uint8
	// -rank
	for r := row - 1; r >= 0; r -= 1 {
		move = field(board, r, col)
		to = idx(r, col)
		capture = move & enemy == enemy
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			fields = append(fields, to)
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
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			fields = append(fields, to)
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
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			fields = append(fields, to)
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
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			fields = append(fields, to)
			if capture {
				break
			}
		} else {
			break
		}
	}
	return fields
}

func rook(board []uint8, piece uint8, enemy uint8, row uint8, col uint8) []Move {
	var moves []Move
	var capture = false
	var move uint8
	var from uint8 = idx(row, col)
	var to uint8
	// -rank
	for r := row - 1; r >= 0; r -= 1 {
		move = field(board, r, col)
		to = idx(r, col)
		capture = move & enemy == enemy
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
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
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
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
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
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
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
			if capture {
				break
			}
		} else {
			break
		}
	}
	return moves
}

func queenAttacks(board []uint8, enemy uint8, row uint8, col uint8) []uint8 {
	var fields []uint8
	var capture bool
	var move uint8
	var to uint8
	// Diagonale 1 → -row,+col
	// TODO limit j to board limits
	for j := 1; j <= 7; j += 1 {
		var r uint8 = uint8(j * -1)
		var c uint8 = uint8(j)
		move = field(board, row + r, col + c)
		to = idx(row + r, col + c)
		capture = move & enemy == enemy
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			fields = append(fields, to)
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
		var r uint8 = uint8(j)
		var c uint8 = uint8(j)
		move = field(board, row + r, col + c)
		to = idx(row + r, col + c)
		capture = move & enemy == enemy
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			fields = append(fields, to)
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
		var r uint8 = uint8(j)
		var c uint8 = uint8(j * -1)
		move = field(board, row + r, col + c)
		to = idx(row + r, col + c)
		capture = move & enemy == enemy
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			fields = append(fields, to)
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
		var r uint8 = uint8(j * -1)
		var c uint8 = uint8(j * -1)
		move = field(board, row + r, col + c)
		to = idx(row + r, col + c)
		capture = move & enemy == enemy
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			fields = append(fields, to)
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
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			fields = append(fields, to)
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
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			fields = append(fields, to)
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
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			fields = append(fields, to)
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
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			fields = append(fields, to)
			if capture {
				break
			}
		} else {
			break
		}
	}
	return fields
}

func queen(board []uint8, piece uint8, enemy uint8, row uint8, col uint8) []Move {
	var moves []Move
	var capture = false
	var move uint8
	var from uint8 = idx(row, col)
	var to uint8
	// Diagonale 1 → -row,+col
	// TODO limit j to board limits
	for j := 1; j <= 7; j += 1 {
		var r uint8 = uint8(j * -1)
		var c uint8 = uint8(j)
		move = field(board, row + r, col + c)
		to = idx(row + r, col + c)
		capture = move & enemy == enemy
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
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
		var r uint8 = uint8(j)
		var c uint8 = uint8(j)
		move = field(board, row + r, col + c)
		to = idx(row + r, col + c)
		capture = move & enemy == enemy
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
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
		var r uint8 = uint8(j)
		var c uint8 = uint8(j * -1)
		move = field(board, row + r, col + c)
		to = idx(row + r, col + c)
		capture = move & enemy == enemy
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
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
		var r uint8 = uint8(j * -1)
		var c uint8 = uint8(j * -1)
		move = field(board, row + r, col + c)
		to = idx(row + r, col + c)
		capture = move & enemy == enemy
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
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
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
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
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
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
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
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
		if to != INVALID_MOVE && (move == EMPTY || capture) {
			moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
			if capture {
				break
			}
		} else {
			break
		}
	}
	return moves
}

func kingAttacks(board []uint8, enemy uint8, row uint8, col uint8) []uint8 {
	var fields []uint8
	var capture = false
	var move uint8
	var to uint8

	// top, -1,0
	move = field(board, row - 1, col)
	to = idx(row - 1, col)
	capture = move & enemy == enemy
	if to != INVALID_MOVE && (move == EMPTY || capture) {
		fields = append(fields, to)
	}

	// top/right, -1,+1
	move = field(board, row - 1, col + 1)
	to = idx(row - 1, col + 1)
	capture = move & enemy == enemy
	if to != INVALID_MOVE && (move == EMPTY || capture) {
		fields = append(fields, to)
	}

	// right, 0,+1
	move = field(board, row, col + 1)
	to = idx(row, col + 1)
	capture = move & enemy == enemy
	if to != INVALID_MOVE && (move == EMPTY || capture) {
		fields = append(fields, to)
	}

	// bottom/right, +1,+1
	move = field(board, row + 1, col + 1)
	to = idx(row + 1, col + 1)
	capture = move & enemy == enemy
	if to != INVALID_MOVE && (move == EMPTY || capture) {
		fields = append(fields, to)
	}

	// bottom, +1,0
	move = field(board, row + 1, col)
	to = idx(row + 1, col)
	capture = move & enemy == enemy
	if to != INVALID_MOVE && (move == EMPTY || capture) {
		fields = append(fields, to)
	}

	// bottom/left, +1,-1
	move = field(board, row + 1, col - 1)
	to = idx(row + 1, col - 1)
	capture = move & enemy == enemy
	if to != INVALID_MOVE && (move == EMPTY || capture) {
		fields = append(fields, to)
	}

	// left, 0,-1
	move = field(board, row, col - 1)
	to = idx(row, col - 1)
	capture = move & enemy == enemy
	if to != INVALID_MOVE && (move == EMPTY || capture) {
		fields = append(fields, to)
	}

	// top/left, -1,-1
	move = field(board, row - 1, col - 1)
	to = idx(row - 1, col - 1)
	capture = move & enemy == enemy
	if to != INVALID_MOVE && (move == EMPTY || capture) {
		fields = append(fields, to)
	}

	return fields
}

func king(board []uint8, attacked[]uint8, piece uint8, enemy uint8, row uint8, col uint8) []Move {
	var moves []Move
	var capture = false
	var move uint8
	var from uint8 = idx(row, col)
	var to uint8
	var isAttackedField bool

	// top, -1,0
	move = field(board, row - 1, col)
	to = idx(row - 1, col)
	capture = move & enemy == enemy
	isAttackedField = isAttacked(attacked, row - 1, col)
	if to != INVALID_MOVE && !isAttackedField && (move == EMPTY || capture) {
		moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
	}

	// top/right, -1,+1
	move = field(board, row - 1, col + 1)
	to = idx(row - 1, col + 1)
	capture = move & enemy == enemy
	isAttackedField = isAttacked(attacked, row - 1, col + 1)
	if to != INVALID_MOVE && !isAttackedField && (move == EMPTY || capture) {
		moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
	}

	// right, 0,+1
	move = field(board, row, col + 1)
	to = idx(row, col + 1)
	capture = move & enemy == enemy
	isAttackedField = isAttacked(attacked, row, col + 1)
	if to != INVALID_MOVE && !isAttackedField && (move == EMPTY || capture) {
		moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
	}

	// bottom/right, +1,+1
	move = field(board, row + 1, col + 1)
	to = idx(row + 1, col + 1)
	capture = move & enemy == enemy
	isAttackedField = isAttacked(attacked, row + 1, col + 1)
	if to != INVALID_MOVE && !isAttackedField && (move == EMPTY || capture) {
		moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
	}

	// bottom, +1,0
	move = field(board, row + 1, col)
	to = idx(row + 1, col)
	capture = move & enemy == enemy
	isAttackedField = isAttacked(attacked, row + 1, col)
	if to != INVALID_MOVE && !isAttackedField && (move == EMPTY || capture) {
		moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
	}

	// bottom/left, +1,-1
	move = field(board, row + 1, col - 1)
	to = idx(row + 1, col - 1)
	capture = move & enemy == enemy
	isAttackedField = isAttacked(attacked, row + 1, col - 1)
	if to != INVALID_MOVE && !isAttackedField && (move == EMPTY || capture) {
		moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
	}

	// left, 0,-1
	move = field(board, row, col - 1)
	to = idx(row, col - 1)
	capture = move & enemy == enemy
	isAttackedField = isAttacked(attacked, row, col - 1)
	if to != INVALID_MOVE && !isAttackedField && (move == EMPTY || capture) {
		moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
	}

	// top/left, -1,-1
	move = field(board, row - 1, col - 1)
	to = idx(row - 1, col - 1)
	capture = move & enemy == enemy
	isAttackedField = isAttacked(attacked, row - 1, col - 1)
	if to != INVALID_MOVE && !isAttackedField && (move == EMPTY || capture) {
		moves = append(moves, Move{Piece: piece, Prev: board[to], From: from, To: to})
	}

	return moves
}

func Attacks(board []uint8, color uint8) []uint8 {
	var fields []uint8
	var col uint8
	var row uint8
	var piece uint8
	var enemy uint8
	var moves []uint8
	var i uint8 = 0

	if color == COLOR_WHITE {
		enemy = COLOR_BLACK
	} else {
		enemy = COLOR_WHITE
	}

	for ; i < 64; i += 1 {
		// get piece
		piece = board[i]

		// check if empty
		if piece == EMPTY {
			continue
		}

		// col & row
		row, col = pos(uint8(i))

		////////////
		/// PAWN ///
		////////////
		if piece == PAWN | color {
			moves = pawnAttacks(board, color, enemy, row, col)
			fields = append(fields, moves...)
			continue
		}

		//////////////
		/// KNIGHT ///
		//////////////
		if piece == KNIGHT | color {
			moves = knightAttacks(board, enemy, row, col)
			fields = append(fields, moves...)
			continue
		}

		//////////////
		/// BISHOP ///
		//////////////
		if piece == BISHOP | color {
			moves = bishopAttacks(board, enemy, row, col)
			fields = append(fields, moves...)
			continue
		}

		////////////
		/// ROOK ///
		////////////
		if piece == BISHOP | color {
			moves = rookAttacks(board, enemy, row, col)
			fields = append(fields, moves...)
			continue
		}

		/////////////
		/// QUEEN ///
		/////////////
		if piece == QUEEN | color {
			moves = queenAttacks(board, enemy, row, col)
			fields = append(fields, moves...)
			continue
		}

		////////////
		/// KING ///
		////////////
		if piece == KING | color {
			moves = kingAttacks(board, enemy, row, col)
			fields = append(fields, moves...)
			continue
		}
	}
	return fields
}

func isAttacked(attacked []uint8, row uint8, col uint8) bool {
	var i = idx(row, col)
	return Contains(attacked, i)
}

func filterKingAttacks(board []uint8, moves []Move, enemy uint8, king uint8) []Move {
	var list []Move
	var attacked []uint8
	var row, col uint8 = pos(king)

	for _, move :=  range moves {
		board = MakeMove(board, move)
		attacked = Attacks(board, enemy)
		if !isAttacked(attacked, row, col) {
			list = append(list, move)
		}
		board = UndoMove(board, move)
	}
	return list
}

func Generate(board []uint8, color uint8) []Move {
	var moves []Move
	var col uint8
	var row uint8
	var piece uint8
	var enemy uint8
	var i uint8 = 0

	if color == COLOR_WHITE {
		enemy = COLOR_BLACK
	} else {
		enemy = COLOR_WHITE
	}

	var attacked = Attacks(board, enemy)
	var kingPos uint8

	for i = 0; i < 64; i += 1 {
		// get piece
		piece = board[i]

		// if king return with pos
		if piece == KING | color {
			kingPos = uint8(i)
			break
		}
	}

	// generate moves
	for i = 0; i < 64; i += 1 {
		// get piece
		piece = board[i]

		// check if empty
		if piece == EMPTY {
			continue
		}

		// col & row
		row, col = pos(uint8(i))

		////////////
		/// PAWN ///
		////////////
		if piece == PAWN | color {
			moves = append(moves, pawn(board, piece, color, enemy, row, col)...)
			continue
		}

		//////////////
		/// KNIGHT ///
		//////////////
		if piece == KNIGHT | color {
			moves = append(moves, knight(board, piece, enemy, row, col)...)
			continue
		}

		//////////////
		/// BISHOP ///
		//////////////
		if piece == BISHOP | color {
			moves = append(moves, bishop(board, piece, enemy, row, col)...)
			continue
		}

		////////////
		/// ROOK ///
		////////////
		if piece == ROOK | color {
			moves = append(moves, rook(board, piece, enemy, row, col)...)
			continue
		}

		/////////////
		/// QUEEN ///
		/////////////
		if piece == QUEEN | color {
			moves = append(moves, queen(board, piece, enemy, row, col)...)
			continue
		}

		////////////
		/// KING ///
		////////////
		if piece == KING | color {
			moves = append(moves, king(board, attacked, piece, enemy, row, col)...)
			continue
		}

	}

	// remove moves for pieces that are pinned (king would be in check)
	moves = filterKingAttacks(board, moves, enemy, kingPos)

	return moves
}
