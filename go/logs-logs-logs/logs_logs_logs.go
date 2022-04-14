package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, r := range log {
		if r == '❗' {
			return "recommendation"
		} else if r == '🔍' {
			return "search"
		} else if r == '☀' {
			return "weather"
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var str string

	for _, r := range log {
		if r == oldRune {
			str += string(newRune)
		} else {
			str += string(r)
		}
	}

	return str
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	numOfRunes := utf8.RuneCountInString(log)

	return numOfRunes <= limit
}
