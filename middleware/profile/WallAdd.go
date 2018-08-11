package profile

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//WallAdd ... Iunno
func WallAdd(w http.ResponseWriter, r *http.Request) {
	// If followingStatus = no
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	r.ParseForm()
	// docID := r.FormValue("docID")
	text := html.EscapeString(r.FormValue("text"))
	textRunes := []rune(text)
	postID, err := uses.UserNewEntry(client.Eclient, docID.(string), textRunes)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	postIDArray := []string{postID} // just an array with 1 entry
	jEntry, err := uses.LoadEntries(client.Eclient, postIDArray, docID.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	data, err := json.Marshal(jEntry)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	fmt.Fprintln(w, string(data))
}
