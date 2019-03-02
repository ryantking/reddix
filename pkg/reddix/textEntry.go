package reddix

import (
	"github.com/RyanTKing/reddix/pkg/event"
)

// EnterInputMode enters into mini buffer input mode
func (reddix *Reddix) EnterInputMode(target string) {
	reddix.mode = TextEntry
	reddix.miniBuffer.BeginInput(reddix.win, target)
}

// ExitInputMode exits the mini buffer from input mode
func (reddix *Reddix) ExitInputMode() {
	switch reddix.miniBuffer.Target {
	case "username":
		reddix.sess.Username = reddix.miniBuffer.GetInput()
		reddix.EnterInputMode("password")
		reddix.miniBuffer.HiddenInput = true
		return
	case "password":
		reddix.sess.Password = reddix.miniBuffer.GetInput()
		reddix.Login()
	case "subreddit":
		subreddit := reddix.miniBuffer.GetInput()
		reddix.sess.Subreddit = subreddit
		reddix.bottomMenu.Left = subreddit
		if subreddit == "" {
			reddix.bottomMenu.Left = frontpage
		}
		reddix.bottomMenu.Draw(reddix.win)
		err := reddix.browse.RefreshPosts(reddix.sess, reddix.win)
		if err != nil {
			reddix.miniBuffer.SetStatus(reddix.win, err.Error(), true)
		}
	}

	reddix.mode = Browse
	reddix.win.SetCursor(0, 0)
	reddix.win.Refresh()
}

func (reddix *Reddix) handleTextEntryKey(ev *event.Keyboard) {
	switch ev.Key {
	case event.KeyEsc:
		reddix.mode = Browse
		reddix.win.SetCursor(0, 0)
		reddix.win.Refresh()
	case event.KeyBackspace:
		reddix.miniBuffer.Backspace(reddix.win)
	case event.KeyEnter:
		reddix.ExitInputMode()
	default:
		reddix.miniBuffer.Append(reddix.win, ev.Key)
	}
}
