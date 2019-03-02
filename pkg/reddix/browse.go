package reddix

import (
	"github.com/RyanTKing/reddix/pkg/event"
)

func (reddix *Reddix) handleBrowseKey(ev *event.Keyboard) {
	switch ev.Key {
	case "l":
		if reddix.sess.LoggedIn() {
			reddix.Logout()
			return
		}

		reddix.EnterInputMode("username")
	case "r":
		reddix.EnterInputMode("subreddit")
	default:
		reddix.browse.HandleKey(reddix.sess, reddix.win, ev)
	}
}
