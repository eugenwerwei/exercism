/*
  Package scrabble provides a function Score() which returns a number,
  based on the rules of the game scrabble.

  Example:
    Score("MILK")
    -> 10
*/

package scrabble

import "strings"

// Score returns a calculated number as an int,
// based on a passed string argument.
func Score(word string) int {
	if len(word) == 0 {
		return 0
	}

	score := 0

	for _, r := range word {
		switch strings.ToUpper(string(r)) {
		case "D", "G":
			score += 2
		case "B", "C", "M", "P":
			score += 3
		case "F", "H", "V", "W", "Y":
			score += 4
		case "K":
			score += 5
		case "J", "X":
			score += 8
		case "Q", "Z":
			score += 10
		default:
			score++
		}
	}

	return score
}
