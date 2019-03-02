package reddix

import (
	"github.com/RyanTKing/reddix/pkg/session"
	"github.com/RyanTKing/reddix/pkg/view"
	"github.com/RyanTKing/reddix/pkg/window"
)

// Mode represents a mode the reddix application can be in
type Mode int

// Reddix holds the state of the application
type Reddix struct {
	mode Mode

	win  window.Window
	sess *session.Session

	topMenu    *Menu
	bottomMenu *Menu
	miniBuffer *MiniBuffer

	browse *view.Browse
}

// Menu holds the values in the menu
type Menu struct {
	Y      int
	Width  int
	Left   string
	Center string
	Right  string
}

// MiniBuffer is a buffer at the bottom of the screen for messages and text entry
type MiniBuffer struct {
	Y     int
	Width int

	InputMode   bool
	Target      string
	Input       []string
	HiddenInput bool

	Message string
	Error   bool
}
