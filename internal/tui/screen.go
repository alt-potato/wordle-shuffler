package tui

import (
	"os"
	"strings"
)

type Cell struct {
	Rune rune
	Fg   int
	Bg   int
}

type Screen struct {
	width, height int
	cells         [][]Cell
}

func NewScreen(w, h int) *Screen {
	cells := make([][]Cell, h)
	for i := range cells {
		cells[i] = make([]Cell, w)
		for j := range cells[i] {
			cells[i][j] = Cell{Rune: ' ', Fg: -1, Bg: -1}
		}
	}

	return &Screen{w, h, cells}
}

// sets a specific cell coordinate, ignoring out-of-bounds coordinates.
func (s *Screen) SetCell(r, c int, ch rune, fg, bg int) {
	if r >= 0 && r < s.height && c >= 0 && c < s.width {
		s.cells[r][c] = Cell{Rune: ch, Fg: fg, Bg: bg}
	}
}

// returns a shallow copy of the cell at the given coordinates, 
// or an empty cell if out-of-bounds.
func (s *Screen) Cell(r, c int) Cell {
	if r >= 0 && r < s.height && c >= 0 && c < s.width {
		return s.cells[r][c]
	} else {
		return Cell{}
	}
}

func (s *Screen) Render() {
	var out strings.Builder
	out.WriteString(clearScreen + home)

	for row := 0; row < s.height; row++ {
		for col := 0; col < s.width; col++ {
			out.WriteRune(s.cells[row][col].Rune)
		}

		// no newline on last line (causes cutoff of first line)
		if row < s.height - 1 {
			out.WriteString("\r\n")
		}
	}

	out.WriteString(resetModes) // reset attributes
	os.Stdout.WriteString(out.String())
}

// Same as Render(), but only renders the per-cell differences from the last screen.
func (s *Screen) RenderDelta(prev *Screen) {
	var out strings.Builder
	out.WriteString(home) // no clear

	// track cursor position for jumps
	cursorRow, cursorCol := 0, 0

	for row := 0; row < s.height; row++ {
		for col := 0; col < s.width; col++ {
			currVal := s.cells[row][col]
			prevVal := prev.cells[row][col]

			if currVal == prevVal {
				continue // skip
			}

			// can't assume cursor is in the right place since skips occur
			if row != cursorRow || col != cursorRow {
				out.WriteString(moveCursor(row, col))
			}

			out.WriteRune(currVal.Rune)
			cursorCol++
		}

		// no newline on last line (causes cutoff of first line)
		if row < s.height - 1 {
			out.WriteString("\r\n")
		}
	}

	out.WriteString(resetModes) // reset attributes
	os.Stdout.WriteString(out.String())
}

// Creates a deep copy of the current Screen.
func (s *Screen) Clone() *Screen {
	clone := NewScreen(s.width, s.height)
	for row := 0; row < s.height; row++ {
		copy(clone.cells[row], s.cells[row])
	}

	return clone
}

// Clears the current Screen. 
func (s *Screen) Clear() {
	for i := range s.cells {
		for j := range s.cells[i] {
			s.cells[i][j] = Cell{Rune: ' ', Fg: -1, Bg: -1}
		}
	}
}
