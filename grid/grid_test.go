package grid

import (
	"testing"
)

func TestNewGrid(t *testing.T) {
	got := NewGrid(6, 4)
	want := Grid{
		Width:  6,
		Height: 4,
		Cells:  make([][]bool, 4),
	}

	for i := range want.Cells {
		want.Cells[i] = make([]bool, 6)
	}

	// test dimensions
	if got.Height != want.Height || got.Width != want.Width {
		t.Errorf("got [ht: %d, wd: %d], want [ht: %d, wd: %d]", got.Height, got.Width, want.Height, want.Width)
	}

	// test cell initialization
	for i := range got.Height {
		for j := range got.Width {
			if got.Cells[i][j] != false {
				t.Errorf("want all cells as false, got true at [%d][%d]", i, j)
			}
		}
	}
}

func TestSetState(t *testing.T) {
	t.Run("valid dimensions", func(t *testing.T) {
		grid := NewGrid(3, 3)
		state := [][]bool{
			{true, false, false},
			{true, false, true},
			{false, false, true},
		}

		grid.SetState(state)

		for i := range grid.Height {
			for j := range grid.Width {
				if state[i][j] != grid.Cells[i][j] {
					t.Errorf("cell[%d][%d]: got %v, want %v", i, j, state[i][j], grid.Cells[i][j])
				}
			}
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Oscillator - Blinker", func(t *testing.T) {
		grid := NewGrid(3, 3)
		state := [][]bool{
			{false, true, false},
			{false, true, false},
			{false, true, false},
		}
		grid.SetState(state)
		grid.Update()

		got := grid.Cells
		want := [][]bool{
			{false, false, false},
			{true, true, true},
			{false, false, false},
		}

		for i := range got {
			for j := range got[0] {
				if got[i][j] != want[i][j] {
					t.Errorf("cell[%d][%d]: got %v, want %v", i, j, got[i][j], want[i][j])
				}
			}
		}
	})

	t.Run("Still Life - Block", func(t *testing.T) {
		grid := NewGrid(4, 4)
		state := [][]bool{
			{false, false, false, false},
			{false, true, true, false},
			{false, true, true, false},
			{false, false, false, false},
		}
		grid.SetState(state)
		grid.Update()

		got := grid.Cells
		want := [][]bool{
			{false, false, false, false},
			{false, true, true, false},
			{false, true, true, false},
			{false, false, false, false},
		}

		for i := range got {
			for j := range got[0] {
				if got[i][j] != want[i][j] {
					t.Errorf("cell[%d][%d]: got %v, want %v", i, j, got[i][j], want[i][j])
				}
			}
		}
	})
}
