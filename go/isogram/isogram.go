/*
	Package isogram provides the function IsIsogram(),
	to determine if a string is an isogram.

	Example:
		IsIsogram("background")
		-> true
*/
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns true if a character in a word is unique.
func IsIsogram(word string) bool {
	s := strings.ToLower(word)

	for i, r := range s {
		if unicode.IsLetter(r) && strings.ContainsRune(s[i+1:], r) {
			return false
		}
	}

	return true
}
