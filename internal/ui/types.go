package ui

import (
	"image"
)

// Color represents an xterm color (0-255)
type Color int

// Modifier is a text modifier such as bold or underline
type Modifier uint

// Style holds information about the style for an object
type Style struct {
	FG       Color
	BG       Color
	Modifier Modifier
}

// Cell is a single cell of the terminal
type Cell struct {
	Rune  rune
	Style *Style
}

// Buffer is a section of the terminal that can be drawn to
type Buffer struct {
	image.Rectangle
	Cells map[image.Point]*Cell
}

// Drawable is anything that can be drawn onto a buffer
type Drawable interface {
	GetRect() image.Rectangle
	SetRect(x1, y1, x2, y2 int)
	Draw(*Buffer)
}

// Block is a basic UI block that can have borders and a title
type Block struct {
	Border       bool
	BorderStyle  *Style
	BorderLeft   bool
	BorderRight  bool
	BorderTop    bool
	BorderBottom bool

	image.Rectangle
	Inner image.Rectangle

	Title      string
	TitleStyle *Style
}

// RootTheme represents the them of the entire application
type RootTheme struct {
	Default *Style

	Block BlockTheme

	Menu      MenuTheme
	TextEntry TextEntryTheme
	Error     ErrorTheme
	Posts     PostsTheme
}

// BlockTheme represents the theme used for a block
type BlockTheme struct {
	Title  *Style
	Border *Style
}

// MenuTheme represents the them used for a menu
type MenuTheme struct {
	Text *Style
}

// TextEntryTheme represents the them used for text entry
type TextEntryTheme struct {
	Text *Style
}

// ErrorTheme represents the theme used for errors
type ErrorTheme struct {
	Text *Style
}

// PostsTheme represents the theme for posts
type PostsTheme struct {
	Title    *Style
	Subtitle *Style
	UpVote   *Style
	DownVote *Style
	Score    *Style
	Link     *Style
	Selected *Style
}
