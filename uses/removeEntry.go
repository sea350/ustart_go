package uses

import (
	delete "github.com/sea350/ustart_go/delete"
	getUser "github.com/sea350/ustart_go/get/user"
	postUser "github.com/sea350/ustart_go/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//RemoveEntry ...
//Removes entry from ES
func RemoveEntry(eclient *elastic.Client, usrID, entryID string) error {
	errDelete := delete.Entry(eclient, entryID)

	if errDelete != nil {
		panic(errDelete)
	}
	usr, errGetUser := getUser.UserByID(eclient, usrID)
	//entry, errGet := getEntry.EntryByID(eclient, entryID)

	if errGetUser != nil {
		panic(errGetUser)
	}

	removeIdx := -1
	for idx := range usr.EntryIDs {
		if usr.EntryIDs[idx] == entryID {
			removeIdx = idx
		}
	}

	var updatedEntries []string
	//update the user entries array
	if removeIdx+1 < len(usr.EntryIDs) {
		updatedEntries = append(usr.EntryIDs[:removeIdx], usr.EntryIDs[removeIdx+1:]...)
	} else {
		updatedEntries = usr.EntryIDs[:removeIdx]
	}

	errUpdate := postUser.UpdateUser(eclient, usrID, "EntryIDs", updatedEntries)
	if errUpdate != nil {
		panic(errUpdate)
	}

	return errUpdate
}
