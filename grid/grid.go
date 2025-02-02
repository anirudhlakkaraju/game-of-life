package grid

// Grid defines a grid and it's dimensions
type Grid struct {
	Width  int
	Height int

	Cells [][]bool
}

// NewGrid creates a grid
func NewGrid(width, height int) *Grid {
	cells := make([][]bool, height)
	for i := range height {
		cells[i] = make([]bool, width)
	}
	return &Grid{width, height, cells}
}
