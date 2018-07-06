package profile

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//WallAdd ... Iunno
func WallAdd(w http.ResponseWriter, r *http.Request) {
	// If followingStatus = no
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	r.ParseForm()
	// docID := r.FormValue("docID")
	text := r.FormValue("text")
	textRunes := []rune(text)
	postID, err := uses.UserNewEntry(client.Eclient, session.Values["DocID"].(string), textRunes)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	postIDArray := []string{postID} // just an array with 1 entry
	jEntry, err := uses.LoadEntries(client.Eclient, postIDArray)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	data, err := json.Marshal(jEntry)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	fmt.Fprintln(w, string(data))
}
