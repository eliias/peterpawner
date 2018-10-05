package chess

const Empty uint8 = 0

const King uint8 = 1 << 0
const Queen uint8 = 1 << 1
const Rook uint8 = 1 << 2
const Bishop uint8 = 1 << 3
const Knight uint8 = 1 << 4
const Pawn uint8 = 1 << 5

const ColorWhite = 1 << 6
const ColorBlack = 1 << 7

const WhiteKing = King | ColorWhite
const WhiteQueen = Queen | ColorWhite
const WhiteRook = Rook | ColorWhite
const WhiteBishop = Bishop | ColorWhite
const WhiteKnight = Knight | ColorWhite
const WhitePawn = Pawn | ColorWhite

const BlackKing = King | ColorBlack
const BlackQueen = Queen | ColorBlack
const BlackRook = Rook | ColorBlack
const BlackBishop = Bishop | ColorBlack
const BlackKnight = Knight | ColorBlack
const BlackPawn = Pawn | ColorBlack

const InvalidMove uint8 = 255
