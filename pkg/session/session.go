package session

import (
	"github.com/RyanTKing/reddix/internal/config"
	"github.com/jzelinskie/geddit"
	"github.com/zalando/go-keyring"
)

const (
	keyringService = "reddix"
	userAgent      = "gedditAgent v1"
)

// New returns a new reddit session
func New() *Session {
	sess := Session{
		DefaultSess: geddit.NewSession(userAgent),
	}

	return &sess
}

// Login attempts to log into a reddit session
func (s *Session) Login() (bool, error) {
	sess, err := geddit.NewLoginSession(s.Username, s.Password, userAgent)
	if err != nil {
		return false, err
	}

	s.LoginSess = sess
	err = keyring.Set(keyringService, s.Username, s.Password)
	if err != nil {
		return false, err
	}

	return true, nil
}

// UserLogin attempts to log into a reddit session with a given username
func (s *Session) UserLogin(username string) (bool, error) {
	s.Username = username
	password, err := keyring.Get(keyringService, username)
	if err != nil {
		return false, err
	}

	s.Password = password

	return s.Login()
}

// DefaultLogin attempts to log in to a reddit session using stored credentials
func (s *Session) DefaultLogin() (bool, error) {
	username, err := config.GetDefaultUser()
	if err != nil {
		return false, err
	}
	if username == "" {
		return false, nil
	}
	password, err := keyring.Get(keyringService, username)
	if err != nil {
		return false, err
	}

	s.Username = username
	s.Password = password
	return s.Login()
}

// Logout ends the current session
func (s *Session) Logout() error {
	if s.Username == "" {
		return nil
	}
	username := s.Username
	s.Username = ""
	s.Password = ""
	s.LoginSess = nil
	defaultUser, err := config.GetDefaultUser()
	if err != nil {
		return err
	}
	if username == defaultUser {
		err := config.DeleteDefaultUser()
		if err != nil {
			return err
		}
	}

	keyring.Delete(keyringService, username)
	return err
}

// LoggedIn returns whether or not the user is logged in
func (s *Session) LoggedIn() bool {
	return s.LoginSess != nil
}
