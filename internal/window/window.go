package window

import (
	"github.com/RyanTKing/reddix/internal/reddit"
	"github.com/RyanTKing/reddix/internal/ui"
	"github.com/RyanTKing/reddix/internal/ui/elements"
	termbox "github.com/nsf/termbox-go"
)

const (
	menuLeft   = "q:quit r:subreddit u:user h:help"
	menuCenter = "reddix"
	menuRight  = "l:login"
)

// New returns a new window
func New() (*Window, error) {
	if err := termbox.Init(); err != nil {
		return nil, err
	}
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCursor(0, 0)

	w, h := termbox.Size()
	sess, err := reddit.NewSession()
	topMenu := elements.NewMenu(menuLeft, menuCenter, menuRight)
	botMenu := elements.NewMenu("frontpage", "", "Not Logged In")
	if sess.Username != "" {
		botMenu.Left = sess.Username
	}
	win := Window{
		Width:      w,
		Height:     h,
		TopMenu:    topMenu,
		BottomMenu: botMenu,
		Sess:       sess,
	}
	if err != nil {
		win.Err = elements.NewError(err.Error())
	}

	return &win, nil
}

// Run starts the main event loop
func (win *Window) Run() {
	if win.Sess.Username != "" {
		win.enterTextEntryMode("password")
	}
	win.draw()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			var err error
			switch win.mode {
			case Browse:
				err = win.handleBrowseKey(ev)
			case TextEntry:
				err = win.handleTextEntryKey(ev)
			}

			if err != nil {
				win.Err = elements.NewError(err.Error())
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

func (win *Window) draw() {
	if win.TopMenu.Size().X != win.Width {
		win.TopMenu.SetRect(0, 0, win.Width, 1)
		win.drawItem(win.TopMenu)
	}
	if win.BottomMenu.Size().X != win.Width || win.BottomMenu.Max.Y != win.Height-1 {
		win.BottomMenu.SetRect(0, win.Height-2, win.Width, win.Height-1)
		win.drawItem(win.BottomMenu)
	}
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

	termbox.Flush()
}

func (win *Window) drawItem(item ui.Drawable) {
	buf := ui.NewBuffer(item.GetRect())
	item.Draw(buf)
	for p, cell := range buf.Cells {
		if p.In(buf.Rectangle) {
			fg := termbox.Attribute(cell.Style.FG + 1)
			bg := termbox.Attribute(cell.Style.BG + 1)
			termbox.SetCell(p.X, p.Y, cell.Rune, fg, bg)
		}
	}
}

// Close closes the current window
func (win *Window) Close() {
	termbox.Close()
}
