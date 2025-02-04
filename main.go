package main

import (
	"game-of-life/grid"
	"time"

	"github.com/rivo/tview"
)

func main() {

	width := 80
	height := 24
	grid := grid.NewGrid(width, height)
	state := getInitialState(width, height)
	grid.SetState(state)

	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetText(grid.Render()).
		SetTextAlign(tview.AlignCenter)

	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			grid.Update()

			app.QueueUpdateDraw(func() {
				textView.SetText(grid.Render())
			})
		}
	}()

	// Run the app full-screen
	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}

}

func getInitialState(width, height int) [][]bool {
	// Create an empty grid
	state := make([][]bool, height)
	for i := range state {
		state[i] = make([]bool, width)
	}

	// Set the initial cells as per your given pattern
	state[12][40] = true
	state[12][41] = true
	state[13][40] = true
	state[13][41] = true
	state[14][38] = true
	state[14][39] = true
	state[14][40] = true
	state[15][39] = true

	return state
}
