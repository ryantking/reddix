package elements

import (
	"fmt"
	"image"

	"github.com/RyanTKing/reddix/internal/ui"
)

// NewMenu creates a new menu element
func NewMenu(left, center, right string) *Menu {
	menu := Menu{
		Block:     *ui.NewBlock(),
		TextStyle: ui.Theme.Menu.Text,
		Left:      left,
		Center:    center,
		Right:     right,
	}
	menu.Border = false

	return &menu
}

func (m *Menu) drawLeft(buf *ui.Buffer) int {
	off := m.Min.X
	buf.SetString(image.Pt(off, m.Min.Y), fmt.Sprintf(" %s", m.Left), m.TextStyle)
	off += len(m.Left) + 1
	return off
}

func (m *Menu) drawCenter(buf *ui.Buffer, off int) int {
	if m.Size().X < len(m.Left)+len(m.Center)+len(m.Right)+4 {
		return off
	}

	sepCell := ui.NewCell(' ', ui.NewStyle(ui.ColorBlack, ui.ColorWhite))
	diff := m.Size().X - len(m.Right) - off - 1
	if diff < len(m.Center)+1 || len(m.Center) == 0 {
		for i := 0; i < diff; i++ {
			buf.SetCell(image.Pt(off+i, m.Min.Y), sepCell)
		}
		off += diff
		return off
	}

	diffLeft := m.Size().X/2 - len(m.Center)/2 - off
	if diffLeft < 0 {
		diffLeft = 0
	}
	for i := 0; i < diffLeft; i++ {
		buf.SetCell(image.Pt(off+i, m.Min.Y), sepCell)
	}
	off += diffLeft
	buf.SetString(image.Pt(m.Min.X+off, m.Min.Y), m.Center, m.TextStyle)
	off += len(m.Center)
	diffRight := diff - len(m.Center) - diffLeft
	for i := 0; i < diffRight; i++ {
		buf.SetCell(image.Pt(off+i, m.Min.Y), sepCell)
	}
	off += diffRight

	return off
}

func (m *Menu) drawRight(buf *ui.Buffer, off int) {
	buf.SetString(image.Pt(off, m.Min.Y), fmt.Sprintf("%s ", m.Right), m.TextStyle)
}

// Draw draws the menu to the buffer
func (m *Menu) Draw(buf *ui.Buffer) {
	off := m.drawLeft(buf)
	off = m.drawCenter(buf, off)
	m.drawRight(buf, off)
}
