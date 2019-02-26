package window

import (
	termbox "github.com/nsf/termbox-go"
)

func (win *Window) handleBrowseKey(ev termbox.Event) {
	switch ev.Ch {
	case 'q':
		win.done = true
	case 'l':
		if win.Sess.LoggedIn {
			win.logout()
			break
		}

		win.enterInputMode("username")
	case 'j':
		if win.selected < len(win.posts)-1 {
			win.selected++
			if win.selected >= win.lastPost {
				win.postOffset++
			}

			win.drawPosts()
		}
	case 'k':
		if win.selected > 0 {
			win.selected--
			if win.selected < win.postOffset {
				win.postOffset--
			}

			win.drawPosts()
		}
	case 'r':
		win.enterInputMode("subreddit")
	}
}
