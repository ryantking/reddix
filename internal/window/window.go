package window

import (
	"github.com/RyanTKing/reddix/internal/reddit"
	"github.com/RyanTKing/reddix/internal/ui"

	termbox "github.com/nsf/termbox-go"
	"github.com/spf13/viper"
)

const (
	frontpage = "frontpage"
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
	username := viper.GetString("username")
	sess := reddit.NewSession(username)
	win := Window{
		Width:      w,
		Height:     h,
		TopMenu:    makeTopMenu(w),
		BottomMenu: makeBottomMenu(w, h),
		Sess:       sess,
		subreddit:  subreddit,
		miniBuffer: makeMiniBuffer(w, h),
	}

	if !viper.GetBool("anonymous") {
		win.login()
	}

	err := win.refreshPosts()
	if err != nil {
		win.setStatus(err.Error(), true)
		return &win, nil
	}

	return &win, nil
}

// Run starts the main event loop
func (win *Window) Run() {
	if win.Sess.Username != "" && !win.Sess.LoggedIn {
		win.enterInputMode("password")
	}

	win.drawTopMenu()
	win.drawBottomMenu()
	win.drawPosts()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch win.mode {
			case Browse:
				win.handleBrowseKey(ev)
			case TextEntry:
				win.handleTextEntryKey(ev)
			}
		case termbox.EventResize:
			win.resize(ev.Width, ev.Height)
		case termbox.EventError:
			win.setStatus(ev.Err.Error(), true)
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

func (win *Window) resize(width, height int) {
	if width != win.Width {
		win.TopMenu.SetRect(0, 0, width, 1)
		win.drawTopMenu()
	}

	if width != win.Width || height != win.Height {
		win.BottomMenu.SetRect(0, height-2, width, height-1)
		win.drawBottomMenu()
		win.drawPosts()
	}

	win.Width = width
	win.Height = height
}

func (win *Window) draw(item ui.Drawable) {
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
