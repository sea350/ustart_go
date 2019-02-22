package post

import (
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	elastic "github.com/olivere/elastic"
)

//AppendSentProjReq ... appends to either sent project request arrays within user
//takes in eclient, user ID, the project ID
func AppendSentProjReq(eclient *elastic.Client, usrID string, projID string) error {
	ProjectLock.Lock()
	defer ProjectLock.Unlock()

	usr, err := get.UserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.SentProjReq = append(usr.SentProjReq, projID)

	err = UpdateUser(eclient, usrID, "SentProjReq", usr.SentProjReq)

	return err

}
