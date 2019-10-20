package isbn

/*

BenchmarkIsValidISBN-4   	 5000000	       341 ns/op

*/

//IsValidISBN - validates string as being ISBN-10 or not
//(x1*10 + x2*9 + x3*8 + x4*7 + x5*6 + x6*5 + x7*4 + x8*3 + x9*2 + x10*1) mod 11 == 0
func IsValidISBN(isbn string) bool {

	sum := 0
	k := 10
	for i := range isbn {
		r := isbn[i]
		if r >= 48 && r <= 57 {
			/*num, err := strconv.Atoi(string(r))
			if err != nil {
				//fmt.Printf("Error %v for %s\n", err, string(r))
				return false
			}*/
			num := int(r - '0')
			sum += num * k
			k--
		} else if r == 'X' && k == 1 {
			sum += 10 * k
			k--
		}
	}
	if k != 0 {
		return false
	}

	return sum%11 == 0
}
