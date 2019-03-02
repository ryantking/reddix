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
	Fill         bool
	FillStyle    *Style
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

	Menu       MenuTheme
	Text       TextTheme
	Post       PostTheme
	PostNumber PostNumberTheme
}

// BlockTheme represents the theme used for a block
type BlockTheme struct {
	Title  *Style
	Fill   *Style
	Border *Style
}

// MenuTheme represents the them used for a menu
type MenuTheme struct {
	Text *Style
}

// TextTheme represents the them used for text entry
type TextTheme struct {
	Text  *Style
	Error *Style
}

// PostTheme represents the theme for a post
type PostTheme struct {
	Title    *Style
	Subtitle *Style
	UpVote   *Style
	DownVote *Style
	Score    *Style
	Link     *Style
}

// PostNumberTheme represents the theme for a post number
type PostNumberTheme struct {
	Num *Style
}
