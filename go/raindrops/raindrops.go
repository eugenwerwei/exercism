/*
	Package raindrops provides the function Convert(),
	which returns a formatted string based on an argument.

	Example:
		Convert(28)
		-> "Plong"
*/
package raindrops

import "fmt"

// Convert takes an int as an argument and returns a message as a string,
// based on a set of remainder checks of the argument.
func Convert(number int) string {
	result := ""

	if number%3 == 0 {
		result += "Pling"
	}

	if number%5 == 0 {
		result += "Plang"
	}

	if number%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		return fmt.Sprintf("%d", number)
	} else {
		return result
	}
}
