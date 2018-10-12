package post

import (
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendSentEventReq ... appends to either sent event request arrays within user
//takes in eclient, user ID, the event ID
func AppendSentEventReq(eclient *elastic.Client, usrID string, eventID string) error {

	EventLock.Lock()
	defer EventLock.Unlock()

	usr, err := get.UserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.SentEventReq = append(usr.SentEventReq, eventID)

	err = UpdateUser(eclient, usrID, "SentEventReq", usr.SentEventReq)

	return err

}
