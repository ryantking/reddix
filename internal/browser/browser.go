package browser

import (
	"errors"
	"os/exec"
	"runtime"
)

var (
	// ErrUnknownPlatform is thrown when a platform cannot be determined
	ErrUnknownPlatform = errors.New("could not determine platform or platform is unsupported")
)

// Open opens a link in a browser
func Open(url string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		return ErrUnknownPlatform
	}
}
