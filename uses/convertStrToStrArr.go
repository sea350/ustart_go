package uses

import "strings"

//ConvertStrToStrArr ... given a json format array as a string type, returns string array
func ConvertStrToStrArr(unbrokenString string) []string {
	stringArray := strings.Split(unbrokenString, `","`)
	if len(stringArray) > 0 {
		stringArray[0] = strings.Trim(stringArray[0], `["`)
		stringArray[len(stringArray)-1] = strings.Trim(stringArray[len(stringArray)-1], `"]`)
		if stringArray[0] == `` && len(stringArray) == 1 {
			var empty []string
			stringArray = empty
		}
	}
	return stringArray
}
