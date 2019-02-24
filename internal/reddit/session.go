package reddit

import (
	"github.com/RyanTKing/reddix/internal/secrets"

	"github.com/jzelinskie/geddit"
)

const (
	userAgent = "gedditAgent v1"
)

// NewSession returns a new reddix session
func NewSession(username string) *Session {
	sess := Session{
		Username:    username,
		DefaultSess: geddit.NewSession(userAgent),
	}

	return &sess
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

func (s *Session) loadCreds() (bool, error) {
	store := secrets.GetNativeStore()

	if s.Username == "" {
		ok, username, password, err := store.LoadDefault()
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}

		s.Username = username
		s.Password = password
		return true, nil
	}

	ok, password, err := store.Load(s.Username)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}
	s.Password = password

	return true, nil
}

// Login attempts to log in to a reddit session
func (s *Session) Login() (bool, error) {
	if s.Username == "" || s.Password == "" {
		ok, err := s.loadCreds()
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
	}

	sess, err := geddit.NewLoginSession(s.Username, s.Password, "gedditAgent v1")
	if err != nil {
		return false, err
	}

	s.LoginSess = sess
	s.LoggedIn = true
	err = s.saveCreds()
	return true, nil
}

// Logout ends the current session
func (s *Session) Logout() error {
	if s.Username == "" {
		return ErrNotLoggedIn
	}
	username := s.Username
	s.Username = ""
	s.Password = ""
	s.LoggedIn = false
	s.LoginSess = nil
	store := secrets.GetNativeStore()
	err := store.Delete(username)
	return err
}
