package post

import (
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendSentProjReq ... appends to either sent project request arrays within user
//takes in eclient, user ID, the project ID
func AppendSentProjReq(eclient *elastic.Client, usrID string, projID string) error {
	ProjectLock.Lock()

	usr, err := get.UserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.SentProjReq = append(usr.SentProjReq, projID)

	err = UpdateUser(eclient, usrID, "SentProjReq", usr.SentProjReq)

	defer ProjectLock.Unlock()
	return err

}
