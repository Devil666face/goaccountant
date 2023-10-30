package utils

import (
	"github.com/microcosm-cc/bluemonday"
)

var policy = bluemonday.StrictPolicy()

func validate(in string) bool {
	if len([]rune(in)) == len([]rune(policy.Sanitize(in))) {
		return true
	}
	return false
}

func ValidateUserInputs(fields ...string) bool {
	for _, in := range fields {
		if !validate(in) {
			return false
		}
	}
	return true
}

func ValidateUserInput(field string) bool {
	return validate(field)
}
