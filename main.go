package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"alt-potato/wordle-shuffler/internal/tui"

	"golang.org/x/term"
)

func main() {
	termW, termH, _ := term.GetSize(int(os.Stdout.Fd()))

	// signal handler for cmd-c exit
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
