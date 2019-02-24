package ui

import (
	"image"

	"github.com/RyanTKing/reddix/internal/ui/symbols"
)

// NewBlock creates an empty block
func NewBlock() *Block {
	block := Block{
		Border:       true,
		BorderStyle:  Theme.Block.Border,
		BorderLeft:   true,
		BorderRight:  true,
		BorderTop:    true,
		BorderBottom: true,

		TitleStyle: Theme.Block.Title,
	}

	return &block
}

func (b *Block) drawBorder(buf *Buffer) {
	vertCell := NewCell(symbols.VerticalLine, b.BorderStyle)
	horCell := NewCell(symbols.HorizontalLine, b.BorderStyle)

	if b.BorderTop {
		buf.Fill(image.Rect(b.Min.X, b.Min.Y, b.Max.X, b.Min.Y+1), horCell)
	}
	if b.BorderBottom {
		buf.Fill(image.Rect(b.Min.X, b.Max.Y-1, b.Max.X, b.Max.Y), horCell)
	}
	if b.BorderLeft {
		buf.Fill(image.Rect(b.Min.X, b.Min.Y, b.Min.X+1, b.Max.Y), vertCell)
	}
	if b.BorderTop {
		buf.Fill(image.Rect(b.Max.X-1, b.Min.Y, b.Max.X, b.Max.Y), vertCell)
	}

	if b.BorderTop && b.BorderLeft {
		buf.SetCell(b.Min, NewCell(symbols.WindowTopLeft, b.BorderStyle))
	}
	if b.BorderTop && b.BorderRight {
		buf.SetCell(image.Pt(b.Max.X-1, b.Min.Y), NewCell(symbols.WindowTopRight, b.BorderStyle))
	}
	if b.BorderBottom && b.BorderLeft {
		buf.SetCell(image.Pt(b.Min.X, b.Max.Y-1), NewCell(symbols.WindowBottomLeft, b.BorderStyle))
	}
	if b.BorderBottom && b.BorderRight {
		buf.SetCell(b.Max.Sub(image.Pt(1, 1)), NewCell(symbols.WindowBottomRight, b.BorderStyle))
	}
}

// Draw draws the block to a buffe3r
func (b *Block) Draw(buf *Buffer) {
	if b.Border {
		b.drawBorder(buf)
	}
	buf.SetString(image.Pt(b.Min.X+2, b.Min.Y), b.Title, b.TitleStyle)
}

// SetRect sets the rectangle bounds for the block
func (b *Block) SetRect(x1, y1, x2, y2 int) {
	b.Rectangle = image.Rect(x1, y1, x2, y2)
	b.Inner = image.Rect(b.Min.X+1, b.Min.Y+1, b.Max.X-1, b.Max.Y-1)
}

// GetRect returns the rectangle the block inhabits
func (b *Block) GetRect() image.Rectangle {
	return b.Rectangle
}
