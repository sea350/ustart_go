package uses

import (
	"context"
	"errors"
	"time"

	getCode "github.com/sea350/ustart_go/get/guestCode"
	"github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ValidGuestCode ... Check if guest code is valid
func ValidGuestCode(eclient *elastic.Client, guestCode string) (bool, error) {
	ctx := context.Background()

	termQuery := elastic.NewTermQuery("Code", guestCode)
	var codeID string

	searchResult, err := eclient.Search().
		Index(globals.GuestCodeIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return false, err
	}

	if searchResult.TotalHits() != 1 {
		return false, err
	}

	//searching for then getting the code object, not sure if this is the best way to do it
	for _, element := range searchResult.Hits.Hits {
		codeID = element.Id
	}
	codeObj, err := getCode.GuestCodeByID(eclient, codeID)

	//Check if code expired (time and number of uses)
	if codeObj.Expiration.Before(time.Now()) {
		return false, errors.New("Code Expired")
	}
	if codeObj.NumUses-len(codeObj.Users) < 1 {
		return false, errors.New("Exceeded number of uses")
	}

	return true, err

}
