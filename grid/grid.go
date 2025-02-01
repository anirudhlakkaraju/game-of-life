package grid

// Grid defines a grid and it's dimensions
type Grid struct {
	Width  int
	Height int
}

// NewGrid creates a grid
func NewGrid(width, height int) Grid {
	return Grid{Width: width, Height: height}
}
