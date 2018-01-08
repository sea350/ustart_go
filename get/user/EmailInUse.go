package get

import (
	"context"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//EmailInUse ...
func EmailInUse(eclient *elastic.Client, theEmail string) (bool, error) {
	ctx := context.Background()
	//username:=EmailToUsername(theEmail)
	//termQuery := elastic.NewTermQuery("Username",username)
	termQuery := elastic.NewTermQuery("Email", theEmail)
	searchResult, err := eclient.Search().
		Index(globals.UserIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return false, err
	} //email might not be in use, but it's still an error

	exists := searchResult.TotalHits() > 0

	return exists, err

}
