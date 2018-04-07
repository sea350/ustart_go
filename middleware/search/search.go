package search

import (
	"fmt"
	"net/http"

	"github.com/sea350/ustart_go/middleware/client"
)

//Page ... draws search page
func Page(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	cs := client.ClientSide{}

	query := r.FormValue("query")

	fmt.Fprintln(w, query)

	client.RenderTemplate(w, r, "template2-nil", cs)
	client.RenderTemplate(w, r, "leftnav-nil", cs)
	client.RenderTemplate(w, r, "search-nil", cs)
}
