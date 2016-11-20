package main

import "fmt"
import "./moves"
import "./debug"

func main() {
  fmt.Println("\n-----------------\npeterpawner 1.0\nHannes Moser 2016\n-----------------\n")

  var board = moves.Start
  fmt.Println(debug.Board(board))

  var depth = 7
  fmt.Println("Perft (", depth, "):", moves.Perft(depth))

  /*
  var list = moves.Generate(board, moves.COLOR_WHITE)
  var n [64]int16
  for _, move := range list {
    n = moves.Add(board, move)
    fmt.Println(debug.Board(n))
  }
  */
}
