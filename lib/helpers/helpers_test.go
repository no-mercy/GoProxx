package helpers

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
  assert.Equal(t, 2, Max(1, 2), "Should return max value")
  assert.Equal(t, 4, Max(4, -3), "Should return max value")
}

func TestMin(t *testing.T) {
  assert.Equal(t, 1, Min(1, 2), "Should return min value")
  assert.Equal(t, -3, Min(4, -3), "Should return min value")
}

func TestParseUserInput(t *testing.T) {
  x, y := ParseUserInput("100,500")
  assert.Equal(t, 100, x, "Should be converted to int and assigned to first value")
  assert.Equal(t, 500, y, "Should be converted to int and assigned to second value")
}
