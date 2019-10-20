package luhn

/*

BenchmarkValid-4   	 5000000	       275 ns/op

*/

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Valid checks if the input string is a Luhn number
func Valid(input string) bool {
	sum := 0
	k := 0
	input = reverse(input)
	for i := range input {
		r := input[i]
		if r >= 48 && r <= 57 { // we only care about digits
			num := int(r - '0') // convert to int
			if k%2 == 1 {       // position where we duplicate the number
				num = num + num
				if num > 9 {
					num -= 9
				}
			}
			sum += num
			k++
		} else if r != ' ' { // if not space (which is ignored)
			return false
		}
	}

	return k > 1 && sum%10 == 0

}
