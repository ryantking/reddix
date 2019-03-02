package window

import (
	"github.com/RyanTKing/reddix/pkg/event"
	"github.com/RyanTKing/reddix/pkg/ui"
)

// Window represents a window that can be drawn to by the Reddix instance
type Window interface {
	// Size returns the size of the window
	Size() (int, int)

	// Poll is a blocking call that gets the next event
	Poll() event.Event

	// Draw is used to render a drawable UI element to the window
	Draw(ui.Drawable)

	// SetCursor sets the cursor location
	SetCursor(x, y int)

	// Refresh refreshes the window
	Refresh() error

	// Quit sends a signal to end the event loop
	Quit()
}
