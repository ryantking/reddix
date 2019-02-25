package window

import (
	termbox "github.com/nsf/termbox-go"
)

func (win *Window) handleBrowseKey(ev termbox.Event) (bool, error) {
	if ev.Ch == 'q' {
		win.done = true
	}
	if ev.Ch == 'l' {
		if win.Sess.LoggedIn {
			err := win.Sess.Logout()
			win.TopMenu.Right = menuRight1
			win.BottomMenu.Right = "Not Logged In"
			return true, err
		}

		win.enterTextEntryMode("username")
		return true, nil
	}
	if ev.Ch == 'j' {
		if win.Posts.Selected < len(win.Posts.Posts)-1 {
			win.Posts.Selected++
			if win.Posts.Selected >= win.Posts.LastPost {
				win.Posts.Offset++
			}
			return true, nil
		}
		return false, nil
	}
	if ev.Ch == 'k' {
		if win.Posts.Selected > 0 {
			win.Posts.Selected--
			if win.Posts.Selected < win.Posts.Offset {
				win.Posts.Offset--
			}
			return true, nil
		}
		return false, nil
	}
	if ev.Ch == 'r' {
		win.enterTextEntryMode("subreddit")
		return true, nil
	}

	return false, nil
}
