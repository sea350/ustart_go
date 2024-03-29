package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//DeleteLink ... deletes QuickLink
func DeleteLink(eclient *elastic.Client, usrID string, link types.Link) error {
	ctx := context.Background()

	ProcLock.Lock()

	usr, err := get.UserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	index := -1
	for i := range usr.QuickLinks {
		if usr.QuickLinks[i] == link {
			index = i
		}
	}
	if index < 0 {
		return errors.New("index does not exist")
	}

	usr.QuickLinks = append(usr.QuickLinks[:index], usr.QuickLinks[index+1:]...)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"Quicklinks": usr.QuickLinks}).
		Do(ctx)

	defer ProcLock.Unlock()
	return err

}
