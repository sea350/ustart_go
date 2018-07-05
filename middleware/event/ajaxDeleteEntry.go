package event

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sea350/ustart_go/delete"
	get "github.com/sea350/ustart_go/get/entry"
	getEvnt "github.com/sea350/ustart_go/get/event"
	client "github.com/sea350/ustart_go/middleware/client"
	postEntry "github.com/sea350/ustart_go/post/entry"
	postEvnt "github.com/sea350/ustart_go/post/event"
)

//AjaxDeleteEventEntry ... removes an event associated entry
func AjaxDeleteEventEntry(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	eventID := r.FormValue("eventID") //This needs to be made on the front end side
	entryID := r.FormValue("entryID")

	entry, err := get.EntryByID(client.Eclient, entryID)
	if err != nil {
		log.Println("Error: middleware/event/ajaxDeleteEntry line 27")
		log.Println(err)
	}

	err = delete.Entry(client.Eclient, entryID)
	if err != nil {
		log.Println("Error: middleware/event/ajaxDeleteEntry line 33")
		log.Println(err)
	}

	//removing reference to entry in user
	evnt, err := getEvnt.EventByID(client.Eclient, eventID)
	if err != nil {
		log.Println("Error: middleware/event/ajaxDeleteEntry line 40")
		log.Println(err)
	}

	removeIdx := -1
	for idx := range evnt.EntryIDs {
		if evnt.EntryIDs[idx] == entryID {
			removeIdx = idx
		}
	}
	if removeIdx != -1 {
		var updatedEntries []string
		//update the user entires array
		if removeIdx+1 < len(evnt.EntryIDs) {
			updatedEntries = append(evnt.EntryIDs[:removeIdx], evnt.EntryIDs[removeIdx+1:]...)
		} else {
			updatedEntries = evnt.EntryIDs[:removeIdx]
		}

		err = postEvnt.UpdateEvent(client.Eclient, eventID, "EntryIDs", updatedEntries)
		if err != nil {
			log.Println("Error: middleware/event/ajaxDeleteEntry line 63")
			log.Println(err)
		}
	}
	//if reply, remove reference from parent
	if entry.Classification == 1 {

		parent, err := get.EntryByID(client.Eclient, entry.ReferenceEntry)
		if err != nil {
			log.Println("err: middleware/event/ajaxDeleteEntry line 73")
			log.Println(err)
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

			err = postEntry.UpdateEntry(client.Eclient, entry.ReferenceEntry, "ReplyIDs", updatedReplies)
			if err != nil {
				log.Println("Error: middleware/event/ajaxDeleteEntry line 92")
				log.Println(err)
			}
		}
	}

	//if share, remove from reference entry
	if entry.Classification == 2 {
		parent, err := get.EntryByID(client.Eclient, entry.ReferenceEntry)
		if err != nil {
			log.Println("Error: middleware/event/ajaxDeleteEntry line 102")
			log.Println(err)
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

			err = postEntry.UpdateEntry(client.Eclient, entry.ReferenceEntry, "ShareIDs", updatedShares)
			if err != nil {
				log.Println("Error: middleware/project/ajaxDeleteEntry line 122")
				log.Println(err)
			}
		}
	}

	fmt.Fprintln(w, "We done!")
}
