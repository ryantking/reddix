package secrets

// Store represents anything that can store and retrievel credentials
type Store interface {
	Save(username, password string) error
	Load(username string) (bool, string, error)
	LoadDefault() (bool, string, string, error)
	Delete(username string) error
}

// DefaultStore is an empty store for when no service is available
type DefaultStore struct{}

// KeychainStore uses macOS keychain as a secret backend
type KeychainStore struct{}
