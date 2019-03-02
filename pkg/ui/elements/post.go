package elements

import (
	"fmt"
	"image"

	"github.com/RyanTKing/reddix/pkg/ui"
	"github.com/RyanTKing/reddix/pkg/ui/symbols"
)

// NewPost creates a new post from a geddit submission item
func NewPost(title []string, author, submitted, score string, numComments int) *Post {
	p := Post{
		Block:         *ui.NewBlock(),
		Title:         title,
		Author:        author,
		Submitted:     submitted,
		Score:         score,
		NumComments:   numComments,
		TitleStyle:    ui.Theme.Post.Title,
		SubtitleStyle: ui.Theme.Post.Subtitle,
		UpVoteStyle:   ui.Theme.Post.UpVote,
		DownVoteStyle: ui.Theme.Post.DownVote,
		ScoreStyle:    ui.Theme.Post.Score,
		LinkStyle:     ui.Theme.Post.Link,
	}
	p.Block.Border = false

	return &p
}

// Draw draws the post to the given buffer
func (p *Post) Draw(buf *ui.Buffer) {
	p.drawScore(buf)
	p.drawTitle(buf)
	p.drawSubtitle(buf)
	p.drawComments(buf)
}

func (p *Post) drawScore(buf *ui.Buffer) {
	buf.SetString(image.Pt(p.Min.X+2, p.Min.Y), string(symbols.UpVote), p.UpVoteStyle)
	buf.SetString(image.Pt(p.Min.X+2, p.Min.Y+2), string(symbols.DownVote), p.DownVoteStyle)
	scoreOff := p.Min.X + (5-len(p.Score))/2
	buf.SetString(image.Pt(scoreOff, p.Min.Y+1), p.Score, p.ScoreStyle)
}

func (p *Post) drawTitle(buf *ui.Buffer) {
	pt := image.Pt(p.Min.X+6, p.Min.Y)
	for _, row := range p.Title {
		buf.SetString(pt, row, p.TitleStyle)
		pt.Y++
	}
}

func (p *Post) drawSubtitle(buf *ui.Buffer) {
	pt := image.Pt(p.Min.X+6, p.Min.Y+len(p.Title))
	submitted := fmt.Sprintf("submitted %s ago by", p.Submitted)
	buf.SetString(pt, submitted, p.SubtitleStyle)
	pt.X += len(submitted) + 1
	buf.SetString(pt, p.Author, p.LinkStyle)
	if p.Subreddit != "" {
		pt.X += len(p.Author) + 1
		buf.SetString(pt, "to", p.SubtitleStyle)
		pt.X += 3
		buf.SetString(pt, fmt.Sprintf("r/%s", p.Subreddit), p.LinkStyle)
	}
}

func (p *Post) drawComments(buf *ui.Buffer) {
	comments := "1 comment"
	if p.NumComments != 1 {
		comments = fmt.Sprintf("%d comments", p.NumComments)
	}
	buf.SetString(image.Pt(p.Min.X+6, p.Min.Y+len(p.Title)+1), comments, p.SubtitleStyle)
}
