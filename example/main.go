package main

import (
	"fmt"

	"github.com/thxrsxm/rnbw"
)

func rainbow() {
	lines := []string{
		"▗▄▄▖  ▗▄▖ ▗▄▄▄▖▗▖  ▗▖▗▄▄▖  ▗▄▖ ▗▖ ▗▖",
		"▐▌ ▐▌▐▌ ▐▌  █  ▐▛▚▖▐▌▐▌ ▐▌▐▌ ▐▌▐▌ ▐▌",
		"▐▛▀▚▖▐▛▀▜▌  █  ▐▌ ▝▜▌▐▛▀▚▖▐▌ ▐▌▐▌ ▐▌",
		"▐▌ ▐▌▐▌ ▐▌▗▄█▄▖▐▌  ▐▌▐▙▄▞▘▝▚▄▞▘▐▙█▟▌",
	}
	colors := []rnbw.Color{
		rnbw.Red,
		rnbw.Yellow,
		rnbw.Green,
		rnbw.Blue,
	}
	for i := range lines {
		rnbw.Println(colors[i], lines[i])
	}
}

func main() {
	// Banner
	rainbow()
	fmt.Println()
	// Foreground colors
	rnbw.Println(rnbw.Green, "Success!")
	rnbw.Println(rnbw.Red, "Error!")
	rnbw.Printf(rnbw.Blue, "Hello, %s!\n", "World")
	fmt.Println()
	// Background colors
	rnbw.BgPrintln(rnbw.Yellow, "Warning background")
	fmt.Println()
	// Combined foreground and background
	rnbw.StyledPrintln(rnbw.Gray, rnbw.Red, "Gray text on red background")
	fmt.Println()
	rnbw.StyledPrintf(rnbw.Yellow, rnbw.Blue, "Score: %d", 100)
	fmt.Println()
}
