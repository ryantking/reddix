package ui

// CellClear represents an empty cell
var CellClear = &Cell{
	Rune:  ' ',
	Style: StyleClear,
}

// NewCell creates a new cell with an optional style
func NewCell(ch rune, args ...interface{}) *Cell {
	cell := Cell{
		Rune:  ch,
		Style: StyleClear,
	}
	if len(args) == 1 {
		cell.Style = args[0].(*Style)
	}

	return &cell
}
