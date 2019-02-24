package secrets

import "runtime"

var store Store

// GetNativeStore returns the native store for the current system
func GetNativeStore() Store {
	if store != nil {
		return store
	}

	switch runtime.GOOS {
	case "darwin":
		store = &KeychainStore{}
	}

	return store
}
