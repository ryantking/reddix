// +build !darwin

package secrets

var store Store

// GetStore store returns the secret store
func GetStore() Store {
	if store != nil {
		return store
	}

	store = &DefaultStore{}

	return store
}

// Save is a dummy function for the default store
func (*DefaultStore) Save(username, password string) error {
	return nil
}

// Load is a dummy function for the default store
func (*DefaultStore) Load(username string) (bool, string, error) {
	return false, "", nil
}

// LoadDefault is a dummy function for the default store
func (*DefaultStore) LoadDefault() (bool, string, string, error) {
	return false, "", "", nil
}

// Delete is a dummy function for the default store
func (*DefaultStore) Delete(username string) error {
	return nil
}
