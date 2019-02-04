package uses

import (
	"regexp"
	"strings"
)

//ValidUsername ... checks if the username entered is properly formatted
func ValidUsername(uname string) bool {
	//checking for proper email format
	// rxEmail := regexp.MustCompile(`[A-Za-z\d][\w\d-]+[A-Za-z\d]`) //double check if you need slashes
	//taken from new-reg-nil.html
	rxEmail := regexp.MustCompile(`^[a-zA-Z0-9]*$`)
	if !rxEmail.MatchString(uname) {
		return false
	}

	for i := range blacklistedUsernames {
		if strings.ToLower(uname) == blacklistedUsernames[i] {
			return false
		}
	}

	return true

}

var blacklistedUsernames = []string{
	"admin",
	"administrator",
	"support",
	"ustart",
	"ustartsupport",
	"ustart_support",
	"ustartadmin",
	"ustart_admin",
	"ustartadministrator",
	"ustart_administrator",
	"rozbiani",
	"ryanteyomrozbiani",
	"teyom",
	"teyomrozbiani",
	"ryan",
}
