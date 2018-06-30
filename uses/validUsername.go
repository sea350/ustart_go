package uses

import (
	"regexp"
)

//ValidUsername ... checks if the username entered is properly formatted
func ValidUsername(uname string) bool {
	//checking for proper email format
	rxEmail := regexp.MustCompile(`[a-zA-Z0-9][^-\s][a-zA-Z0-9_][^-\s]+[a-zA-Z0-9][^-\s]`) //double check if you need slashes
	//taken from new-reg-nil.htmll
	if !rxEmail.MatchString(uname) {
		return false
	}

	return true

}
