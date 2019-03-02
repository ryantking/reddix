package session

import (
	"github.com/RyanTKing/reddix/internal/secrets"
	"github.com/jzelinskie/geddit"
)

const (
	userAgent = "gedditAgent v1"
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
	err = s.saveCreds()
	return true, err
}

// UserLogin attempts to log into a reddit session with a given username
func (s *Session) UserLogin(username string) (bool, error) {
	s.Username = username
	store := secrets.GetNativeStore()
	ok, password, err := store.Load(username)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	s.Password = password
	return s.Login()
}

// DefaultLogin attempts to log in to a reddit session using stored credentials
func (s *Session) DefaultLogin() (bool, error) {
	store := secrets.GetNativeStore()
	ok, username, password, err := store.LoadDefault()
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
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
	store := secrets.GetNativeStore()
	err := store.Delete(username)
	return err
}

// LoggedIn returns whether or not the user is logged in
func (s *Session) LoggedIn() bool {
	return s.LoginSess != nil
}

func (s *Session) saveCreds() error {
	store := secrets.GetNativeStore()

	err := store.Save(s.Username, s.Password)
	if err == secrets.ErrUserAlreadyExists {
		err = store.Delete(s.Username)
		if err != nil {
			return err
		}
		err = store.Save(s.Username, s.Password)
		if err != nil {
			return err
		}

		return nil
	}

	return err
}
