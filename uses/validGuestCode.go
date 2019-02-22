package uses

import (
	"errors"
	"fmt"
	"time"

	getCode "github.com/sea350/ustart_go/get/guestCode"
	elastic "github.com/olivere/elastic"
)

//ValidGuestCode ... Check if guest code is valid
func ValidGuestCode(eclient *elastic.Client, guestCode string) (bool, error) {
	codeObj, err := getCode.GuestCodeByID(eclient, guestCode)
	if err != nil {
		return false, err
	}

	//Check if code expired (time and number of uses)

	if codeObj.Expiration.Before(time.Now()) && (codeObj.Classification == 2 || codeObj.Classification == 3) {
		return false, errors.New("Code Expired")
	}
	if codeObj.NumUses-len(codeObj.Users) < 0 && (codeObj.Classification == 1 || codeObj.Classification == 3) {
		fmt.Println(codeObj.NumUses - len(codeObj.Users))
		return false, errors.New("Access code is no longer valid")
	}

	return true, err

}
