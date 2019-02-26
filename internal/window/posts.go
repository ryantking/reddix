package window

import (
	"fmt"
	"time"

	"github.com/RyanTKing/reddix/internal/ui"

	"github.com/RyanTKing/reddix/internal/ui/elements"
)

func (win *Window) refreshPosts() error {
	if win.subreddit == "" {
		posts, err := win.Sess.Frontpage()
		if err != nil {
			return err
		}

		win.posts = posts
		win.BottomMenu.Left = "frontpage"
		return nil
	}

	posts, err := win.Sess.Subreddit(win.subreddit)
	if err != nil {
		return err
	}

	win.posts = posts
	win.BottomMenu.Left = fmt.Sprintf("r/%s", win.subreddit)
	return nil
}

func (win *Window) clearPosts() {
	block := ui.NewBlock()
	block.Border = false
	block.SetRect(0, 1, win.Width, win.Height-2)
	win.drawItem(block)
}

func (win *Window) drawPosts() {
	win.clearPosts()
	row := 1
	maxNumLen := len(fmt.Sprint(win.postOffset + (win.Height-3)/3))
	titleLen := win.Width - maxNumLen - 7
	for i, post := range win.posts[win.postOffset:] {
		title := ui.ParseText(post.Title, titleLen)
		if row >= win.Height-2 {
			win.lastPost = win.postOffset + i
			return
		}
		height := 2 + len(title)
		if row+height >= win.Height-2 {
			height = win.Height - row - 2
		}

		pn := elements.NewPostNumber(win.postOffset + i)
		pn.SetRect(0, row, maxNumLen, row+height)
		win.drawItem(pn)

		submitted := parseTime(post.DateCreated)
		score := parseScore(post.Score)
		p := elements.NewPost(title, post.Author, submitted, score, post.NumComments)
		if win.subreddit == "" {
			p.Subreddit = post.Subreddit
		}
		p.SetRect(maxNumLen+1, row, win.Width, row+height)
		if win.postOffset+i == win.selected {
			p.TitleStyle = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold|ui.ModifierUnderline)
		}
		win.drawItem(p)
		row += height + 1
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
