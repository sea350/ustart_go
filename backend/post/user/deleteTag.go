package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/backend/get/user"
	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteTag ... deletes a tag
func DeleteTag(eclient *elastic.Client, usrID string, tag string, idx int) error {
	ctx := context.Background()

	TagLock.Lock()

	usr, err := get.UserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	index := -1
	for i := range usr.Tags {
		if usr.Tags[i] == tag {
			index = i
		}
	}
	if index < 0 {
		return errors.New("index does not exist")
	}
	usr.Tags = append(usr.Tags[:idx], usr.Tags[idx+1:]...)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"Tags": usr.Tags}).
		Do(ctx)

	defer TagLock.Unlock()
	return err

}
