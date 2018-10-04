# Attacks

The attack array combines all possible attacks of all pieces. The
integer values represent a combination of all pieces that can move
to a certain field. Use bit operations to identify a valid field
for a given piece.

```go
package chess

const CENTER uint8 = 112

const attacks = []uint8{
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}
```

# Pieces

```go
package chess

const KING uint8 = 1 << 0 // 1
const QUEEN uint8 = 1 << 1 // 2
const ROOK uint8 = 1 << 2 // 4
const BISHOP uint8 = 1 << 3 // 8
const KNIGHT uint8 = 1 << 4 // 16
const PAWN uint8 = 1 << 5 // 32
```

```go
package chess

var field uint8

// -1, -1
field = KING | QUEEN | BISHOP
//  0, -1
field = KING | QUEEN | ROOK
// +1, -1
field = KING | QUEEN | BISHOP
// +1,  0
field = KING | QUEEN | ROOK
// +1, +1
field = KING | QUEEN | BISHOP
//  0, +1
field = KING | QUEEN | ROOK
// -1, +1
field = KING | QUEEN | BISHOP
// -1,  0
field = KING | QUEEN | ROOK
```
