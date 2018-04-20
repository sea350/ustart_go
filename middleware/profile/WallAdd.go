package profile

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	docID := r.FormValue("docID")
	text := r.FormValue("text")
	textRunes := []rune(text)
	postID, err := uses.UserNewEntry(client.Eclient, docID, textRunes)
	if err != nil {
		fmt.Println(err)
	}
	postIDArray := []string{postID} // just an array with 1 entry
	jEntry, err5 := uses.LoadEntries(client.Eclient, postIDArray)
	if err5 != nil {
		fmt.Println(err5)
	}

	data, err := json.Marshal(jEntry)

	if err != nil {
		fmt.Println(err)
	}

	/*output := stringHTML.AddClass0Entry(jEntry[0].Image,
	jEntry[0].FirstName,
	string(jEntry[0].Element.Content),
	jEntry[0].ElementID,
	string(jEntry[0].NumLikes),
	string(jEntry[0].NumReplies),
	string(jEntry[0].NumShares))*/

	fmt.Fprintln(w, string(data))
}
