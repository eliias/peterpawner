package generator

import "github.com/eliias/peterpawner/pkg/chess"

func pos(i uint8) (row uint8, col uint8) {
	row = i / 8
	col = i - row*8
	return
}

func Idx(row uint8, col uint8) uint8 {
	if row < 0 || row > 7 {
		return chess.InvalidMove
	}
	if col < 0 || col > 7 {
		return chess.InvalidMove
	}
	return row*8 + col
}

func field(board []uint8, row uint8, col uint8) uint8 {
	if row < 0 || row > 7 {
		return chess.InvalidMove
	}
	if col < 0 || col > 7 {
		return chess.InvalidMove
	}
	var i = Idx(row, col)
	return board[i]
}

func Attacks(board []uint8, color uint8, enPassant []uint8) []uint8 {
	var fields []uint8
	var col uint8
	var row uint8
	var piece uint8
	var enemy uint8
	var moves []uint8
	var i uint8 = 0

	if color == chess.ColorWhite {
		enemy = chess.ColorBlack
	} else {
		enemy = chess.ColorWhite
	}

	for ; i < 64; i += 1 {
		// get piece
		piece = board[i]

		// check if empty
		if piece == chess.Empty {
			continue
		}

		// col & row
		row, col = pos(uint8(i))

		////////////
		/// Pawn ///
		////////////
		if piece == chess.Pawn|color {
			moves = pawnAttacks(board, color, enemy, row, col, enPassant)
			fields = append(fields, moves...)
			continue
		}

		//////////////
		/// Knight ///
		//////////////
		if piece == chess.Knight|color {
			moves = knightAttacks(board, enemy, row, col)
			fields = append(fields, moves...)
			continue
		}

		//////////////
		/// Bishop ///
		//////////////
		if piece == chess.Bishop|color {
			moves = bishopAttacks(board, enemy, row, col)
			fields = append(fields, moves...)
			continue
		}

		////////////
		/// Rook ///
		////////////
		if piece == chess.Bishop|color {
			moves = rookAttacks(board, enemy, row, col)
			fields = append(fields, moves...)
			continue
		}

		/////////////
		/// Queen ///
		/////////////
		if piece == chess.Queen|color {
			moves = queenAttacks(board, enemy, row, col)
			fields = append(fields, moves...)
			continue
		}

		////////////
		/// King ///
		////////////
		if piece == chess.King|color {
			moves = kingAttacks(board, enemy, row, col)
			fields = append(fields, moves...)
			continue
		}
	}
	return fields
}

func isAttacked(attacked []uint8, row uint8, col uint8) bool {
	var i = Idx(row, col)
	return chess.Contains(attacked, i)
}

func filterByPiece(board []uint8, moves []*chess.Move, enemy uint8, piece uint8, enPassant []uint8) []*chess.Move {
	var list []*chess.Move
	var attacked []uint8
	var row, col uint8 = pos(piece)

	for _, move := range moves {
		board = chess.MakeMove(board, move)
		attacked = Attacks(board, enemy, enPassant)
		if !isAttacked(attacked, row, col) {
			list = append(list, move)
		}
		board = chess.UndoMove(board, move)
	}
	return list
}

func Generate(board []uint8, color uint8, enPassant []uint8) []*chess.Move {
	var moves []*chess.Move
	var col uint8
	var row uint8
	var piece uint8
	var enemy uint8
	var i uint8 = 0

	if color == chess.ColorWhite {
		enemy = chess.ColorBlack
	} else {
		enemy = chess.ColorWhite
	}

	var attacked = Attacks(board, enemy, enPassant)
	var kingPos uint8

	for i = 0; i < 64; i += 1 {
		// get piece
		piece = board[i]

		// if king return with pos
		if piece == chess.King|color {
			kingPos = uint8(i)
			break
		}
	}

	// generate moves
	for i = 0; i < 64; i += 1 {
		// get piece
		piece = board[i]

		// check if empty
		if piece == chess.Empty {
			continue
		}

		// col & row
		row, col = pos(uint8(i))

		////////////
		/// Pawn ///
		////////////
		if piece == chess.Pawn|color {
			moves = append(moves, pawn(board, piece, color, enemy, row, col, enPassant)...)
			continue
		}

		//////////////
		/// Knight ///
		//////////////
		if piece == chess.Knight|color {
			moves = append(moves, knight(board, piece, enemy, row, col)...)
			continue
		}

		//////////////
		/// Bishop ///
		//////////////
		if piece == chess.Bishop|color {
			moves = append(moves, bishop(board, piece, enemy, row, col)...)
			continue
		}

		////////////
		/// Rook ///
		////////////
		if piece == chess.Rook|color {
			moves = append(moves, rook(board, piece, enemy, row, col)...)
			continue
		}

		/////////////
		/// Queen ///
		/////////////
		if piece == chess.Queen|color {
			moves = append(moves, queen(board, piece, enemy, row, col)...)
			continue
		}

		////////////
		/// King ///
		////////////
		if piece == chess.King|color {
			moves = append(moves, king(board, attacked, piece, enemy, row, col)...)
			continue
		}
	}

	// remove moves for pieces that are pinned (king would be in check)
	moves = filterByPiece(board, moves, enemy, kingPos, enPassant)

	return moves
}
