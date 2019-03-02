package view

import (
	"github.com/RyanTKing/reddix/pkg/event"
	"github.com/RyanTKing/reddix/pkg/session"
	"github.com/RyanTKing/reddix/pkg/window"
	"github.com/jzelinskie/geddit"
)

// View is a reddix view that holds UI element, and respond to various events
type View interface {
	// Rect returns the top left and bottom right coordinate of the view
	Rect() (int, int, int, int)

	// HandleKey handles a keyboard event
	HandleKey(sess *session.Session, win window.Window, key *event.Keyboard) error

	// HandleResize handles a resize event
	HandleResize(win window.Window, width, height int)
}

// Browse is a view that allows the browsing of a subreddit
type Browse struct {
	x      int
	y      int
	width  int
	height int

	posts      []*geddit.Submission
	selected   int
	postOffset int
	lastPost   int
}
