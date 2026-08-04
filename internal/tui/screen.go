package tui

import (
	"os"
	"strings"
)

// ref: https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797
const (
	esc         = "\x1B"
	clearScreen = esc + "[2J"
	home        = esc + "[H"
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

	os.Stdout.WriteString(out.String())
}
