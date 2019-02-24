package window

import (
	"github.com/RyanTKing/reddix/internal/ui/elements"
	termbox "github.com/nsf/termbox-go"
)

func (win *Window) enterTextEntryMode(target string) {
	win.mode = TextEntry
	hidden := false
	if target == "password" {
		hidden = true
	}
	win.textEntryTarget = target
	win.TextEntry = elements.NewTextEntry(target, hidden)
	win.TextEntry.SetRect(0, win.Height-1, win.Width, win.Height)
	termbox.SetCursor(len(target)+2, win.Height-1)
	win.draw()
}

func (win *Window) handleTextEntryKey(ev termbox.Event) error {
	if ev.Key == termbox.KeyEsc {
		win.mode = Browse
		termbox.SetCursor(0, 0)
		win.draw()
		return nil
	}

	if ev.Ch != 0 {
		win.TextEntry.Text += string(ev.Ch)
		termbox.SetCursor(len(win.TextEntry.Prefix)+len(win.TextEntry.Text)+2, win.Height-1)
		win.draw()
		return nil
	}

	if ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
		if len(win.TextEntry.Text) > 0 {
			win.TextEntry.Text = win.TextEntry.Text[0 : len(win.TextEntry.Text)-1]
			termbox.SetCursor(len(win.TextEntry.Prefix)+len(win.TextEntry.Text)+2, win.Height-1)
			win.draw()
		}
		return nil
	}

	if ev.Key == termbox.KeyEnter {
		switch win.textEntryTarget {
		case "username":
			win.Sess.Username = win.TextEntry.Text
			win.enterTextEntryMode("password")
		case "password":
			win.Sess.Password = win.TextEntry.Text
			if err := win.Sess.Login(); err != nil {
				win.mode = Browse
				return err
			}
			win.TextEntry = elements.NewTextEntry("", false)
			win.TextEntry.Text = "login succeeded"
			win.mode = Browse
			win.draw()
		}
		return nil
	}

	return nil
}
