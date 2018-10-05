package generator

import "github.com/eliias/peterpawner/pkg/chess"

func pawn(board []uint8, piece uint8, color uint8, enemy uint8, row uint8, col uint8, enPassantFields []uint8) []*chess.Move {
	var moves []*chess.Move
	var capture = false
	var move uint8
	var from = Idx(row, col)
	var to uint8
	var trow uint8
	var enPassantRow uint8
	var enPassant uint8
	var isEnPassantTarget bool
	// move one up, target must be empty
	if color == chess.ColorBlack {
		trow = row + 1
	} else {
		trow = row - 1
	}
	move = field(board, trow, col)
	if move == chess.Empty {
		to = Idx(trow, col)
		moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	// move two up, if in starting position & next two fields are clear
	// also save En passant square
	if move == chess.Empty && (color == chess.ColorWhite && row == 6 || color == chess.ColorBlack && row == 1) {
		if color == chess.ColorBlack {
			trow = row + 2
			enPassantRow = row + 1
		} else {
			trow = row - 2
			enPassantRow = row - 1
		}
		move = field(board, trow, col)
		if move == chess.Empty {
			to = Idx(trow, col)
			enPassant = Idx(enPassantRow, col)
			moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to, EnPassant: enPassant})
		}
	}
	// attack left/up if enemy piece is there or enpassant square
	if color == chess.ColorBlack {
		trow = row + 1
	} else {
		trow = row - 1
	}
	move = field(board, trow, col-1)
	capture = move&enemy == enemy
	to = Idx(trow, col-1)
	isEnPassantTarget = chess.Contains(enPassantFields, to)
	if move != chess.InvalidMove && (capture || isEnPassantTarget) {
		moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	// attack right/up if enemy piece is there or enpassant square
	move = field(board, trow, col+1)
	capture = move&enemy == enemy
	to = Idx(trow, col+1)
	isEnPassantTarget = chess.Contains(enPassantFields, to)
	if move != chess.InvalidMove && (capture || isEnPassantTarget) {
		moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	return moves
}

func knight(board []uint8, piece uint8, enemy uint8, row uint8, col uint8) []*chess.Move {
	var moves []*chess.Move
	var capture = false
	var move uint8
	var from = Idx(row, col)
	var to uint8
	// Quadrant 1
	// -2, +1
	move = field(board, row-2, col+1)
	to = Idx(row-2, col+1)
	capture = move&enemy == enemy
	if to != chess.InvalidMove && (move == chess.Empty || capture) {
		moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	// -1, +2
	move = field(board, row-1, col+2)
	to = Idx(row-1, col+2)
	capture = move&enemy == enemy
	if to != chess.InvalidMove && (move == chess.Empty || capture) {
		moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	// Quadrant 2
	// +2, +1
	move = field(board, row+2, col+1)
	to = Idx(row+2, col+1)
	capture = move&enemy == enemy
	if to != chess.InvalidMove && (move == chess.Empty || capture) {
		moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	// +1, +2
	move = field(board, row+1, col+2)
	to = Idx(row+1, col+2)
	capture = move&enemy == enemy
	if to != chess.InvalidMove && (move == chess.Empty || capture) {
		moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	// Quadrant 3
	// +2, -1
	move = field(board, row+2, col-1)
	to = Idx(row+2, col-1)
	capture = move&enemy == enemy
	if to != chess.InvalidMove && (move == chess.Empty || capture) {
		moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	// +1, -2
	move = field(board, row+1, col-2)
	to = Idx(row+1, col-2)
	capture = move&enemy == enemy
	if to != chess.InvalidMove && (move == chess.Empty || capture) {
		moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	// Quadrant 4
	// -2, +1
	move = field(board, row-2, col-1)
	to = Idx(row-2, col-1)
	capture = move&enemy == enemy
	if to != chess.InvalidMove && (move == chess.Empty || capture) {
		moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	// -1, +2
	move = field(board, row-1, col-2)
	to = Idx(row-1, col-2)
	capture = move&enemy == enemy
	if to != chess.InvalidMove && (move == chess.Empty || capture) {
		moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
	}
	return moves
}

func bishop(board []uint8, piece uint8, enemy uint8, row uint8, col uint8) []*chess.Move {
	var moves []*chess.Move
	var capture = false
	var move uint8
	var from = Idx(row, col)
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
			moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
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
			moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
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
			moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
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
			moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
			if capture {
				break
			}
		} else {
			break
		}
	}
	return moves
}

func rook(board []uint8, piece uint8, enemy uint8, row uint8, col uint8) []*chess.Move {
	var moves []*chess.Move
	var capture = false
	var move uint8
	var from = Idx(row, col)
	var to uint8
	// -rank
	for r := row - 1; r >= 0; r -= 1 {
		move = field(board, r, col)
		to = Idx(r, col)
		capture = move&enemy == enemy
		if to != chess.InvalidMove && (move == chess.Empty || capture) {
			moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
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
			moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
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
			moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
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
			moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
			if capture {
				break
			}
		} else {
			break
		}
	}
	return moves
}

func queen(board []uint8, piece uint8, enemy uint8, row uint8, col uint8) []*chess.Move {
	var moves []*chess.Move
	var capture = false
	var move uint8
	var from = Idx(row, col)
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
			moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
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
			moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
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
			moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
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
			moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
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
			moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
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
			moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
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
			moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
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
			moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
			if capture {
				break
			}
		} else {
			break
		}
	}
	return moves
}

func king(board []uint8, attacked []uint8, piece uint8, enemy uint8, row uint8, col uint8) []*chess.Move {
	var moves []*chess.Move
	var capture = false
	var move uint8
	var from = Idx(row, col)
	var to uint8
	var isAttackedField bool

	// top, -1,0
	move = field(board, row-1, col)
	to = Idx(row-1, col)
	capture = move&enemy == enemy
	isAttackedField = isAttacked(attacked, row-1, col)
	if to != chess.InvalidMove && !isAttackedField && (move == chess.Empty || capture) {
		moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
	}

	// top/right, -1,+1
	move = field(board, row-1, col+1)
	to = Idx(row-1, col+1)
	capture = move&enemy == enemy
	isAttackedField = isAttacked(attacked, row-1, col+1)
	if to != chess.InvalidMove && !isAttackedField && (move == chess.Empty || capture) {
		moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
	}

	// right, 0,+1
	move = field(board, row, col+1)
	to = Idx(row, col+1)
	capture = move&enemy == enemy
	isAttackedField = isAttacked(attacked, row, col+1)
	if to != chess.InvalidMove && !isAttackedField && (move == chess.Empty || capture) {
		moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
	}

	// bottom/right, +1,+1
	move = field(board, row+1, col+1)
	to = Idx(row+1, col+1)
	capture = move&enemy == enemy
	isAttackedField = isAttacked(attacked, row+1, col+1)
	if to != chess.InvalidMove && !isAttackedField && (move == chess.Empty || capture) {
		moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
	}

	// bottom, +1,0
	move = field(board, row+1, col)
	to = Idx(row+1, col)
	capture = move&enemy == enemy
	isAttackedField = isAttacked(attacked, row+1, col)
	if to != chess.InvalidMove && !isAttackedField && (move == chess.Empty || capture) {
		moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
	}

	// bottom/left, +1,-1
	move = field(board, row+1, col-1)
	to = Idx(row+1, col-1)
	capture = move&enemy == enemy
	isAttackedField = isAttacked(attacked, row+1, col-1)
	if to != chess.InvalidMove && !isAttackedField && (move == chess.Empty || capture) {
		moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
	}

	// left, 0,-1
	move = field(board, row, col-1)
	to = Idx(row, col-1)
	capture = move&enemy == enemy
	isAttackedField = isAttacked(attacked, row, col-1)
	if to != chess.InvalidMove && !isAttackedField && (move == chess.Empty || capture) {
		moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
	}

	// top/left, -1,-1
	move = field(board, row-1, col-1)
	to = Idx(row-1, col-1)
	capture = move&enemy == enemy
	isAttackedField = isAttacked(attacked, row-1, col-1)
	if to != chess.InvalidMove && !isAttackedField && (move == chess.Empty || capture) {
		moves = append(moves, &chess.Move{Piece: piece, Prev: board[to], From: from, To: to})
	}

	return moves
}
