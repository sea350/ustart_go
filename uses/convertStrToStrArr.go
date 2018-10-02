package uses

import (
	"encoding/json"
)

//ConvertStrToStrArr ... given a json format array as a string type, returns string array
func ConvertStrToStrArr(unbrokenString string) []string {
	var stringArray []string
	_ = json.Unmarshal([]byte(unbrokenString), &stringArray)
	return stringArray
}
