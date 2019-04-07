package config

import (
	"io/ioutil"
	"os"
)

var defaultUserFile = os.ExpandEnv("$HOME/.config/reddix/default_user")

// GetDefaultUser returns the username of the default user if it exists
func GetDefaultUser() (string, error) {
	b, err := ioutil.ReadFile(defaultUserFile)
	if os.IsNotExist(err) {
		return "", nil
	}
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// SetDefaultUser stores the username of the default user
func SetDefaultUser(username string) error {
	f, err := os.OpenFile(defaultUserFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	_, err = f.WriteString(username)
	return err
}

// DeleteDefaultUser deletes the default user
func DeleteDefaultUser() error {
	return os.Remove(defaultUserFile)
}
