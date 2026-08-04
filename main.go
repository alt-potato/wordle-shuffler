package main

import (
	"fmt"
	"os"

	"alt-potato/wordle-shuffler/internal/tui"

	"golang.org/x/term"
)

func main() {
	termW, termH, _ := term.GetSize(int(os.Stdout.Fd()))
	screen := tui.NewScreen(termW, termH)

	screen.SetCell(0, 0, 'h', -1, -1)
	screen.SetCell(1, 1, 'i', -1, -1)
	screen.SetCell(3, 3, 'c', -1, -1)
	screen.SetCell(4, 4, 'h', -1, -1)
	screen.SetCell(5, 5, 'a', -1, -1)
	screen.SetCell(6, 6, 't', -1, -1)

	screen.Render()

	fmt.Scanln()
}
