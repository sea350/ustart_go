package uses

import (
	"strings"
)

func StringChecker(entryCode string, website string) bool { //all 3 in one, checks iframe embed, div class embed, regular URL check
	ifValidWebsite := strings.Contains(entryCode, website)
	if !ifValidWebsite {
		return false
	}
	primaryChecker1 := "<iframe"
	primaryChecker2 := "<div class"

	ifValidPrefix1 := strings.HasPrefix(entryCode, primaryChecker1)
	ifValidPrefix2 := strings.HasPrefix(entryCode, primaryChecker2)

	if ifValidPrefix1 { //Iframe checker
		secondaryChecker1 := "</iframe>"
		ifValidSuffix1 := strings.HasSuffix(entryCode, secondaryChecker1)
		if !ifValidSuffix1 {
			return false
		}
	}
	if ifValidPrefix2 { //Divclass checker
		secondaryChecker2 := "</script>"
		ifValidSuffix2 := strings.HasSuffix(entryCode, secondaryChecker2)
		if !ifValidSuffix2 {
			return false
		}
	}

	if strings.HasPrefix(entryCode, "http") || strings.HasPrefix(entryCode, "https") { //Regular Chekcer
		if strings.Contains(entryCode, "www."+website) == false {
			return false
		}
	}

	return true

}
