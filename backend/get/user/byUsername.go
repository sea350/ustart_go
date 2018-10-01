package get

import (
	"context"
	"strings"

	globals "github.com/sea350/ustart_go/backend/globals"
	types "github.com/sea350/ustart_go/backend/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UserByUsername ...
func UserByUsername(eclient *elastic.Client, username string) (types.User, error) {

	ctx := context.Background()

	//username:= EmailToUsername(email) //for username query
	termQuery := elastic.NewTermQuery("Username", strings.ToLower(username))
	searchResult, err := eclient.Search().
		Index(globals.UserIndex).
		Query(termQuery).
		Do(ctx)

	var usr types.User
	if err != nil {
		return usr, err
	}
	var result string

	for _, element := range searchResult.Hits.Hits {

		result = element.Id
		break
	}

	usr, err = UserByID(eclient, result)

	return usr, err

}
