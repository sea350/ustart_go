package project

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sea350/ustart_go/uses"

	get "github.com/sea350/ustart_go/get/project"
	client "github.com/sea350/ustart_go/middleware/client"
)

//AjaxLoadProjectEntries ... pulls all entries for a given project and fprints it back as a json array
func AjaxLoadProjectEntries(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	wallID := r.FormValue("UNKNOWN")
	proj, err := get.ProjectByID(client.Eclient, wallID)
	if err != nil {
		fmt.Println("err middleware/project/AjaxLoadProjectEntries line 25")
		fmt.Println(err)
	}
	entries, err := uses.LoadEntries(client.Eclient, proj.EntryIDs)
	if err != nil {
		fmt.Println("err middleware/project/AjaxLoadProjectEntries line 30")
		fmt.Println(err)
	}

	data, err := json.Marshal(entries)
	if err != nil {
		fmt.Println("err middleware/project/AjaxLoadProjectEntries line 37")
		fmt.Println(err)
	}

	fmt.Fprintln(w, string(data))
}
