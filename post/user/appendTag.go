package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendTag ... appends a new tag
func AppendTag(eclient *elastic.Client, usrID string, tag string) error {
	ctx := context.Background()

	tagLock.Lock()
	defer tagLock.Unlock()
	usr, err := get.GetUserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.Tags = append(usr.Tags, tag)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"Tags": usr.Tags}).
		Do(ctx)

	return err

}
