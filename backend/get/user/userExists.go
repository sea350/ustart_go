package get

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UserExists ... checks if user exists
func UserExists(eclient *elastic.Client, usrID string) (bool, error) {
	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.UserIndex).Do(ctx)
	if !exists {
		return false, errors.New("index does not exist")
	}
	if err != nil {
		return false, err
	}

	searchResult, err := eclient.Get().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Do(ctx)

	if err != nil {
		return false, err
	} //email might not be in use, but it's still an error

	return searchResult.Found, err

}
