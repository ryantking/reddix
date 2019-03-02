package elements

import (
	"github.com/RyanTKing/reddix/pkg/ui"
)

// Menu is a menu that is drawn across the screen
type Menu struct {
	ui.Block
	TextStyle *ui.Style
	Left      string
	Center    string
	Right     string
}

// Text is a text field
type Text struct {
	ui.Block
	Text      string
	TextStyle *ui.Style
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
