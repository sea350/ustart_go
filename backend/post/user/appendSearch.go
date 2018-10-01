package post

import (
	"context"

	get "github.com/sea350/ustart_go/backend/get/user"
	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendSearch ...
func AppendSearch(eclient *elastic.Client, usrID string, newSearch string) error {
	ctx := context.Background()
	usr, err := get.UserByID(eclient, usrID)
	if err != nil {
		return err
	}

	usr.SearchHist = append(usr.SearchHist, newSearch)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"SearchHist": usr.SearchHist}).
		Do(ctx)

	return err

}
