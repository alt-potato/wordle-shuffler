package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"alt-potato/wordle-shuffler/internal/tui"

	"golang.org/x/term"
)

func main() {
	termW, termH, _ := term.GetSize(int(os.Stdout.Fd()))

	// signal handeler for cmd-c exit
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	screen := tui.NewScreen(termW, termH)

	// start tick loop
	const fps = 12
	ticker := time.NewTicker(time.Second / fps)
	defer ticker.Stop()

	frameCount := 0
	// startTime := time.Now()

main:
	for {
		select {
		case <-sigCh:
			break main
		case <-ticker.C:
			frameCount++

			setDemo(screen, frameCount)

			screen.Render()
		}
	}
}

func setDemo(screen *tui.Screen, tick int) {
	// probably a better way to do this but whatever
	type coord struct {
		row, col int
	}

	vectors := [...]struct {
		facing, length int // facing: NESW -> 0123 (too lazy to define in code)
	}{
		{1, 8},
		{2, 3},
		{3, 8},
		{0, 3},
	}

	pathLength := 0
	for _, v := range vectors {
		pathLength += v.length
	}

	path := make([]coord, pathLength)

	// build path (starting position is NOT counted (easier that way lol))
	cursor := coord{2, 1}
	curr := 0
	for _, v := range vectors {
	draw:
		for range v.length {
			switch v.facing {
			case 0:
				cursor.row -= 1 // north
			case 1:
				cursor.col += 1 // east
			case 2:
				cursor.row += 1 // south
			case 3:
				cursor.col -= 1 // west
			default:
				pathLength -= 1 // don't count invalid instruction
				break draw // don't draw!
				// should probably break out before the draw loop but whatever
			}

			path[curr] = cursor
			curr++
		}
	}

	snake := "******"
	snake += strings.Repeat(" ", len(path)-len(snake)) // match length of path bc i'm lazy

	pos := tick % len(path)
	for _, char := range snake {
		pos = (pos + 1) % len(path)
		screen.SetCell(path[pos].row, path[pos].col, char, -1, -1)
	}
}
