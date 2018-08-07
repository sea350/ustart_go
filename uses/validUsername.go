package uses

import (
	"regexp"
)

//ValidUsername ... checks if the username entered is properly formatted
func ValidUsername(uname string) bool {
	//checking for proper email format
	rxEmail := regexp.MustCompile(`[A-Za-z\d][\w\d-]+[A-Za-z\d]`) //double check if you need slashes
	//taken from new-reg-nil.htmll
	if !rxEmail.MatchString(uname) {
		return false
	}

	return true

}
