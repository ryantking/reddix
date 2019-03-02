package reddix

import (
	"github.com/RyanTKing/reddix/pkg/event"
	"github.com/RyanTKing/reddix/pkg/session"
	"github.com/RyanTKing/reddix/pkg/view"
	"github.com/RyanTKing/reddix/pkg/window"
)

// New creates a new reddix instance to draw on the given window and attempts to log
// in to the default session
func New(win window.Window) *Reddix {
	reddix := newDefault(win)
	ok, err := reddix.sess.DefaultLogin()
	if err != nil {
		reddix.miniBuffer.SetStatus(win, err.Error(), true)
	}
	if ok {
		reddix.topMenu.Right = topMenuRight2
		reddix.bottomMenu.Right = reddix.sess.Username
	}

	return reddix
}

// NewWithUser creates a new reddix instance attempting to login as a user
func NewWithUser(win window.Window, username string) *Reddix {
	reddix := newDefault(win)
	ok, err := reddix.sess.UserLogin(username)
	if err != nil {
		reddix.miniBuffer.SetStatus(win, err.Error(), true)
	}
	if ok {
		reddix.topMenu.Right = topMenuRight2
		reddix.bottomMenu.Right = reddix.sess.Username
	}

	return reddix
}

// NewAnonymous creates a new anonymous reddix instance
func NewAnonymous(win window.Window) *Reddix {
	return newDefault(win)
}

// Run starts listening for events and handling them accordingly
func (reddix *Reddix) Run() {
	reddix.topMenu.Draw(reddix.win)
	reddix.bottomMenu.Draw(reddix.win)
	reddix.browse.RefreshPosts(reddix.sess, reddix.win)

	if reddix.sess.Username != "" && !reddix.sess.LoggedIn() {
		reddix.EnterInputMode("password")
		reddix.miniBuffer.HiddenInput = true
	}

	for {
		switch ev := reddix.win.Poll().(type) {
		case *event.Keyboard:
			reddix.handleKey(ev)
		case *event.Resize:
			reddix.topMenu.Resize(reddix.win, 0, ev.Width)
			reddix.bottomMenu.Resize(reddix.win, ev.Height-2, ev.Width)
			reddix.miniBuffer.Resize(reddix.win, ev.Height-1, ev.Width)
		case *event.Error:
		case *event.Quit:
			return
		}
	}
}

// Login attempts to login the curretn reddix instance to reddit
func (reddix *Reddix) Login() {
	loggedIn, err := reddix.sess.Login()
	if err != nil {
		reddix.miniBuffer.SetStatus(reddix.win, err.Error(), true)
		return
	}
	if !loggedIn {
		return
	}

	reddix.miniBuffer.SetStatus(reddix.win, "login succeeded", false)
	reddix.topMenu.Right = topMenuRight2
	reddix.topMenu.Draw(reddix.win)
	reddix.bottomMenu.Right = reddix.sess.Username
	reddix.bottomMenu.Draw(reddix.win)
	err = reddix.browse.RefreshPosts(reddix.sess, reddix.win)
	if err != nil {
		reddix.miniBuffer.SetStatus(reddix.win, err.Error(), true)
	}
}

// Logout ends the current session
func (reddix *Reddix) Logout() {
	err := reddix.sess.Logout()
	if err != nil {
		reddix.miniBuffer.SetStatus(reddix.win, err.Error(), true)
		return
	}

	reddix.miniBuffer.SetStatus(reddix.win, "logout succeeded", false)
	reddix.topMenu.Right = topMenuRight1
	reddix.topMenu.Draw(reddix.win)
	reddix.bottomMenu.Right = bottomMenuRight
	reddix.bottomMenu.Draw(reddix.win)
	err = reddix.browse.RefreshPosts(reddix.sess, reddix.win)
	if err != nil {
		reddix.miniBuffer.SetStatus(reddix.win, err.Error(), true)
	}
}

func newDefault(win window.Window) *Reddix {
	w, h := win.Size()
	sess := session.New()
	reddix := Reddix{
		mode:       Browse,
		win:        win,
		sess:       sess,
		topMenu:    NewMenu(0, w, topMenuLeft, topMenuCenter, topMenuRight1),
		bottomMenu: NewMenu(h-2, w, frontpage, "", bottomMenuRight),
		miniBuffer: NewMiniBuffer(h-1, w),
		browse:     view.NewBrowse(0, 1, w, h-3),
	}

	return &reddix
}

func (reddix *Reddix) handleKey(ev *event.Keyboard) {
	switch reddix.mode {
	case Browse:
		reddix.handleBrowseKey(ev)
	case TextEntry:
		reddix.handleTextEntryKey(ev)
	}
}
