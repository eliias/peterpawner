package debug

func Board(board [64]int16) string {
  var str = ""
  var col int
  var row int = 0
  for i := 0; i < 100; i += 1 {
    if i % 10 == 0 && i > 0 {
      row += 1
      str += "\n"
    }
    col = i - row * 10

    if col == 0 && row == 0 || col == 9 && row == 0 || col == 0 && row == 9 || col == 9 && row == 9 {
      str += "+"
    } else if row == 0 && col > 0 && col < 9 {
      str += "-"
    } else if row == 9 && col > 0 && col < 9 {
      str += "-"
    } else if col == 0 || col == 9 {
      str += "|"
    } else {
      str += PieceName(board[col + 8 * (row - 1) - 1])
    }
  }

  return str
}
