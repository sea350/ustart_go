package get

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UserIDByEmail ...
func UserIDByEmail(eclient *elastic.Client, email string) (string, error) {

	ctx := context.Background()

	//username:= EmailToUsername(email)

	//termQuery := elastic.NewTermQuery("Username",username)

	termQuery := elastic.NewTermQuery("Email", email)
	searchResult, err := eclient.Search().
		Index(globals.EntryIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	var result string //save id to a result variable
	if searchResult.TotalHits() > 1 {
		return result, errors.New("More than one user found")
	}

	for _, element := range searchResult.Hits.Hits { //interate through hits, get the element id
		result = element.Id

	}

	return result, err //return id

}
