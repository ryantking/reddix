package session

import "github.com/jzelinskie/geddit"

// Session is a reddit session
type Session struct {
	Username  string
	Password  string
	Subreddit string

	DefaultSess *geddit.Session
	LoginSess   *geddit.LoginSession
}
