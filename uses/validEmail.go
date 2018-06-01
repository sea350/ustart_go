package uses

import (
	"strings"
)

var validExtension = []string{`nyu.edu`}

//ValidEmail ... checks if the email entered is allowed on our service
func ValidEmail(email string) bool {
	splitEmail := strings.Split(email, `@`)
	if len(splitEmail) < 2 {
		return false
	}
	for _, ext := range validExtension {
		if splitEmail[len(splitEmail)-1] != ext {
			return false
		}
	}

	return true

}
