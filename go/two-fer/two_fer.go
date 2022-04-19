/*
	Package twofer provides a function ShareWith() which prints a message.

	Example:
		ShareWith("Eugen")
		-> "One for Eugen, one for me."
		ShareWith("")
		-> "One for you, one for me."
*/
package twofer

import "fmt"

// ShareWith returns a formatted string based on the passed argument: "name".
func ShareWith(name string) string {
	if name != "" {
		return fmt.Sprintf("One for %s, one for me.", name)
	}

	return fmt.Sprint("One for you, one for me.")
}
