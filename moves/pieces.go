package moves

const COLOR_WHITE int16 = 1 << 7
const COLOR_BLACK int16 = 1 << 8

const EMPTY int16 = 1 << 9

const KING int16 = 1 << 1
const QUEEN int16 = 1 << 2
const ROOK int16 = 1 << 3
const BISHOP int16 = 1 << 4
const KNIGHT int16 = 1 << 5
const PAWN int16 = 1 << 6

const W_KING int16 = KING | COLOR_WHITE
const W_QUEEN int16 = QUEEN | COLOR_WHITE
const W_ROOK int16 = ROOK | COLOR_WHITE
const W_BISHOP int16 = BISHOP | COLOR_WHITE
const W_KNIGHT int16 = KNIGHT | COLOR_WHITE
const W_PAWN int16 = PAWN | COLOR_WHITE

const B_KING int16 = KING | COLOR_BLACK
const B_QUEEN int16 = QUEEN | COLOR_BLACK
const B_ROOK int16 = ROOK | COLOR_BLACK
const B_BISHOP int16 = BISHOP | COLOR_BLACK
const B_KNIGHT int16 = KNIGHT | COLOR_BLACK
const B_PAWN int16 = PAWN | COLOR_BLACK
