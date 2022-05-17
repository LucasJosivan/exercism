package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	character := map[string]string{"U+2757": "recommendation", "U+1F50D": "search", "U+2600": "weather"}
	for _, char := range log {
		for key, value := range character {
			if fmt.Sprintf("%U", char) == key {
				return value
			}
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, fmt.Sprintf("%c", oldRune), fmt.Sprintf("%c", newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
