package main

import (
  "fmt"
  "proxx/lib/grid"
  h "proxx/lib/helpers"
)

func main() {
  grid.InitializeGrid(h.GetEnv("GRID_SIZE", 10), h.GetEnv("BLACKHOLES_COUNT", 10))
  input := ""
  turnResult := 0
  for turnResult < 2 {
    grid.PrintMatrix()
    fmt.Println("Enter cell coordinates like 0,0")
    fmt.Scanln(&input)
    x, y := h.ParseUserInput(input)
    turnResult = grid.OpenCell(x,y)
  }
  if (turnResult == 666) {
    fmt.Println("You have been sucked into black hole")
  } else {
    fmt.Println("You survived")
  }
}
