package get

import (
	"context"
	"errors"
	"strings"

	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UserIDByEmail ...
func UserIDByEmail(eclient *elastic.Client, email string) (string, error) {

	ctx := context.Background()

	//username:= EmailToUsername(email)

	//termQuery := elastic.NewTermQuery("Username",username)

	termQuery := elastic.NewTermQuery("Email", strings.ToLower(email))
	searchResult, err := eclient.Search().
		Index(globals.UserIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	var result string //save id to a result variable
	if searchResult.TotalHits() > 1 {
		return result, errors.New("More than one user found")
	}

	if searchResult.TotalHits() < 1 {
		return result, errors.New("No hits, check your index or search criteria")
	}

	for _, element := range searchResult.Hits.Hits { //interate through hits, get the element id
		result = element.Id
		break
	}

	return result, err //return id

}
