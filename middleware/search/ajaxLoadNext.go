package search

import (
	"net/http"

	"github.com/sea350/ustart_go/middleware/client"
)

func AjaxLoadNext(w http.ResponseWriter, r *http.Request, scrollID string) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
}
