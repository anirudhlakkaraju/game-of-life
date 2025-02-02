package grid

import (
	"fmt"
	"log"
)

const (
	LIVECELL = "■"
	DEADCELL = " "
)

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

// SetState sets the cells state for the grid
func (g *Grid) SetState(state [][]bool) {
	if len(state) != g.Height || len(state[0]) != g.Width {
		log.Fatalf("Given State (%d, %d) does not match Grid size (%d, %d)",
			len(state), len(state[0]), g.Height, g.Width)
	}
	g.Cells = state
}

// dirs describes all 8 directions
var dirs = [][]int{
	// Top
	{1, 1},
	{1, -1},
	{1, 0},

	// Middle
	{0, -1},
	{0, 1},

	// Bottom
	{-1, 1},
	{-1, -1},
	{-1, 0},
}

// GetNeighbors returns count of live neighboring cells
func (g *Grid) GetNeighbors(row, col int) int {
	nbhs := 0
	for _, dir := range dirs {
		r := row + dir[0]
		c := col + dir[1]

		if r >= 0 && r < g.Height && c >= 0 && c < g.Width {
			// neighboring cell is alive
			if g.Cells[r][c] == true {
				nbhs += 1
			}
		}
	}

	return nbhs
}

// Render displays the cells
func (g *Grid) Render() {
	for _, row := range g.Cells {
		for _, cell := range row {
			if cell {
				fmt.Printf(LIVECELL)
			} else {
				fmt.Printf(DEADCELL)
			}
		}
		fmt.Println()
	}
}
