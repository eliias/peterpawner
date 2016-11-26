package main

import (
  "github.com/eliias/peterpawner/debug"
  "github.com/eliias/peterpawner/moves"
  "fmt"
)

func main() {
  fmt.Println("\n-----------------\npeterpawner 1.0\nHannes Moser 2016\n-----------------\n")

  var board = moves.Start
  var color = moves.COLOR_WHITE
  fmt.Println(debug.Board(board))

  var depth = 3
  fmt.Println("Perft (", depth, "):", moves.Perft(board, depth, color))

  //fmt.Println(debug.Moves(board, depth, color))
}
