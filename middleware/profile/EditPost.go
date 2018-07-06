package profile

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	data, err := json.Marshal(editedEntry)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	fmt.Fprintln(w, string(data))

}
