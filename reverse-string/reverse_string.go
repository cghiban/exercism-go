package reverse

// Reverse reverses the given input string
func Reverse(input string) string {
	output := ""
	for _, r := range input {
		output = string(r) + output
	}
	return output
}
