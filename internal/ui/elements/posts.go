package elements

import (
	"fmt"
	"image"
	"time"

	"github.com/RyanTKing/reddix/internal/ui/symbols"

	"github.com/RyanTKing/reddix/internal/ui"
	"github.com/jzelinskie/geddit"
)

// NewPosts creates a new posts element
func NewPosts(posts []*geddit.Submission) *Posts {
	p := Posts{
		Block:         *ui.NewBlock(),
		Posts:         posts,
		TitleStyle:    ui.Theme.Posts.Title,
		SubtitleStyle: ui.Theme.Posts.Subtitle,
		UpVoteStyle:   ui.Theme.Posts.UpVote,
		DownVoteStyle: ui.Theme.Posts.DownVote,
		ScoreStyle:    ui.Theme.Posts.Score,
		LinkStyle:     ui.Theme.Posts.Link,
		SelectedStyle: ui.Theme.Posts.Selected,
	}
	p.Block.Border = false

	return &p
}

// Draw draws the post to the buffer
func (p *Posts) Draw(buf *ui.Buffer) {
	row := p.Min.Y
	maxNumLen := len(fmt.Sprint(p.Offset + (p.Max.Y-p.Min.Y)/3))
	for i, post := range p.Posts[p.Offset:] {
		if row >= p.Max.Y {
			p.LastPost = p.Offset + i
			return
		}

		off := p.Min.X
		buf.SetString(image.Pt(off, row+1), fmt.Sprint(p.Offset+i), p.TitleStyle)

		off += maxNumLen + 1
		buf.SetString(image.Pt(off+2, row), string(symbols.UpVote), p.UpVoteStyle)
		buf.SetString(image.Pt(off+2, row+2), string(symbols.DownVote), p.DownVoteStyle)
		score := parseScore(post.Score)
		scoreOff := off
		if len(score) < 3 {
			scoreOff++
		}
		if len(score) < 2 {
			scoreOff++
		}
		buf.SetString(image.Pt(scoreOff, row+1), score, p.ScoreStyle)

		off += 6
		titleStyle := p.TitleStyle
		if p.Selected == p.Offset+i {
			titleStyle = p.SelectedStyle
		}

		titleRows := ui.ParseText(post.Title, p.Size().X-off)
		for i, titleRow := range titleRows {
			buf.SetString(image.Pt(off, row+i), titleRow, titleStyle)
		}
		submitted := parseTime(post.DateCreated)
		subtitle := fmt.Sprintf("%s by", submitted)
		buf.SetString(image.Pt(off, row+len(titleRows)), subtitle, p.SubtitleStyle)
		subtitleOff := off + len(subtitle) + 1
		buf.SetString(image.Pt(subtitleOff, row+len(titleRows)), post.Author, p.LinkStyle)
		if p.Frontpage {
			subtitleOff += len(post.Author) + 1
			buf.SetString(image.Pt(subtitleOff, row+len(titleRows)), "to", p.SubtitleStyle)
			subtitleOff += 3
			buf.SetString(image.Pt(subtitleOff, row+len(titleRows)), fmt.Sprintf("r/%s", post.Subreddit), p.LinkStyle)
		}

		comments := "1 comment"
		if post.NumComments != 1 {
			comments = fmt.Sprintf("%d comments", post.NumComments)
		}
		buf.SetString(image.Pt(off, row+len(titleRows)+1), comments, p.SubtitleStyle)

		row += 3 + len(titleRows)
	}
}

func parseScore(score int) string {
	if score < 10000 && score > -10000 {
		return fmt.Sprint(score)
	}

	if score > 100000 || score < -100000 {
		return fmt.Sprintf("%dk", score/1000)
	}

	return fmt.Sprintf("%.1fk", float32(score)/1000)
}

func parseTime(created float64) string {
	t := time.Unix(int64(created), int64(0))
	now := time.Now().UTC()
	y1, M1, d1 := t.Date()
	y2, M2, d2 := now.Date()

	h1, m1, s1 := t.Clock()
	h2, m2, s2 := now.Clock()

	year := int(y2 - y1)
	month := int(M2 - M1)
	day := int(d2 - d1)
	hour := int(h2 - h1)
	min := int(m2 - m1)
	sec := int(s2 - s1)

	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	if year > 1 {
		return fmt.Sprintf("submitted %d years ago", year)
	}
	if year == 1 {
		return "submitted 1 year ago"
	}
	if month > 1 {
		return fmt.Sprintf("submitted %d months ago", month)
	}
	if month == 1 {
		return "submitted 1 month ago"
	}
	if day > 1 {
		return fmt.Sprintf("submitted %d days ago", day)
	}
	if day == 1 {
		return "submitted 1 day ago"
	}
	if hour > 1 {
		return fmt.Sprintf("submitted %d hours ago", hour)
	}
	if hour == 1 {
		return "submitted 1 hour ago"
	}
	if min > 1 {
		return fmt.Sprintf("submitted %d minutes ago", min)
	}
	if min == 1 {
		return "submitted 1 minute ago"
	}

	return "submitted just now"
}
