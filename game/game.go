package game

import (
	"fmt"
	g "game-of-life/grid"
	"time"
)

func Start(grid *g.Grid) {
	timer := time.NewTicker(100 * time.Millisecond)
	for {
		<-timer.C
		clearScreen()
		grid.Render()
		grid.Update()
	}
}

func clearScreen() {
	// Hide cursor before clearing screen
	fmt.Print("\033[?25l")

	// Clear screen
	fmt.Print("\033[H\033[2J")

	// Show cursor after clearing screen
	fmt.Print("\033[?25h")
}
