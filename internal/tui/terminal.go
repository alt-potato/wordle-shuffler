package tui

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

// ref: https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797
const (
	ETX = 0x03
	ESC = 0x1B

	clearScreen = "\x1B[2J"
	home        = "\x1B[H"
	resetModes  = "\x1B[0m"

	hideCursor       = "\x1B[?25l"
	showCursor       = "\x1B[?25h"
	enableAltBuffer  = "\x1B[?1049h"
	disableAltBuffer = "\x1B[?1049l"
)

type TerminalState struct {
	fd int
	oldTermState *term.State
}

// get terminal and enter raw mode with alternate screen buffer
func NewTerminal() (*TerminalState, error) {
	fd := int(os.Stdout.Fd())

	// save state and enter raw mode
	oldState, err := term.MakeRaw(fd)
    if err != nil {
        return nil, err
    }

	// enter alternate screen buffer and hide cursor
	fmt.Print(enableAltBuffer + hideCursor)

	return &TerminalState{fd, oldState}, nil
}

// restore previous terminal
func (t *TerminalState) Close() {
	// show cursor and restore original buffer
	fmt.Print(showCursor + disableAltBuffer)
	term.Restore(t.fd, t.oldTermState)
}

func (t *TerminalState) TermSize() (int, int, error) {
	return term.GetSize(t.fd)
}

// buffered async reader for capturing keyboard input
func SetupInput(sigCh chan os.Signal) {
	go func() {
		buf := make([]byte, 32) // input buffer

		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				continue // who cares about errors amirite
			}

			for i := range n {
				switch buf[i] {
				case ETX, ESC:
					// leaves
					sigCh <- syscall.SIGINT
				}
			}
		}
	}()
}
