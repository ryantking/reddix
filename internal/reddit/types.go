package reddit

import "github.com/jzelinskie/geddit"

// Session holds information about the current user session
type Session struct {
	LoggedIn bool
	Username string
	Password string

	DefaultSess *geddit.Session
	LoginSess   *geddit.LoginSession
}
