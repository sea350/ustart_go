package uses

import (
	"log"

	delete "github.com/sea350/ustart_go/delete"
	get "github.com/sea350/ustart_go/get/entry"
	getUser "github.com/sea350/ustart_go/get/user"
	postEntry "github.com/sea350/ustart_go/post/entry"
	postUser "github.com/sea350/ustart_go/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//RemoveEntry ...
//Removes entry from ES
func RemoveEntry(eclient *elastic.Client, entryID string) (string, error) {

	entry, err := get.EntryByID(eclient, entryID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return ``, err
	}

	err = delete.Entry(eclient, entryID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return ``, err
	}

	//removing refrence to entry in user
	usr, err := getUser.UserByID(eclient, entry.PosterID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return ``, err
	}

	removeIdx := -1
	for idx := range usr.EntryIDs {
		if usr.EntryIDs[idx] == entryID {
			removeIdx = idx
		}
	}
	if removeIdx != -1 {
		var updatedEntries []string
		//update the user entries array
		if len(usr.EntryIDs) != 0 {
			if removeIdx+1 < len(usr.EntryIDs) {
				updatedEntries = append(usr.EntryIDs[:removeIdx], usr.EntryIDs[removeIdx+1:]...)
			} else {
				updatedEntries = usr.EntryIDs[:removeIdx]
			}
		}

		err = postUser.UpdateUser(eclient, entry.PosterID, "EntryIDs", updatedEntries)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return ``, err
		}
	}

	//if reply, remove reference from parent
	if entry.Classification == 1 {

		parent, err := get.EntryByID(eclient, entry.ReferenceEntry)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return ``, err
		}

		removeIdx := -1
		for idx := range parent.ReplyIDs {
			if parent.ReplyIDs[idx] == entryID {
				removeIdx = idx
			}
		}
		if removeIdx != -1 {
			var updatedReplies []string
			//remove from parent
			if removeIdx+1 < len(parent.ReplyIDs) {
				updatedReplies = append(parent.ReplyIDs[:removeIdx], parent.ReplyIDs[removeIdx+1:]...)
			} else {
				updatedReplies = parent.ReplyIDs[:removeIdx]
			}

			err = postEntry.UpdateEntry(eclient, entry.ReferenceEntry, "ReplyIDs", updatedReplies)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
				return ``, err
			}
		}
	}

	//if share, remove from reference entry
	if entry.Classification == 2 {
		parent, err := get.EntryByID(eclient, entry.ReferenceEntry)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return ``, err
		}
		removeIdx := -1
		for idx := range parent.ShareIDs {
			if parent.ShareIDs[idx] == entryID {
				removeIdx = idx
			}
		}
		if removeIdx != -1 {
			var updatedShares []string
			//remove from parent
			if removeIdx+1 < len(parent.ShareIDs) {
				updatedShares = append(parent.ShareIDs[:removeIdx], parent.ShareIDs[removeIdx+1:]...)
			} else {
				updatedShares = parent.ShareIDs[:removeIdx]
			}

			err = postEntry.UpdateEntry(eclient, entry.ReferenceEntry, "ShareIDs", updatedShares)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
				return ``, err
			}
		}
	}

	return entry.ReferenceEntry, err
}
