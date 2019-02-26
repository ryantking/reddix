package elements

import (
	"image"

	"github.com/RyanTKing/reddix/internal/ui"
)

// NewText creates a new text element
func NewText() *Text {
	txt := Text{
		Block:      *ui.NewBlock(),
		TextStyle:  ui.Theme.Text.Text,
		ErrorStyle: ui.Theme.Text.Error,
	}
	txt.Block.Border = false

	return &txt
}

// Draw draws the text to the screen
func (txt *Text) Draw(buf *ui.Buffer) {
	buf.Fill(txt.GetRect(), ui.CellClear)
	style := txt.TextStyle
	if txt.Error {
		style = txt.ErrorStyle
	}
	buf.SetString(image.Pt(txt.Min.X, txt.Min.Y), txt.Text, style)
}
