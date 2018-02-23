package profile

import (
	"fmt"
	"net/http"
)

//EditModal ... Iunno
func EditModal(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		//No docID in session
		http.Redirect(w, r, "/~", http.StatusFound)
	}

	fmt.Fprintln(w, "nothing")
}
