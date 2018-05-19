package profile

import (
	"fmt"
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//ShareEntry ... Creates a new shared entry for user
func ShareEntry(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	originalPost := r.FormValue("postid")
	newContent := []rune(r.FormValue("content"))
	err := uses.UserShareEntry(client.Eclient, docID.(string), originalPost, newContent)
	if err != nil {
		fmt.Println("err: middleware/profile/shareentry line 25")
		fmt.Println(err)
	}

	fmt.Fprintln(w, "complete")
}
