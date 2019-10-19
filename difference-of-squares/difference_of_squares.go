package diffsquares

// SquareOfSum computes square of sum of 1st n integers [1, n]
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares computes sum of top squares
func SumOfSquares(n int) int {
	sum := n * (n + 1) * (n + n + 1) / 6
	return sum
}

// Difference computes SquareOfSum(n) - SumOfSquares(n)
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
