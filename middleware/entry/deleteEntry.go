package entry

import (
	"fmt"
	
	"net/http"

	get "github.com/sea350/ustart_go/get/entry"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/entry"
)

//DeleteEntry ... Can delete any post NEEDS SECURITY CHECK
//designed for ajax
func DeleteEntry(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	entryID := r.FormValue("postid")

	entry, err := get.EntryByID(client.Eclient, entryID)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		return
	}

	err = post.UpdateEntry(client.Eclient, entryID, "Visible", false)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		return
	}

	// err = delete.Entry(eclient, entryID)
	// if err != nil {
	// 	
	// 	client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	// 	return
	// }

	switch entry.Classification {
	case 0:
		for i := range entry.ShareIDs {
			err := post.UpdateEntry(client.Eclient, entry.ShareIDs[i], "ReferenceEntry", ``)
			if err != nil {
				
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			}
		}
	case 1:
		err = post.DeleteReplyID(client.Eclient, entry.ReferenceEntry, entryID)
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			return
		}
	case 2:
		err = post.DeleteShareID(client.Eclient, entry.ReferenceEntry, entryID)
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
	}

	fmt.Fprintln(w, entry.ReferenceEntry)
}
