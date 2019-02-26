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
	TextEntry  *elements.TextEntry
	Err        *elements.Error

	posts           []*geddit.Submission
	postOffset      int
	lastPost        int
	selected        int
	subreddit       string
	textEntryTarget string
	mode            Mode
	done            bool
}
