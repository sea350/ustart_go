package post

import (
	get "github.com/sea350/ustart_go/get/user"
	elastic "github.com/olivere/elastic"
)

//AppendSentEventReq ... appends to either sent event request arrays within user
//takes in eclient, user ID, the event ID
func AppendSentEventReq(eclient *elastic.Client, usrID string, eventID string) error {

	EventLock.Lock()
	defer EventLock.Unlock()

	usr, err := get.UserByID(eclient, usrID)
	if err != nil {
		return err
	}

	usr.SentEventReq = append(usr.SentEventReq, eventID)

	err = UpdateUser(eclient, usrID, "SentEventReq", usr.SentEventReq)

	return err

}
