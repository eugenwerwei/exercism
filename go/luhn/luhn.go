/*
	Package luhn provides a function Valid(),
	to validate numbers based on the Luhn Algorithm.

	Example:
		Valid("8569 6195 0383 3437")
		-> true
*/

package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid returns true if a string passes the conditions of Luhn Algorithm.
func Valid(id string) bool {
	id = strings.TrimSpace(id)

	if len(id) < 2 || id == "0" {
		return false
	}

	sum := 0
	isSecond := false

	for i := len(id) - 1; i >= 0; i-- {
		if unicode.IsLetter(rune(id[i])) {
			return false
		}

		if unicode.IsSpace(rune(id[i])) {
			continue
		}

		digit, _ := strconv.Atoi(string(id[i]))

		if isSecond {
			if digit*2 > 9 {
				sum += digit*2 - 9
			} else {
				sum += digit * 2
			}
		} else {
			sum += digit
		}

		isSecond = !isSecond
	}

	return sum%10 == 0
}
