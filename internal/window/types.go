package window

import (
	"github.com/RyanTKing/reddix/internal/reddit"
	"github.com/RyanTKing/reddix/internal/ui/elements"
	"github.com/jzelinskie/geddit"
)

// Mode is the current window mode
type Mode int

// Window holds information about the current reddit window
type Window struct {
	Width  int
	Height int
	Sess   *reddit.Session

	TopMenu    *elements.Menu
	BottomMenu *elements.Menu

	miniBuffer      *MiniBuffer
	posts           []*geddit.Submission
	postOffset      int
	lastPost        int
	selected        int
	subreddit       string
	textEntryTarget string
	mode            Mode
	done            bool
}

// MiniBuffer is the buffer at the bottom of the screen for messages and text entry
type MiniBuffer struct {
	Text *elements.Text

	InputMode   bool
	Target      string
	Input       string
	HiddenInput bool

	Message string
	Error   bool
}
