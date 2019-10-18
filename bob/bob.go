package bob

import (
	"regexp"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	//fmt.Println(remark)

	if matched, err := regexp.MatchString(`^\s*$`, remark); matched && err == nil {
		return "Fine. Be that way!"
	}

	if matched, err := regexp.MatchString(`(^[A-Z ]+!?$)|(^[^a-z]+!$)`, remark); matched && err == nil {
		return "Whoa, chill out!"
	}

	if matched, err := regexp.MatchString(`^[A-Z ]+\?$`, remark); matched && err == nil {
		return "Calm down, I know what I'm doing!"
	}

	if matched, err := regexp.MatchString(`\?\s*$`, remark); matched && err == nil {
		return "Sure."
	}
	return "Whatever."
}
