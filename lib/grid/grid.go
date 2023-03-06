package grid

import (
    "proxx/lib/cell"
  h "proxx/lib/helpers"
    "math/rand"
    "time"
    "fmt"
    "strings"
)

var cells[][] cell.Cell
var gridSize int
var blackHolesCount int
var openedCells int

type gridCoordinates struct {
  x, y int
}

func InitializeGrid(size, desiredHolesCount int) {
  gridSize = size
  cells = make([][] cell.Cell, size, size)
  for n := range cells {
    cells[n] = make([] cell.Cell, size, size)
  }
  seedBlackHoles(desiredHolesCount)
  populateNearbyBlackHolesCounters()
}

func PrintMatrix() {
  fmt.Println(strings.Repeat("-", gridSize*4))
  for y := 0; y < gridSize; y++ {
    fmt.Print("| ")
    for x := 0; x < gridSize; x++ {
      cell := cells[x][y]
      if (cell.IsOpened) {
        fmt.Print(cell.BlackHolesAround)
      } else {
        fmt.Print(" ")
      }
      fmt.Print(" | ")
    }
    fmt.Println("")
    fmt.Println(strings.Repeat("-", gridSize*4))
  }
}

func OpenCell(x, y int) int {
  cell := cells[x][y]
  if (cell.IsOpened) {
    return 0
  }
  if (cell.IsBlackHole) {
    return 666
  }
  cell.Open()
  cells[x][y] = cell
  openedCells += 1
  if (cell.BlackHolesAround == 0) {
    cellNeighboursCoordsList := getCellNeighbourCoorinates(x,y)
    for i := range cellNeighboursCoordsList {
      coords := cellNeighboursCoordsList[i]
      if(cells[coords.x][coords.y].IsBlackHole) {
        continue
      }
      OpenCell(coords.x, coords.y)
    }
  }
  if (openedCells == (gridSize * gridSize - blackHolesCount)) {
    return 777
  } else {
    return 1
  }
}

// private

func seedBlackHoles(desiredHolesCount int) {
  blackHolesCount = desiredHolesCount
  if (desiredHolesCount < 1 || gridSize*gridSize < desiredHolesCount) {
    fmt.Println("out of bound desired black holes count")
    return
  }
  holesPlaced := 0
  rand.Seed(time.Now().UnixNano())
  for holesPlaced < desiredHolesCount {
    x := rand.Intn(gridSize)
    y := rand.Intn(gridSize)
    cell := cells[x][y]
    if (!cell.IsBlackHole) {
      cell.MarkAsBlackHole()
      cells[x][y] = cell
      holesPlaced += 1
    }
  }
}

func getCellNeighbourCoorinates(cellX, cellY int) [] gridCoordinates {
  var result [] gridCoordinates
  for y := h.Max(0, cellY - 1); y <= h.Min(cellY + 1, gridSize - 1); y++ {
    for x := h.Max(0, cellX - 1); x <= h.Min(cellX + 1, gridSize - 1); x++ {
      if(y == cellY && x == cellX) {
        continue
      }
      result = append(result, gridCoordinates{x: x, y: y})
    }
  }
  return result
}

func populateNearbyBlackHolesCounters() {
  for y := 0; y < gridSize; y++ {
    for x := 0; x < gridSize; x++ {
      cellNeighboursCoordsList := getCellNeighbourCoorinates(x,y)
      blackHolesCounter := 0
      for i := range cellNeighboursCoordsList {
        coords := cellNeighboursCoordsList[i]
        if(cells[coords.x][coords.y].IsBlackHole) {
          blackHolesCounter += 1
        }
      }
      cell := cells[x][y]
      cell.AssignBlackHolesAround(blackHolesCounter)
      cells[x][y] = cell
    }
  }
}
