package uses

import (
	"strings"
)

//StringChecker ...
func StringChecker(entryCode string, website string) bool {
	//A lot of these changes are based on this document
	//https://docs.google.com/document/d/1H6AG11pxFkTyxSsXM8ZjSRRsTlTE4Xb5ZsCqzvNSU-g/edit

	//The Entry Code could be an embed code, a url, or anything else that's related to our website

	//Checks if the website is in our EntryCode
	ifValidWebsite := strings.Contains(entryCode, website)
	if !ifValidWebsite {
		return false
	}

	//3 following Embed checkers if our EntryCode is an embed code
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

	//If EntryCode is a url and has "http" or "https", check for the "www" keyword
	if strings.HasPrefix(entryCode, "http") || strings.HasPrefix(entryCode, "https") { //Regular Chekcer
		if strings.Contains(entryCode, "www."+website) == false {
			return false
		}
	}
	if website == "pinterest.com" {
		blockPoint := strings.Split(entryCode, "/")
		if !(len(blockPoint[3]) == 18 || len(blockPoint[3]) == 55 || len(blockPoint[3]) == 87) {
			return false
		}
	}

	if website == "instagram.com" {
		blockPoint := strings.Split(entryCode, "/")
		if len(blockPoint[4]) != 11 {
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
