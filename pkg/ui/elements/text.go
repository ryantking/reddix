package elements

import (
	"image"

	"github.com/RyanTKing/reddix/pkg/ui"
)

// NewText creates a new text element
func NewText(s string) *Text {
	txt := Text{
		Block:     *ui.NewBlock(),
		Text:      s,
		TextStyle: ui.Theme.Text.Text,
	}

	return &txt
}

// NewError creates a new text element with error formatting
func NewError(s string) *Text {
	txt := Text{
		Block:     *ui.NewBlock(),
		Text:      s,
		TextStyle: ui.Theme.Text.Error,
	}

	return &txt
}

// Draw draws the text to the screen
func (txt *Text) Draw(buf *ui.Buffer) {
	buf.SetString(image.Pt(txt.Min.X, txt.Min.Y), txt.Text, txt.TextStyle)
}
