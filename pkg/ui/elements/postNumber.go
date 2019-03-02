package elements

import (
	"fmt"
	"image"

	"github.com/RyanTKing/reddix/pkg/ui"
)

// NewPostNumber creates a new post number from the given number
func NewPostNumber(num int) *PostNumber {
	pn := PostNumber{
		Block:    *ui.NewBlock(),
		Num:      num,
		NumStyle: ui.Theme.PostNumber.Num,
	}
	pn.Block.Border = false

	return &pn
}

// Draw draws the post number to the buffer
func (pn *PostNumber) Draw(buf *ui.Buffer) {
	num := fmt.Sprint(pn.Num)
	buf.SetString(image.Pt(pn.Min.X, pn.Min.Y+1), num, pn.NumStyle)
}
