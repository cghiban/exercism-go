package accumulate

// Accumulate - given a collection and an operation to perform on
// each element of the collection, returns a new collection
// containing the result of applying that operation to each
// element of the input collection.
func Accumulate(input []string, op func(string) string) []string {

	output := []string{}

	for _, s := range input {
		output = append(output, op(s))
	}

	return output
}
