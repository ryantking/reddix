package ui

const (
	// ColorClear clears the foreground or background of a style
	ColorClear Color = iota - 1
	// ColorBlack represents the black color (0)
	ColorBlack
	// ColorRed represents the red color (1)
	ColorRed
	// ColorGreen represents the green color (2)
	ColorGreen
	// ColorYellow represents the yellow color (3)
	ColorYellow
	// ColorBlue represents the blue color (4)
	ColorBlue
	// ColorMagenta represents the magenta color (5)
	ColorMagenta
	// ColorCyan represents the cyan color (6)
	ColorCyan
	// ColorWhite represents the white color (7)
	ColorWhite

	// ModifierClear removes any modifiers
	ModifierClear Modifier = 0
	// ModifierBold makes the text bold
	ModifierBold Modifier = 1 << 9
	//ModifierUnderline underlines the text
	ModifierUnderline Modifier = 1 << 10
	// ModifierReverse reverses the text
	ModifierReverse Modifier = 1 << 11
)

// StyleClear clears all stypes
var StyleClear = &Style{
	FG:       ColorClear,
	BG:       ColorClear,
	Modifier: ModifierClear,
}

// NewStyle creates a new style from the given foreground cover and optionally a background color and modifier
// The second arg will be used as the background color and the third arg will be used as he modifier
func NewStyle(fg Color, args ...interface{}) *Style {
	style := Style{FG: fg}
	if len(args) >= 1 {
		style.BG = args[0].(Color)
	}
	if len(args) == 2 {
		style.Modifier = args[1].(Modifier)
	}

	return &style
}
