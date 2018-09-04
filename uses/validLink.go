package uses

import (
	"regexp"
)

//ValidLink ... checks if the email entered is allowed on our service
func ValidLink(link string) bool {
	//checking for proper email format
	rxLink := regexp.MustCompile(`^((http[s]?|ftp):\/)?\/?([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?`) //double check if you need slashes
	//taken from reg-check.js
	if !rxLink.MatchString(link) {
		return false
	}

	return true

}
