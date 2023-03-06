package cell

type Cell struct {
  IsBlackHole, IsOpened bool
  BlackHolesAround int
}

func (c *Cell) MarkAsBlackHole() {
  c.IsBlackHole = true
}

func (c *Cell) Open() {
  c.IsOpened = true
}

func (c *Cell) AssignBlackHolesAround(value int) {
  c.BlackHolesAround = value
}
