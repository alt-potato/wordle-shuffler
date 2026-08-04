package tui

import "testing"

// sanity test for new screen creation
func TestNewScreen(t *testing.T) {
	s := NewScreen(80, 24)
	if s.width != 80 || s.height != 24 {
		t.Errorf("NewScreen(80, 24) wrong size: %dx%d", s.width, s.height)
	}

	if len(s.cells) != 24 || len(s.cells[0]) != 80 {
		t.Error("Cell slice dimensions incorrect")
	}
}

// sanity check for new screen initial content
func TestInitialContent(t *testing.T) {
	s := NewScreen(10, 3)

	for r := 0; r < s.height; r++ {
		for c := 0; c < s.width; c++ {
			cell := s.cells[r][c]
			if cell.Rune != ' ' {
				t.Errorf("Expected space at (%d,%d), got %q", r, c, cell.Rune)
			}
            if cell.Fg != -1 || cell.Bg != -1 {
                t.Errorf("Expected default attrs at (%d,%d), got Fg=%d Bg=%d", 
                    r, c, cell.Fg, cell.Bg)
            }
		}
	}
}

// ensure oob writes are ignored and do not panic
func TestOutOfBounds(t *testing.T) {
	s := NewScreen(10, 5)

	s.Set(-1, 0, 'X', -1, -1)  // row too low
    s.Set(5, 0, 'X', -1, -1)   // row too high
    s.Set(0, -1, 'X', -1, -1)  // col too low
    s.Set(0, 10, 'X', -1, -1)  // col too high

	for r := 0; r < s.height; r++ {
		for c := 0; c < s.width; c++ {
			cell := s.cells[r][c]
			if cell.Rune != ' ' {
				t.Error("Out-of-bounds set affected valid cell (somehow)")
			}
		}
	}
}

// confirm that UTF-8 characters can be held and (presumably) rendered
func TestUnicodeHandling(t *testing.T) {
	s := NewScreen(3, 3)

	s.Set(1, 1, 'é', -1, -1) // multi-byte UTF-8 character

	cell := s.cells[1][1]
	if cell.Rune != 'é' {
		t.Errorf("Expected 'é', got %q", cell.Rune)
	}
}
