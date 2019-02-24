package secrets

import "errors"

var (
	// ErrUserAlreadyExists is thrown when attempting to save a user that is already in the store
	ErrUserAlreadyExists = errors.New("the user already exists in the store")

	// ErrUserDoesNotExist is thrown when a user is attempted to be accessed that does not exist in the store
	ErrUserDoesNotExist = errors.New("no entry found for given username")
)
