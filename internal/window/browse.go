package window

import (
	termbox "github.com/nsf/termbox-go"
)

func (win *Window) handleBrowseKey(ev termbox.Event) error {
	if ev.Ch == 'q' {
		win.done = true
	}
	if ev.Ch == 'l' && !win.Sess.LoggedIn {
		win.enterTextEntryMode("username")
	}

	return nil
}
