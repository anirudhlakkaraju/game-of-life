package grid

import (
	"log"
	"strings"
)

const (
	LIVECELL  = "◼"
	DEADCELL  = " "
	LineBreak = "\n"
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

// Render displays the cells
func (g *Grid) Render() string {
	var out strings.Builder

	for _, row := range g.Cells {
		for _, cell := range row {
			if cell {
				out.WriteString(LIVECELL)
			} else {
				out.WriteString(DEADCELL)
			}
		}
		out.WriteString(LineBreak)
	}

	return out.String()
}

// Update applies Conway's rules
// 1. Any live cell with fewer than two live neighbors dies (underpopulation).
// 2. Any live cell with two or three live neighbors survives.
// 3. Any live cell with more than three live neighbors dies (overpopulation).
// 4. Any dead cell with exactly three live neighbors becomes alive (reproduction).
func (g *Grid) Update() {
	// Initialize all cells as dead first
	newCells := make([][]bool, g.Height)
	for i := range newCells {
		newCells[i] = make([]bool, g.Width)
	}

	for i := range g.Height {
		for j := range g.Width {
			newCells[i][j] = g.Cells[i][j]
			nbhs := getNeighbors(i, j, g.Cells)
			if nbhs < 2 {
				newCells[i][j] = false
			}
			if nbhs >= 4 {
				newCells[i][j] = false
			}
			if nbhs == 3 {
				newCells[i][j] = true
			}
		}
	}

	g.Cells = newCells
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

// getNeighbors returns nbhs of live neighboring cells
func getNeighbors(row, col int, state [][]bool) int {
	nbhs := 0
	for _, dir := range dirs {
		r := row + dir[0]
		c := col + dir[1]

		if r >= 0 && r < len(state) && c >= 0 && c < len(state[0]) {
			// neighboring cell is alive
			if state[r][c] == true {
				nbhs += 1
			}
		}
	}

	return nbhs
}
