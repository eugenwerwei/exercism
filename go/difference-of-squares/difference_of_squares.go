/*
	Package diffsquares provides a series of mathematical functions
	for calculations with natural numbers.
*/

package diffsquares

// SquareOfSum returns the square of the sum of N numbers as int.
func SquareOfSum(n int) int {
	return (n * (n + 1) / 2) * (n * (n + 1) / 2)
}

// SumOfSquares returns the sum of the squares of N numbers as int.
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference returns the difference between the square of the sum of N numbers
// and the sum of the squares of N numbers as int.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
