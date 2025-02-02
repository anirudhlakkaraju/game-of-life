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

func TestGetNeighbors(t *testing.T) {
	grid := NewGrid(3, 3)
	state := [][]bool{
		{true, false, false},
		{true, false, true},
		{false, false, true},
	}
	grid.SetState(state)

	t.Run("center of grid", func(t *testing.T) {
		got := grid.GetNeighbors(1, 1)
		want := 4

		if got != want {
			t.Errorf("got %d live neighbors, want %d", got, want)
		}
	})

	t.Run("boundary of grid", func(t *testing.T) {
		got := grid.GetNeighbors(0, 0)
		want := 1

		if got != want {
			t.Errorf("got %d live neighbors, want %d", got, want)
		}
	})
}
