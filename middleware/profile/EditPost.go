package profile

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	"github.com/sea350/ustart_go/uses"
)

//EditPost ... Iunno
func EditPost(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		//No docID in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	postID := r.FormValue("postid")
	newContent := r.FormValue("content")

	editedEntry, err := uses.EditEntry(client.Eclient, postID, "Content", []rune(newContent))
	if err != nil {
		log.Println("Error: middleware/post/EditPost line 26")
		log.Println(err)
	}

	data, err := json.Marshal(editedEntry)
	if err != nil {
		log.Println("Error: middleware/post/EditPost line 32")
		log.Println(err)
	}
	fmt.Fprintln(w, string(data))

}
