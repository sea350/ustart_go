package uses

import (
	"errors"
	"fmt"

	getCode "github.com/sea350/ustart_go/get/guestCode"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ValidGuestCode ... Check if guest code is valid
func ValidGuestCode(eclient *elastic.Client, guestCode string) (bool, error) {
	codeObj, err := getCode.GuestCodeByID(eclient, guestCode)
	if err != nil {
		return false, err
	}

	//Check if code expired (time and number of uses)
	/*
		if codeObj.Expiration.Before(time.Now()) {
			return false, errors.New("Code Expired")
		}*/
	if codeObj.NumUses-len(codeObj.Users) < 0 {
		fmt.Println(codeObj.NumUses - len(codeObj.Users))
		return false, errors.New("Exceeded number of uses")
	}

	return true, err

}
