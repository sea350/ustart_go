package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendLink ... appends new link to QuickLinks
func AppendLink(eclient *elastic.Client, usrID string, link types.Link) error {
	ctx := context.Background()

	ProcLock.Lock()

	usr, err := get.UserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.QuickLinks = append(usr.QuickLinks, link)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"QuickLinks": usr.QuickLinks}).
		Do(ctx)

	defer ProcLock.Unlock()
	return err

}
