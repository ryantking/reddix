package window

import (
	"fmt"

	"github.com/RyanTKing/reddix/internal/reddit"
	"github.com/RyanTKing/reddix/internal/ui"
	"github.com/RyanTKing/reddix/internal/ui/elements"

	termbox "github.com/nsf/termbox-go"
	"github.com/spf13/viper"
)

const (
	menuLeft   = "q:quit r:subreddit u:user h:help"
	menuCenter = "reddix"
	menuRight1 = "l:login"
	menuRight2 = "l:logout"
)

// New returns a new window
func New(subreddit string) (*Window, error) {
	if err := termbox.Init(); err != nil {
		return nil, err
	}
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCursor(0, 0)

	w, h := termbox.Size()
	topMenu := elements.NewMenu(menuLeft, menuCenter, menuRight1)
	botMenu := elements.NewMenu("frontpage", "", "Not Logged In")
	username := viper.GetString("username")
	sess := reddit.NewSession(username)
	win := Window{
		Width:      w,
		Height:     h,
		TopMenu:    topMenu,
		BottomMenu: botMenu,
		Sess:       sess,
		subreddit:  subreddit,
	}

	if !viper.GetBool("anonymous") {
		loggedIn, err := win.Sess.Login()
		if err != nil {
			win.Err = elements.NewError(err.Error())
		}
		if loggedIn {
			win.TopMenu.Right = menuRight2
			win.BottomMenu.Right = win.Sess.Username
		}
	}

	err := win.refreshPosts()
	if err != nil {
		win.Err = elements.NewError(err.Error())
		return &win, nil
	}

	return &win, nil
}

// Run starts the main event loop
func (win *Window) Run() {
	if win.Sess.Username != "" && !win.Sess.LoggedIn {
		win.enterTextEntryMode("password")
	}
	win.draw()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			var redraw bool
			var err error
			switch win.mode {
			case Browse:
				redraw, err = win.handleBrowseKey(ev)
			case TextEntry:
				redraw, err = win.handleTextEntryKey(ev)
			}

			if err != nil {
				win.Err = elements.NewError(err.Error())
			}
			if redraw {
				win.draw()
			}
		case termbox.EventResize:
			win.Width = ev.Width
			win.Height = ev.Height
			win.draw()
		case termbox.EventError:
			win.Err = elements.NewError(ev.Err.Error())
		}

		if win.done {
			return
		}
	}
}

// Close closes the current window
func (win *Window) Close() {
	termbox.Close()
}

func (win *Window) draw() {
	if win.TopMenu.Size().X != win.Width {
		win.TopMenu.SetRect(0, 0, win.Width, 1)
	}
	win.drawItem(win.TopMenu)

	if win.BottomMenu.Size().X != win.Width || win.BottomMenu.Max.Y != win.Height-1 {
		win.BottomMenu.SetRect(0, win.Height-2, win.Width, win.Height-1)
	}
	win.drawItem(win.BottomMenu)

	if win.TextEntry != nil && (win.TextEntry.Size().X != win.Width || win.TextEntry.Max.Y != win.Height) {
		win.TextEntry.SetRect(0, win.Height-1, win.Width, win.Height)
	}
	if win.Err != nil {
		win.Err.SetRect(0, win.Height-1, win.Width, win.Height)
		win.drawItem(win.Err)
		win.Err = nil
	} else if win.TextEntry != nil {
		win.drawItem(win.TextEntry)
	}

	win.Posts.SetRect(0, 1, win.Width, win.Height-2)
	win.drawItem(win.Posts)

	termbox.Flush()
}

func (win *Window) drawItem(item ui.Drawable) {
	buf := ui.NewBuffer(item.GetRect())
	item.Draw(buf)
	for p, cell := range buf.Cells {
		if p.In(buf.Rectangle) {
			fg := termbox.Attribute(cell.Style.FG+1) | termbox.Attribute(cell.Style.Modifier)
			bg := termbox.Attribute(cell.Style.BG + 1)
			termbox.SetCell(p.X, p.Y, cell.Rune, fg, bg)
		}
	}
}

func (win *Window) refreshPosts() error {
	if win.subreddit == "" {
		posts, err := win.Sess.Frontpage()
		if err != nil {
			return err
		}

		win.Posts = elements.NewPosts(posts)
		win.Posts.Frontpage = true
		win.BottomMenu.Left = "frontpage"
		return nil
	}

	posts, err := win.Sess.Subreddit(win.subreddit)
	if err != nil {
		return err
	}

	win.Posts = elements.NewPosts(posts)
	win.Posts.Frontpage = false
	win.BottomMenu.Left = fmt.Sprintf("r/%s", win.subreddit)
	return nil
}
