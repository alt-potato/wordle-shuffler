package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	termW, termH, _ := term.GetSize(int(os.Stdout.Fd()))
	screen := NewScreen(termW, termH)

	screen.Set(0, 0, 'h', -1, -1)
	screen.Set(1, 1, 'i', -1, -1)
	screen.Set(3, 3, 'c', -1, -1)
	screen.Set(4, 4, 'h', -1, -1)
	screen.Set(5, 5, 'a', -1, -1)
	screen.Set(6, 6, 't', -1, -1)

	screen.Render()

	fmt.Scanln()
}
