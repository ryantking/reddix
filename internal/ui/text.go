package ui

import (
	"strings"

	"github.com/RyanTKing/reddix/internal/ui/symbols"
)

// ParseText wraps text to a given width and expands newlines
func ParseText(s string, width int) []string {
	s = strings.Replace(s, string(symbols.LeftSingleQuote), "'", -1)
	s = strings.Replace(s, string(symbols.RightSingleQuote), "'", -1)
	s = strings.Replace(s, string(symbols.LeftDoubleQuote), "\"", -1)
	s = strings.Replace(s, string(symbols.RightDoubleQuote), "\"", -1)

	wrapped := []string{}
	for s != "" {
		if len(s) < width {
			wrapped = append(wrapped, s)
			break
		}

		lastSpace := 0
		for i, c := range s {
			if i >= width {
				break
			}

			if c == '\n' {
				lastSpace = i
				break
			}

			if c == ' ' {
				lastSpace = i
			}
		}

		wrapped = append(wrapped, s[:lastSpace])
		s = s[lastSpace+1 : len(s)]
	}

	return wrapped
}
