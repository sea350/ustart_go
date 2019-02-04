package uses

import (
	"regexp"
	"strings"

	"github.com/TwinProduction/go-away"
)

var blacklistedUsernames = []string{"admin", "administrator", "support", "ustart", "ustartsupport", "ustart_support", "ustartadmin", "ustart_admin", "ustartadministrator", "ustart_administrator"}

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

	return goaway.IsProfane(uname)

}
