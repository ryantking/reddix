package reddit

import (
	"fmt"

	"github.com/RyanTKing/reddix/internal/secrets"

	"github.com/jzelinskie/geddit"
	"github.com/spf13/viper"
)

const (
	userAgent = "gedditAgent v1"
)

// NewSession returns a new reddix session
func NewSession() (*Session, error) {
	sess := Session{}
	store := secrets.GetNativeStore()
	if store == nil {
		return &sess, nil
	}

	username := viper.GetString("username")
	sess.Username = username
	var password string
	var ok bool
	var err error
	if username == "" {
		ok, username, password, err = store.LoadDefault()
		sess.Username = username
	} else {
		ok, password, err = store.Load(sess.Username)
	}
	if err != nil {
		return nil, err
	}
	if !ok {
		return &sess, nil
	}

	fmt.Println(password)

	return &sess, nil
}

// Login attempts to log in to a reddit session
func (s *Session) Login() error {
	if s.Username == "" {
		s.DefaultSess = geddit.NewSession(userAgent)
		return nil
	}

	sess, err := geddit.NewLoginSession(s.Username, s.Password, "gedditAgent v1")
	if err != nil {
		return err
	}

	s.LoginSess = sess
	s.LoggedIn = true
	return nil
}
