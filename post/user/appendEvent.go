package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendEvent ... appends new event to user
//takes in eclient, userID, event ID
func AppendEvent(eclient *elastic.Client, usrID string, evnt types.EventInfo) error {
	ctx := context.Background()

	EventLock.Lock()
	defer EventLock.Unlock()

	usr, err := get.UserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.Events = append(usr.Events, evnt)

	// log.Println("evnt: ", evnt)
	// log.Println("usr.Events: ", len(usr.Events))

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"Events": usr.Events}).
		Do(ctx)

	return err

}
