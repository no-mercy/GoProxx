package cell

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "math/rand"
  "time"
)

var c Cell

func TestMarkAsBlackHole(t *testing.T) {
  assert.Equal(t, false, c.IsBlackHole, "Should be false by default")
  c.MarkAsBlackHole()
  assert.Equal(t, true, c.IsBlackHole, "Should be true after call")
}

func TestOpen(t *testing.T) {
  assert.Equal(t, false, c.IsOpened, "Should be false by default")
  c.Open()
  assert.Equal(t, true, c.IsOpened, "Should be true after call")
}

func AssignBlackHolesAround(t *testing.T) {
  rand.Seed(time.Now().UnixNano())
  value := rand.Intn(100)
  assert.Equal(t, 0, c.BlackHolesAround, "Should be zero by default")
  c.AssignBlackHolesAround(value)
  assert.Equal(t, value, c.BlackHolesAround, "Should have correct value after call")
}
