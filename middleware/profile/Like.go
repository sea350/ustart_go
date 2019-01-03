package profile

import (
	"encoding/json"
	"fmt"
	
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
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	entryID := r.FormValue("PostID")
	if entryID == `` {
		
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"WARNING: no entry Id passed in")
		return
	}

	likeStatus, err := uses.IsLiked(client.Eclient, entryID, docID.(string))
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}
	if likeStatus == true {
		err := uses.UserUnlikeEntry(client.Eclient, entryID, docID.(string))
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
		}
	} else {
		err := uses.UserLikeEntry(client.Eclient, entryID, docID.(string))
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
		}
	}

	entry, err := get.EntryByID(client.Eclient, entryID)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	data, err := json.Marshal(len(entry.Likes))
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}
	fmt.Fprintln(w, string(data))
}
