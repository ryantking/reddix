package elements

import (
	"github.com/RyanTKing/reddix/internal/ui"
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

// Post is an reddit post
type Post struct {
	ui.Block
	Title         []string
	Author        string
	Submitted     string
	Score         string
	NumComments   int
	Subreddit     string
	TitleStyle    *ui.Style
	SubtitleStyle *ui.Style
	UpVoteStyle   *ui.Style
	DownVoteStyle *ui.Style
	ScoreStyle    *ui.Style
	LinkStyle     *ui.Style
}

// PostNumber is the number label of a post
type PostNumber struct {
	ui.Block
	Num      int
	NumStyle *ui.Style
}
