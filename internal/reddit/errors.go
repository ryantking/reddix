package reddit

import "errors"

var (
	// ErrNotLoggedIn is thrown when a logout is called when a user is not logged in
	ErrNotLoggedIn = errors.New("no user is currently logged in")
)
