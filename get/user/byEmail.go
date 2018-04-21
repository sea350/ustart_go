package get

import (
	"context"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UserByEmail ...

func UserByEmail(eclient *elastic.Client, email string) (types.User, error) {

	ctx := context.Background()

	//username:= EmailToUsername(email) //for username query
	termQuery := elastic.NewTermQuery("Email", strings.ToLower(email))
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
