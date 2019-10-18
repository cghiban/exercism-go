package summultiples

//SumMultiples - solves the following problem:
//Given a number, find the sum of all the unique multiples of particular numbers up to but not including that number.
func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	seen := map[int]bool{}
	for i := 1; i < limit; i++ {
		for _, d := range divisors {
			if d != 0 && 0 == i%d {
				if !seen[i] {
					sum += i
					seen[i] = true
				}
			}
		}
	}
	return sum
}
