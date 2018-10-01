package post

import (
	get "github.com/sea350/ustart_go/backend/get/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendReceivedEventReq ... appends to either sent or received event request arrays within user
//takes in eclient, user ID, the event ID
func AppendReceivedEventReq(eclient *elastic.Client, usrID string, eventID string) error {

	EventLock.Lock()

	usr, err := get.UserByID(eclient, usrID)

	usr.ReceivedEventReq = append(usr.ReceivedEventReq, eventID)

	err = UpdateUser(eclient, usrID, "ReceivedEventReq", usr.ReceivedEventReq)

	defer EventLock.Unlock()
	return err
}
