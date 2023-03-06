package grid

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestInitializeGrid(t *testing.T) {
  // initial state
  assert.Equal(t, len(cells), 0, "Should be empty")

  // First init
  newGridSize := 10
  newBlackHolesCount := 5
  InitializeGrid(newGridSize, newBlackHolesCount)
  assert.Equal(t, newGridSize, len(cells), "Grid should be with desired size of y axis")
  assert.Equal(t, newGridSize, len(cells[0]), "Grid should be with desired size of x axis")
  assert.Equal(t, newBlackHolesCount, len(getCellCoords("blackholes")), "Grid should contain desired number of black holes")

  // Reinitialization
  newGridSize = 2
  newBlackHolesCount = 3
  InitializeGrid(newGridSize, newBlackHolesCount)
  assert.Equal(t, newGridSize, len(cells), "Grid should be with new desired size of y axis")
  assert.Equal(t, newGridSize, len(cells[0]), "Grid should be with new desired of of x axis")
  assert.Equal(t, newBlackHolesCount, len(getCellCoords("blackholes")), "Grid should contain new number of black holes")

  // Regular cells should be populated with correct numbers
  regularCell := getCellCoords("regular")[0]
  assert.Equal(t, newBlackHolesCount, cells[regularCell.x][regularCell.y].BlackHolesAround, "Cell should be populated with correct number of black holes around") // true for given 2x2 grid with 3 black holes
}

func TestGetCellNeighbourCoorinates(t *testing.T) {
  InitializeGrid(5, 0)

  //Top Left
  var expectedCoords = [] gridCoordinates {
    gridCoordinates{x:1, y:0},
    gridCoordinates{x:0, y:1},
    gridCoordinates{x:1, y:1},
  }
  coords := getCellNeighbourCoorinates(0, 0)
  assert.Equal(t, expectedCoords, coords, "Should return correct neighbour coords for top left corner")

  //Center
  expectedCoords = [] gridCoordinates {
    gridCoordinates{x:1, y:1},
    gridCoordinates{x:2, y:1},
    gridCoordinates{x:3, y:1},
    gridCoordinates{x:1, y:2},
    gridCoordinates{x:3, y:2},
    gridCoordinates{x:1, y:3},
    gridCoordinates{x:2, y:3},
    gridCoordinates{x:3, y:3},
  }
  coords = getCellNeighbourCoorinates(2, 2)
  assert.Equal(t, expectedCoords, coords, "Should return correct neighbour coords for centered cell")

  //Bottom Right
  expectedCoords = [] gridCoordinates {
    gridCoordinates{x:3, y:3},
    gridCoordinates{x:4, y:3},
    gridCoordinates{x:3, y:4},
  }
  coords = getCellNeighbourCoorinates(4, 4)
  assert.Equal(t, expectedCoords, coords, "Should return correct neighbour coords for bottom right corner")
}

// private

func getCellCoords(cellType string) []gridCoordinates {
  scopeBlackHoles := false
  if (cellType == "blackholes") {
    scopeBlackHoles = true
  }
  var coords [] gridCoordinates
  for y := 0; y < gridSize; y++ {
    for x := 0; x < gridSize; x++ {
      if (cells[x][y].IsBlackHole == scopeBlackHoles) {
        coords = append(coords, gridCoordinates{x: x, y: y})
      }
    }
  }
  return coords
}
