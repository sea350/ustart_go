package search

import (
	"fmt"
	"net/http"

	"github.com/sea350/ustart_go/middleware/client"
	"github.com/sea350/ustart_go/search"
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
	filter := r.FormValue("searchFilterGroup") //can be: all,users,projects

	if filter == `projects` {
		results, _, err := search.ProjectBasic(client.Eclient, query)
		if err != nil {
			fmt.Println("err: middleware/search/search line 26")
		}
		cs.ListOfHeads = results
	}
	if filter == `users` {
		results, err := search.UserBasic(client.Eclient, query)
		if err != nil {
			fmt.Println("err: middleware/search/search line 34")
		}
		cs.ListOfHeads = results
	}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "search-nil", cs)
}