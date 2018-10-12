package uses

import (
	"regexp"
	"strings"
)

var validExtension = []string{`nyu.edu`}

//ValidEmail ... checks if the email entered is allowed on our service
func ValidEmail(email string) bool {
	//checking for proper email format
	rxEmail := regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`) //double check if you need slashes
	//taken from reg-check.js
	if !rxEmail.MatchString(email) {
		return false
	}
	//checking for proper extensions
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

//ValidGuestEmail ... checks if the guest email entered is allowed on our service
func ValidGuestEmail(email string) bool {
	//checking for proper email format
	rxEmail := regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`) //double check if you need slashes
	//taken from reg-check.js
	if !rxEmail.MatchString(email) {
		return false
	}
	//checking for proper extensions
	splitEmail := strings.Split(email, `@`)
	if len(splitEmail) < 2 {
		return false
	}
	return true
}
