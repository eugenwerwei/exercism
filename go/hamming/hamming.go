/*
	Package hamming provides the function Distance(),
	for counting the differences of two DNA sequences.

	Example:
		Distance("GAGCCTACTAACGGGAT", "CGTCGTAATGACGGCCG")
		-> 9, nil
*/

package hamming

import "errors"

// Distance returns the counted differences of two strings as int.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Sequences are not of same length.")
	}

	count := 0

	for i, char := range a {
		if char != rune(b[i]) {
			count++
		}
	}

	return count, nil
}
