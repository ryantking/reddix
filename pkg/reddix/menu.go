package reddix

import (
	"github.com/RyanTKing/reddix/pkg/ui"
	"github.com/RyanTKing/reddix/pkg/ui/elements"
	"github.com/RyanTKing/reddix/pkg/window"
)

const (
	frontpage = "frontpage"

	topMenuLeft     = "q:quit r:subreddit u:user h:help"
	topMenuCenter   = "reddix"
	topMenuRight1   = "l:login"
	topMenuRight2   = "l:logout"
	bottomMenuRight = "Not Logged In"
)

var menuStyle = ui.NewStyle(ui.ColorBlack, ui.ColorWhite)

// NewMenu creates a new menu in the given location
func NewMenu(y, width int, left, center, right string) *Menu {
	menu := Menu{
		Y:      y,
		Width:  width,
		Left:   left,
		Center: center,
		Right:  right,
	}

	return &menu
}

// Resize Handles window resizes and redrawing
func (m *Menu) Resize(win window.Window, y, width int) {
	if y == m.Y && width == m.Width {
		return
	}

	m.Y = y
	m.Width = width
	m.Draw(win)
}

// Draw draws the menu to the window
func (m *Menu) Draw(win window.Window) {
	m.fill(win)
	m.drawLeft(win)
	m.drawCenter(win)
	m.drawRight(win)
	win.Refresh()
}

func (m *Menu) fill(win window.Window) {
	block := ui.NewBlock()
	block.FillStyle = menuStyle
	block.SetRect(0, m.Y, m.Width, m.Y+1)
	win.Draw(block)
}

func (m *Menu) drawLeft(win window.Window) {
	text := elements.NewText(m.Left)
	text.TextStyle = menuStyle
	text.SetRect(1, m.Y, len(m.Left)+1, m.Y+1)
	win.Draw(text)
}

func (m *Menu) drawCenter(win window.Window) {
	if m.Center == "" || len(m.Center) > m.Width-len(m.Left)-len(m.Right)-4 {
		return
	}

	x := m.Width/2 - len(m.Center)/2 + len(m.Center)%2
	if x < len(m.Left)+2 {
		x = len(m.Left) + 2
	}

	text := elements.NewText(m.Center)
	text.TextStyle = menuStyle
	text.SetRect(x, m.Y, x+len(m.Center), m.Y+1)
	win.Draw(text)
}

func (m *Menu) drawRight(win window.Window) {
	if len(m.Right) > m.Width-len(m.Left)-3 {
		return
	}

	x := m.Width - len(m.Right) - 1
	text := elements.NewText(m.Right)
	text.TextStyle = menuStyle
	text.SetRect(x, m.Y, m.Width-1, m.Y+1)
	win.Draw(text)
}
