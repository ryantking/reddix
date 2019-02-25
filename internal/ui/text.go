package ui

import "unicode"

// ParseText wraps text to a given width and expands newlines
func ParseText(s string, width int) []string {
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
				if !unicode.IsPunct(rune(s[i+1])) {
					lastSpace = i
				}
			}
		}

		wrapped = append(wrapped, s[:lastSpace])
		s = s[lastSpace+1 : len(s)]
	}

	return wrapped
}
