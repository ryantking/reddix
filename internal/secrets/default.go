package secrets

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
