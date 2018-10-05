package io

import (
	"github.com/eliias/peterpawner/pkg/chess"
	"github.com/eliias/peterpawner/pkg/debug"
	"github.com/eliias/peterpawner/pkg/generator"
	"strconv"
	"strings"
)

func PosCode(pos string) uint8 {
	var parts = strings.Split(pos, "")
	var file = parts[0]
	var rank = parts[1]
	var r = strings.Index(chess.Rank, rank)
	var f = strings.Index(chess.File, file)
	var i = generator.Idx(uint8(7-r), uint8(f))
	return i
}

func ColorCode(color string) uint8 {
	if color == "w" {
		return chess.ColorWhite
	} else {
		return chess.ColorBlack
	}
}

func Load(fen string) chess.Game {
	// board
	var board []uint8

	// parts
	var parts []string = strings.Split(fen, " ")

	// pieces
	var ranks []string = strings.Split(parts[0], "/")
	for _, rank := range ranks {
		for _, piece := range rank {
			// numbers are numbers of pawns in rank
			if l, err := strconv.Atoi(string(piece)); err == nil {
				for i := 0; i < l; i += 1 {
					board = append(board, chess.Empty)
				}
			} else {
				board = append(board, chess.PieceCode(string(piece)))
			}
		}
	}

	// active color
	var color = ColorCode(parts[1])

	// castling
	var castling = parts[2]
	var whiteCastleKingSide = strings.Contains(castling, "K")
	var whiteCastleQueenSide = strings.Contains(castling, "Q")
	var blackCastleKingSide = strings.Contains(castling, "k")
	var blackCastleQueenSide = strings.Contains(castling, "q")

	// en passant
	var enPassant uint8
	if parts[3] != "-" {
		enPassant = PosCode(parts[3])
	} else {
		enPassant = chess.InvalidMove
	}

	// half move
	var halfmove int
	if hm, err := strconv.Atoi(string(parts[4])); err == nil {
		halfmove = hm
	}

	// full move
	var fullmove int
	if fm, err := strconv.Atoi(string(parts[5])); err == nil {
		fullmove = fm
	}

	// game
	return chess.Game{
		Board:                board,
		Color:                color,
		WhiteCastleKingSide:  whiteCastleKingSide,
		WhiteCastleQueenSide: whiteCastleQueenSide,
		BlackCastleKingSide:  blackCastleKingSide,
		BlackCastleQueenSide: blackCastleQueenSide,
		EnPassant:            enPassant,
		Halfmove:             halfmove,
		Fullmove:             fullmove}
}

func Save(game chess.Game) string {
	var str = ""
	var cnt = 0

	// board
	for i := 0; i < 64; i += 1 {
		// piece
		var piece = game.Board[uint8(i)]

		// empty
		if piece == chess.Empty {
			// found an empty field
			cnt += 1
		}

		// regular piece
		if piece != chess.Empty {
			if cnt > 0 {
				str += strconv.FormatInt(int64(cnt), 10)
			}
			str += debug.PieceName(piece)
			cnt = 0
		}

		// new rank
		if (i+1)%8 == 0 {
			if cnt > 0 {
				str += strconv.FormatInt(int64(cnt), 10)
			}
			if i < 63 {
				str += "/"
			}
			cnt = 0
		}
	}

	// active color
	str += " " + debug.ColorName(game.Color)

	// castling
	var castling = ""
	if game.WhiteCastleKingSide {
		castling += "K"
	}
	if game.WhiteCastleQueenSide {
		castling += "Q"
	}
	if game.BlackCastleKingSide {
		castling += "k"
	}
	if game.BlackCastleQueenSide {
		castling += "q"
	}
	if len(castling) == 0 {
		castling = "-"
	}
	str += " " + castling

	// en passant
	var enPassant = ""
	if game.EnPassant != chess.InvalidMove {
		enPassant = debug.DebugPos(game.EnPassant)
	} else {
		enPassant = "-"
	}
	str += " " + enPassant

	// halfmove
	str += " " + strconv.FormatInt(int64(game.Halfmove), 10)

	// fullmove
	str += " " + strconv.FormatInt(int64(game.Fullmove), 10)

	return str
}
