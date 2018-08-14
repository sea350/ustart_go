package profile

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	get "github.com/sea350/ustart_go/get/entry"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//Like ... likes a post, designed for ajax
func Like(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	entryID := r.FormValue("PostID")
	if entryID == `` {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("WARNING: no entry Id passed in")
		return
	}

	likeStatus, err := uses.IsLiked(client.Eclient, entryID, docID.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	if likeStatus == true {
		err := uses.UserUnlikeEntry(client.Eclient, entryID, docID.(string))
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
	} else {
		err := uses.UserLikeEntry(client.Eclient, entryID, docID.(string))
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
	}

	entry, err := get.EntryByID(client.Eclient, entryID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	data, err := json.Marshal(len(entry.Likes))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	fmt.Fprintln(w, string(data))
}
