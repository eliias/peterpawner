package generator

import "github.com/eliias/peterpawner/pkg/chess"

func pawnAttacks(board []uint8, color uint8, enemy uint8, row uint8, col uint8, enPassantFields []uint8) []uint8 {
	var fields []uint8
	var move uint8
	var trow uint8
	var to uint8
	var capture bool
	var empty bool
	var isEnPassantTarget bool

	// attack left/up if enemy piece is there
	if color == chess.ColorBlack {
		trow = row + 1
	} else {
		trow = row - 1
	}
	move = field(board, trow, col-1)
	capture = move&enemy == enemy
	empty = move == chess.Empty
	to = Idx(trow, col-1)
	isEnPassantTarget = chess.Contains(enPassantFields, to)
	if move != chess.InvalidMove && (capture || empty || isEnPassantTarget) {
		fields = append(fields, to)
	}
	// attack right/up if enemy piece is there
	move = field(board, trow, col+1)
	capture = move&enemy == enemy
	empty = move == chess.Empty
	to = Idx(trow, col+1)
	isEnPassantTarget = chess.Contains(enPassantFields, to)
	if move != chess.InvalidMove && (capture || empty || isEnPassantTarget) {
		fields = append(fields, to)
	}
	return fields
}

func knightAttacks(board []uint8, enemy uint8, row uint8, col uint8) []uint8 {
	var fields []uint8
	var capture bool
	var empty bool
	var move uint8
	var to uint8
	// Quadrant 1
	// -2, +1
	move = field(board, row-2, col+1)
	to = Idx(row-2, col+1)
	capture = move&enemy == enemy
	empty = move == chess.Empty
	if to != chess.InvalidMove && (empty || capture) {
		fields = append(fields, to)
	}
	// -1, +2
	move = field(board, row-1, col+2)
	to = Idx(row-1, col+2)
	capture = move&enemy == enemy
	empty = move == chess.Empty
	if to != chess.InvalidMove && (empty || capture) {
		fields = append(fields, to)
	}
	// Quadrant 2
	// +2, +1
	move = field(board, row+2, col+1)
	to = Idx(row+2, col+1)
	capture = move&enemy == enemy
	empty = move == chess.Empty
	if to != chess.InvalidMove && (empty || capture) {
		fields = append(fields, to)
	}
	// +1, +2
	move = field(board, row+1, col+2)
	to = Idx(row+1, col+2)
	capture = move&enemy == enemy
	empty = move == chess.Empty
	if to != chess.InvalidMove && (empty || capture) {
		fields = append(fields, to)
	}
	// Quadrant 3
	// +2, -1
	move = field(board, row+2, col-1)
	to = Idx(row+2, col-1)
	capture = move&enemy == enemy
	empty = move == chess.Empty
	if to != chess.InvalidMove && (empty || capture) {
		fields = append(fields, to)
	}
	// +1, -2
	move = field(board, row+1, col-2)
	to = Idx(row+1, col-2)
	capture = move&enemy == enemy
	empty = move == chess.Empty
	if to != chess.InvalidMove && (empty || capture) {
		fields = append(fields, to)
	}
	// Quadrant 4
	// -2, +1
	move = field(board, row-2, col-1)
	to = Idx(row-2, col-1)
	capture = move&enemy == enemy
	empty = move == chess.Empty
	if to != chess.InvalidMove && (empty || capture) {
		fields = append(fields, to)
	}
	// -1, +2
	move = field(board, row-1, col-2)
	to = Idx(row-1, col-2)
	capture = move&enemy == enemy
	empty = move == chess.Empty
	if to != chess.InvalidMove && (empty || capture) {
		fields = append(fields, to)
	}
	return fields
}

func bishopAttacks(board []uint8, enemy uint8, row uint8, col uint8) []uint8 {
	var fields []uint8
	var capture bool
	var move uint8
	var to uint8
	// Diagonale 1 → -row,+col
	// TODO limit j to board limits
	for j := 1; j <= 7; j += 1 {
		var r = uint8(j * -1)
		var c = uint8(j)
		move = field(board, row+r, col+c)
		to = Idx(row+r, col+c)
		capture = move&enemy == enemy
		if to != chess.InvalidMove && (move == chess.Empty || capture) {
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
		var r = uint8(j)
		var c = uint8(j)
		move = field(board, row+r, col+c)
		to = Idx(row+r, col+c)
		capture = move&enemy == enemy
		if to != chess.InvalidMove && (move == chess.Empty || capture) {
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
		var r = uint8(j)
		var c = uint8(j * -1)
		move = field(board, row+r, col+c)
		to = Idx(row+r, col+c)
		capture = move&enemy == enemy
		if to != chess.InvalidMove && (move == chess.Empty || capture) {
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
		var r = uint8(j * -1)
		var c = uint8(j * -1)
		move = field(board, row+r, col+c)
		to = Idx(row+r, col+c)
		capture = move&enemy == enemy
		if to != chess.InvalidMove && (move == chess.Empty || capture) {
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

func rookAttacks(board []uint8, enemy uint8, row uint8, col uint8) []uint8 {
	var fields []uint8
	var capture bool
	var move uint8
	var to uint8
	// -rank
	for r := row - 1; r >= 0; r -= 1 {
		move = field(board, r, col)
		to = Idx(r, col)
		capture = move&enemy == enemy
		if to != chess.InvalidMove && (move == chess.Empty || capture) {
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
		to = Idx(r, col)
		capture = move&enemy == enemy
		if to != chess.InvalidMove && (move == chess.Empty || capture) {
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
		to = Idx(row, c)
		capture = move&enemy == enemy
		if to != chess.InvalidMove && (move == chess.Empty || capture) {
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
		to = Idx(row, c)
		capture = move&enemy == enemy
		if to != chess.InvalidMove && (move == chess.Empty || capture) {
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

func queenAttacks(board []uint8, enemy uint8, row uint8, col uint8) []uint8 {
	var fields []uint8
	var capture bool
	var move uint8
	var to uint8
	// Diagonale 1 → -row,+col
	// TODO limit j to board limits
	for j := 1; j <= 7; j += 1 {
		var r = uint8(j * -1)
		var c = uint8(j)
		move = field(board, row+r, col+c)
		to = Idx(row+r, col+c)
		capture = move&enemy == enemy
		if to != chess.InvalidMove && (move == chess.Empty || capture) {
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
		var r = uint8(j)
		var c = uint8(j)
		move = field(board, row+r, col+c)
		to = Idx(row+r, col+c)
		capture = move&enemy == enemy
		if to != chess.InvalidMove && (move == chess.Empty || capture) {
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
		var r = uint8(j)
		var c = uint8(j * -1)
		move = field(board, row+r, col+c)
		to = Idx(row+r, col+c)
		capture = move&enemy == enemy
		if to != chess.InvalidMove && (move == chess.Empty || capture) {
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
		var r = uint8(j * -1)
		var c = uint8(j * -1)
		move = field(board, row+r, col+c)
		to = Idx(row+r, col+c)
		capture = move&enemy == enemy
		if to != chess.InvalidMove && (move == chess.Empty || capture) {
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
		to = Idx(r, col)
		capture = move&enemy == enemy
		if to != chess.InvalidMove && (move == chess.Empty || capture) {
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
		to = Idx(r, col)
		capture = move&enemy == enemy
		if to != chess.InvalidMove && (move == chess.Empty || capture) {
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
		to = Idx(row, c)
		capture = move&enemy == enemy
		if to != chess.InvalidMove && (move == chess.Empty || capture) {
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
		to = Idx(row, c)
		capture = move&enemy == enemy
		if to != chess.InvalidMove && (move == chess.Empty || capture) {
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

func kingAttacks(board []uint8, enemy uint8, row uint8, col uint8) []uint8 {
	var fields []uint8
	var capture = false
	var move uint8
	var to uint8

	// top, -1,0
	move = field(board, row-1, col)
	to = Idx(row-1, col)
	capture = move&enemy == enemy
	if to != chess.InvalidMove && (move == chess.Empty || capture) {
		fields = append(fields, to)
	}

	// top/right, -1,+1
	move = field(board, row-1, col+1)
	to = Idx(row-1, col+1)
	capture = move&enemy == enemy
	if to != chess.InvalidMove && (move == chess.Empty || capture) {
		fields = append(fields, to)
	}

	// right, 0,+1
	move = field(board, row, col+1)
	to = Idx(row, col+1)
	capture = move&enemy == enemy
	if to != chess.InvalidMove && (move == chess.Empty || capture) {
		fields = append(fields, to)
	}

	// bottom/right, +1,+1
	move = field(board, row+1, col+1)
	to = Idx(row+1, col+1)
	capture = move&enemy == enemy
	if to != chess.InvalidMove && (move == chess.Empty || capture) {
		fields = append(fields, to)
	}

	// bottom, +1,0
	move = field(board, row+1, col)
	to = Idx(row+1, col)
	capture = move&enemy == enemy
	if to != chess.InvalidMove && (move == chess.Empty || capture) {
		fields = append(fields, to)
	}

	// bottom/left, +1,-1
	move = field(board, row+1, col-1)
	to = Idx(row+1, col-1)
	capture = move&enemy == enemy
	if to != chess.InvalidMove && (move == chess.Empty || capture) {
		fields = append(fields, to)
	}

	// left, 0,-1
	move = field(board, row, col-1)
	to = Idx(row, col-1)
	capture = move&enemy == enemy
	if to != chess.InvalidMove && (move == chess.Empty || capture) {
		fields = append(fields, to)
	}

	// top/left, -1,-1
	move = field(board, row-1, col-1)
	to = Idx(row-1, col-1)
	capture = move&enemy == enemy
	if to != chess.InvalidMove && (move == chess.Empty || capture) {
		fields = append(fields, to)
	}

	return fields
}
