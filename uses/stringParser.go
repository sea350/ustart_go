package uses

import (
	"strings"
)

//StringChecker ...
func StringChecker(entryCode string, website string) bool {
	ifValidWebsite := strings.Contains(entryCode, website)
	if !ifValidWebsite {
		return false
	}
	primaryChecker1 := "<iframe"
	primaryChecker2 := "<div class"
	primaryChecker3 := "<p data-height"

	ifValidPrefix1 := strings.Contains(entryCode, primaryChecker1)
	ifValidPrefix2 := strings.Contains(entryCode, primaryChecker2)
	ifValidPrefix3 := strings.Contains(entryCode, primaryChecker3)

	if ifValidPrefix1 { //Iframe checker
		secondaryChecker1 := "</iframe>"
		ifValidSuffix1 := strings.Contains(entryCode, secondaryChecker1)
		if !ifValidSuffix1 {
			return false
		}
	}
	if ifValidPrefix2 { //Divclass checker
		secondaryChecker2 := "</script>"
		ifValidSuffix2 := strings.Contains(entryCode, secondaryChecker2)
		if !ifValidSuffix2 {
			return false
		}
	}

	if ifValidPrefix3 { //P data-height checker
		secondaryChecker3 := "</script>"
		ifValidSuffix3 := strings.Contains(entryCode, secondaryChecker3)
		if !ifValidSuffix3 {
			return false
		}
	}

	if strings.HasPrefix(entryCode, "http") || strings.HasPrefix(entryCode, "https") { //Regular Chekcer
		if strings.Contains(entryCode, "www."+website) == false {
			return false
		}
	}

	if website == "pinterest.com" {
		blockPoint := strings.Split(entryCode, "/")
		if !(len(blockPoint[4]) == 18 || len(blockPoint[4]) == 55 || len(blockPoint) == 87) {
			return false
		}
	}

	if website == "instagram.com" {
		blockPoint := strings.Split(entryCode, "/")
		if len(blockPoint[4]) != 10 {
			return false
		}
	}

	if website == "youtube.com" {
		blockPoint := strings.Split(entryCode, "/")
		if strings.Contains(blockPoint[3], "watch?v=") == false {
			return false
		}
	}

	return true

}
