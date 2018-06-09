package uses

import (
	"regexp"
)

//ValidUsername ... checks if the username entered is properly formatted
func ValidUsername(uname string) bool {
	//checking for proper email format
	rxEmail := regexp.MustCompile(`[a-zA-Z0-9][a-zA-Z0-9_]+[a-zA-Z0-9]`) //double check if you need slashes
	//taken from new-reg-nil.html
	if !rxEmail.MatchString(uname) {
		return false
	}

	return true

}