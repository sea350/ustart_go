package project

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sea350/ustart_go/delete"
	get "github.com/sea350/ustart_go/get/entry"
	getProj "github.com/sea350/ustart_go/get/project"
	client "github.com/sea350/ustart_go/middleware/client"
	postEntry "github.com/sea350/ustart_go/post/entry"
	postProj "github.com/sea350/ustart_go/post/project"
)

//AjaxDeleteEntry ... removes a project associated entry
func AjaxDeleteEntry(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	projID := r.FormValue("projID")
	entryID := r.FormValue("entryID")

	entry, err := get.EntryByID(client.Eclient, entryID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	err = delete.Entry(client.Eclient, entryID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	//removing refrence to entry in user
	proj, err := getProj.ProjectByID(client.Eclient, projID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	removeIdx := -1
	for idx := range proj.EntryIDs {
		if proj.EntryIDs[idx] == entryID {
			removeIdx = idx
		}
	}
	if removeIdx != -1 {
		var updatedEntries []string
		//update the user entries array
		if removeIdx+1 < len(proj.EntryIDs) {
			updatedEntries = append(proj.EntryIDs[:removeIdx], proj.EntryIDs[removeIdx+1:]...)
		} else {
			updatedEntries = proj.EntryIDs[:removeIdx]
		}

		err = postProj.UpdateProject(client.Eclient, projID, "EntryIDs", updatedEntries)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
	}

	//if reply, remove reference from parent
	if entry.Classification == 1 {
		parent, err := get.EntryByID(client.Eclient, entry.ReferenceEntry)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
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
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
		}
	}

	//if share, remove from reference entry
	if entry.Classification == 2 {
		parent, err := get.EntryByID(client.Eclient, entry.ReferenceEntry)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
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
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
		}
	}

	fmt.Fprintln(w, "REEEEEEEEE")
}
