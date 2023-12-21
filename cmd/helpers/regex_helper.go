package helpers

import (
	"regexp"
)

func IsBpddiyEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@bpddiy\.co\.id$`)
	return regex.MatchString(email)
}
