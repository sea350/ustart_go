package entry

import (
	"fmt"
	"log"
	"net/http"

	get "github.com/sea350/ustart_go/backend/get/entry"
	client "github.com/sea350/ustart_go/backend/middleware/client"
	post "github.com/sea350/ustart_go/backend/post/entry"
)

//DeleteEntry ... Can delete any post NEEDS SECURITY CHECK
//designed for ajax
func DeleteEntry(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	entryID := r.FormValue("postid")

	entry, err := get.EntryByID(client.Eclient, entryID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	err = post.UpdateEntry(client.Eclient, entryID, "Visible", false)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	// err = delete.Entry(eclient, entryID)
	// if err != nil {
	// 	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 	log.Println(err)
	// 	return
	// }

	switch entry.Classification {
	case 0:
		for i := range entry.ShareIDs {
			err := post.UpdateEntry(client.Eclient, entry.ShareIDs[i], "ReferenceEntry", ``)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
		}
	case 1:
		err = post.DeleteReplyID(client.Eclient, entry.ReferenceEntry, entryID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return
		}
	case 2:
		err = post.DeleteShareID(client.Eclient, entry.ReferenceEntry, entryID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
	}

	fmt.Fprintln(w, entry.ReferenceEntry)
}
