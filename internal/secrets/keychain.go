package secrets

import (
	keychain "github.com/keybase/go-keychain"
)

const (
	keychainService = "reddix"
)

// Save saves a username and password tuple to the macOS keychain service
func (*KeychainStore) Save(username, password string) error {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(keychainService)
	item.SetAccount(username)
	item.SetData([]byte(password))
	err := keychain.AddItem(item)
	if err == keychain.ErrorDuplicateItem {
		return ErrUserAlreadyExists
	}
	if err != nil {
		return err
	}

	return nil
}

// Load loads a password for a given username from the macOS keychain service
func (*KeychainStore) Load(username string) (bool, string, error) {
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassGenericPassword)
	query.SetService(keychainService)
	query.SetAccount(username)
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnData(true)
	results, err := keychain.QueryItem(query)
	if err != nil {
		return false, "", err
	}
	if len(results) == 0 {
		return false, "", nil
	}

	return true, string(results[0].Data), nil
}

// LoadDefault loads the default username and password from the macOS keychain service
func (*KeychainStore) LoadDefault() (bool, string, string, error) {
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassGenericPassword)
	query.SetService(keychainService)
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnData(true)
	results, err := keychain.QueryItem(query)
	if err != nil {
		return false, "", "", err
	}
	if len(results) == 0 {
		return false, "", "", nil
	}

	return true, string(results[0].Account), string(results[0].Data), nil
}

// Delete removes a stored password from the macOS keychain service
func (*KeychainStore) Delete(username string) error {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(keychainService)
	item.SetAccount(username)
	err := keychain.DeleteItem(item)
	if err == keychain.ErrorItemNotFound {
		return ErrUserDoesNotExist
	}
	return err
}
