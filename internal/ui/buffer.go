package ui

import "image"

// NewBuffer returns a new buffer made from a rectangle
func NewBuffer(r image.Rectangle) *Buffer {
	buf := Buffer{
		Rectangle: r,
		Cells:     make(map[image.Point]*Cell),
	}

	buf.Fill(r, CellClear)
	return &buf
}

// GetCell returns a specific cell of the buffer
func (b *Buffer) GetCell(p image.Point) *Cell {
	return b.Cells[p]
}

// SetCell sets a specifc cell to a given one
func (b *Buffer) SetCell(p image.Point, c *Cell) {
	b.Cells[p] = c
}

// Fill fills a rectangle in the buffer with a given cell in the rectangle
func (b *Buffer) Fill(r image.Rectangle, c *Cell) {
	for x := r.Min.X; x < r.Max.X; x++ {
		for y := r.Min.Y; y < r.Max.Y; y++ {
			b.SetCell(image.Pt(x, y), c)
		}
	}
}

// SetString writes a string to the buffer starting from a given point
func (b *Buffer) SetString(p image.Point, s string, style *Style) {
	for i, c := range s {
		b.SetCell(image.Pt(p.X+i, p.Y), NewCell(c, style))
	}
}
