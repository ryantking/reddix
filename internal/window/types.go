package window

import (
	"github.com/RyanTKing/reddix/internal/reddit"
	"github.com/RyanTKing/reddix/internal/ui/elements"
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
	Posts      *elements.Posts

	subreddit       string
	textEntryTarget string
	mode            Mode
	done            bool
}
