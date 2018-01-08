package post

import (
	"context"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendProjReq ... appends to either sent or received project request arrays within user
//takes in eclient, user ID, the project ID
func AppendReceivedProjReq(eclient *elastic.Client, usrID string, projID string) error {
	ctx := context.Background()

	projectLock.Lock()

	usr, err := get.UserByID(eclient, usrID)

	usr.ReceivedProjReq = append(usr.ReceivedProjReq, projID)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"ReceivedProjReq": usr.ReceivedProjReq}).
		Do(ctx)

	defer procLock.Unlock()
	return err
}
