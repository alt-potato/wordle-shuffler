package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"alt-potato/wordle-shuffler/internal/tui"
)

// heh. mango.
func main() {
	termState, err := tui.NewTerminal()
	if err != nil {
		fmt.Println("Could not create terminal: " + err.Error())
	}
	defer termState.Close()

	termW, termH, _ := termState.TermSize()

	// signal handler for forced exit
	// probably won't capture ctrl-c, since we're in raw mode
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// stdin read and ctrl-c routing
	tui.SetupInput(sigCh)

	screen := tui.NewScreen(termW, termH)
	prevScreen := screen.Clone()

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
			// render!
			frameCount++

			setDemo(screen, frameCount)

			screen.RenderDelta(prevScreen)
			screen.Swap(prevScreen)
		}
	}
}
