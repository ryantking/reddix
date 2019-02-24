package elements

import (
	"fmt"
	"image"

	"github.com/RyanTKing/reddix/internal/ui"
)

// NewTextEntry creates a new text entry element
func NewTextEntry(prefix string, hidden bool) *TextEntry {
	textEntry := TextEntry{
		Block:     *ui.NewBlock(),
		TextStyle: ui.Theme.TextEntry.Text,
		Prefix:    prefix,
		Hidden:    hidden,
	}
	textEntry.Border = false

	return &textEntry
}

// Draw draws the text entry field to a buffer
func (te *TextEntry) Draw(buf *ui.Buffer) {
	text := te.Text
	if te.Hidden {
		text = ""
		for i := 0; i < len(te.Text); i++ {
			text += "*"
		}
	}
	if te.Prefix != "" {
		text = fmt.Sprintf("%s: %s", te.Prefix, text)
	}
	buf.SetString(image.Pt(te.Min.X, te.Min.Y), text, te.TextStyle)
	for i := te.Min.X + len(text); i < te.Max.X; i++ {
		buf.SetCell(image.Pt(i, te.Max.Y), ui.CellClear)
	}
}
