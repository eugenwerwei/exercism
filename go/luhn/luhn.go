package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) bool {
	len := len(id)

	id = strings.TrimSpace(id)

	if len < 2 || id == "0" {
		return false
	}

	sum := 0
	isSecond := false

	for i := len - 1; i >= 0; i-- {
		if unicode.IsLetter(rune(id[i])) {
			return false
		}

		if unicode.IsSpace(rune(id[i])) {
			continue
		}

		digit, _ := strconv.Atoi(string(id[i]))

		if isSecond {
			digit = digit * 2

			if digit > 9 {
				sum += digit - 9
			} else {
				sum += digit
			}
		} else {

			sum += digit
		}

		isSecond = !isSecond
	}

	return sum%10 == 0
}
