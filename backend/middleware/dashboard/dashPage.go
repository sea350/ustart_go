package dashboard

import (
	"net/http"

	"github.com/sea350/ustart_go/backend/middleware/client"
)

//Page ... draws dashboard page
func Page(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	//cs := client.ClientSide{}
}
