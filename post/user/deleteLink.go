package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteLink ... deletes QuickLink
func DeleteLink(eclient *elastic.Client, usrID string, link types.Link) error {
	ctx := context.Background()

	procLock.Lock()
	defer procLock.Unlock()
	usr, err := get.GetUserByID(eclient, usrID)
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

	return err

}
