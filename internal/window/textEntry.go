package window

import (
	termbox "github.com/nsf/termbox-go"
)

func (win *Window) handleTextEntryKey(ev termbox.Event) {
	if ev.Key == termbox.KeyEsc {
		win.exitInputMode()
		return
	}

	if ev.Ch != 0 {
		win.miniBufferAppendChar(ev.Ch)
		return
	}

	if ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
		win.miniBufferBackspace()
		return
	}

	if ev.Key == termbox.KeyEnter {
		switch win.miniBuffer.Target {
		case "username":
			win.Sess.Username = win.miniBuffer.Input
			win.enterInputMode("password")
			win.miniBuffer.HiddenInput = true
		case "password":
			win.Sess.Password = win.miniBuffer.Input
			win.exitInputMode()
			win.login()
		case "subreddit":
			win.exitInputMode()
			win.subreddit = win.miniBuffer.Input
			win.refreshPosts()
			win.drawPosts()
		}
	}
}
