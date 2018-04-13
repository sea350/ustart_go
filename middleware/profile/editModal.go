package profile

import (
	"fmt"
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
)

//EditModal ... Iunno
func EditModal(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		//No docID in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	fmt.Fprintln(w, "nothing")
}
