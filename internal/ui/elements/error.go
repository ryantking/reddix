package elements

import (
	"fmt"
	"image"

	"github.com/RyanTKing/reddix/internal/ui"
)

// NewError creates a new error element
func NewError(msg string) *Error {
	err := Error{
		Block:     *ui.NewBlock(),
		TextStyle: ui.Theme.Error.Text,
		Msg:       msg,
	}
	err.Border = false

	return &err
}

// Draw draws the error message to the buffer
func (err *Error) Draw(buf *ui.Buffer) {
	buf.SetString(image.Pt(err.Min.X, err.Min.Y), fmt.Sprintf("error: %s", err.Msg), err.TextStyle)
}
