/*
	Package isogram provides the function IsIsogram(),
	to determine if a string is an isogram.

	Example:
		IsIsogram("background")
		-> true
*/
package isogram

import "strings"

// IsIsogram returns true if a character in a word is unique.
func IsIsogram(word string) bool {
	if word == "" {
		return true
	}

	str := strings.ToLower(word)

	for i := 0; i < len(str)-1; i++ {
		if string(str[i]) == " " || string(str[i]) == "-" {
			continue
		}

		for j := i + 1; j < len(str); j++ {
			if string(str[j]) == " " || string(str[j]) == "-" {
				continue
			}

			if string(str[i]) == string(str[j]) {
				return false
			}
		}
	}

	return true
}
