package elements

import (
	"github.com/RyanTKing/reddix/internal/ui"
	"github.com/jzelinskie/geddit"
)

// Menu is a menu that is drawn across the screen
type Menu struct {
	ui.Block
	TextStyle *ui.Style
	Left      string
	Center    string
	Right     string
}

// Error is an error to be displayed to the GUI
type Error struct {
	ui.Block
	TextStyle *ui.Style
	Msg       string
}

// TextEntry is used to allow a user to enter text
type TextEntry struct {
	ui.Block
	TextStyle *ui.Style
	Prefix    string
	Text      string
	Hidden    bool
}

// Posts is used to draw a series of posts to the screen
type Posts struct {
	ui.Block
	Posts         []*geddit.Submission
	Selected      int
	Offset        int
	LastPost      int
	Frontpage     bool
	TitleStyle    *ui.Style
	SubtitleStyle *ui.Style
	UpVoteStyle   *ui.Style
	DownVoteStyle *ui.Style
	ScoreStyle    *ui.Style
	LinkStyle     *ui.Style
	SelectedStyle *ui.Style
}
