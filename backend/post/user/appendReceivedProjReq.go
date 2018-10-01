package post

import (
	get "github.com/sea350/ustart_go/backend/get/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendReceivedProjReq ... appends to either sent or received project request arrays within user
//takes in eclient, user ID, the project ID
func AppendReceivedProjReq(eclient *elastic.Client, usrID string, projID string) error {

	ProjectLock.Lock()

	usr, err := get.UserByID(eclient, usrID)

	usr.ReceivedProjReq = append(usr.ReceivedProjReq, projID)

	err = UpdateUser(eclient, usrID, "ReceivedProjReq", usr.ReceivedProjReq)

	defer ProjectLock.Unlock()
	return err
}
