package view

import (
	"fmt"
	"time"

	"github.com/RyanTKing/reddix/internal/browser"
	"github.com/RyanTKing/reddix/pkg/event"
	"github.com/RyanTKing/reddix/pkg/session"
	"github.com/RyanTKing/reddix/pkg/ui"
	"github.com/RyanTKing/reddix/pkg/ui/elements"
	"github.com/RyanTKing/reddix/pkg/window"
)

// NewBrowse creates a new browse buffer
func NewBrowse(x, y, width, height int) *Browse {
	return &Browse{x: x, y: y, width: width, height: height}
}

// Rect returns the rectangle of the view
func (browse *Browse) Rect() (int, int, int, int) {
	return browse.x, browse.y, browse.x + browse.width, browse.y + browse.height
}

// HandleKey responds to a keyboard input in a browse view
func (browse *Browse) HandleKey(sess *session.Session, win window.Window, key *event.Keyboard) error {
	switch key.Key {
	case "q":
		win.Quit()
	case "j":
		if browse.selected < len(browse.posts)-1 {
			browse.selected++
			if browse.selected >= browse.lastPost && browse.lastPost != len(browse.posts)-1 {
				browse.postOffset++
			}

			return browse.draw(sess, win)
		}
	case "k":
		if browse.selected > 0 {
			browse.selected--
			if browse.selected < browse.postOffset {
				browse.postOffset--
			}

			return browse.draw(sess, win)
		}
	case "o":
		err := browser.Open(browse.posts[browse.selected].FullPermalink())
		if err != nil {
			return err
		}
	}
	return nil
}

// HandleResize responds to a window resize
func (browse *Browse) HandleResize(win window.Window, width, height int) {
	browse.width = width
	browse.height = height
}

// RefreshPosts gets the posts from the currently selected subreddit and draws them to the screen
func (browse *Browse) RefreshPosts(sess *session.Session, win window.Window) error {
	posts, err := sess.GetPosts()
	if err != nil {
		return err
	}

	browse.posts = posts
	browse.postOffset = 0
	browse.selected = 0
	return browse.draw(sess, win)
}

func (browse *Browse) draw(sess *session.Session, win window.Window) error {
	clear(browse, win)
	row := 0
	titleLen := browse.width - 9
	for i, post := range browse.posts[browse.postOffset:] {
		title := ui.ParseText(post.Title, titleLen)
		if row >= browse.height {
			break
		}
		if row+len(title)+3 >= browse.height {
			browse.lastPost = browse.postOffset + i
		}

		height := 2 + len(title)
		if row+height > browse.height {
			height = browse.height - row
		}

		pn := elements.NewPostNumber(browse.postOffset + i + 1)
		pn.SetRect(browse.x, browse.y+row, browse.x+2, browse.y+row+height)
		win.Draw(pn)

		submitted := parseTime(post.DateCreated)
		score := parseScore(post.Score)
		p := elements.NewPost(title, post.Author, submitted, score, post.NumComments)
		if sess.Subreddit == "" {
			p.Subreddit = post.Subreddit
		}
		p.SetRect(browse.x+3, browse.y+row, browse.x+browse.width, browse.y+row+height)
		if browse.postOffset+i == browse.selected {
			p.TitleStyle = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold|ui.ModifierUnderline)
		}
		win.Draw(p)
		row += height + 1
	}

	return win.Refresh()
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
		return fmt.Sprintf("%d years", year)
	}
	if year == 1 {
		return "1 year"
	}
	if month > 1 {
		return fmt.Sprintf("%d months", month)
	}
	if month == 1 {
		return "1 month"
	}
	if day > 1 {
		return fmt.Sprintf("%d days", day)
	}
	if day == 1 {
		return "1 day"
	}
	if hour > 1 {
		return fmt.Sprintf("%d hours", hour)
	}
	if hour == 1 {
		return "1 hour"
	}
	if min > 1 {
		return fmt.Sprintf("%d minutes", min)
	}
	if min == 1 {
		return "1 minute"
	}

	return "just now"
}
