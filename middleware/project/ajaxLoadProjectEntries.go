package project

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	get "github.com/sea350/ustart_go/get/project"
	client "github.com/sea350/ustart_go/middleware/client"
	scrollpkg "github.com/sea350/ustart_go/properloading"
)

//AjaxLoadProjectEntries ... pulls all entries for a given project and fprints it back as a json array
func AjaxLoadProjectEntries(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	wallID := r.FormValue("userID")
	proj, err := get.ProjectByID(client.Eclient, wallID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	/*
		entries, err := uses.LoadEntries(client.Eclient, proj.EntryIDs)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
	*/

	res, entries, err := scrollpkg.ScrollPageProject(client.Eclient, proj.EntryIDs, "")
	if err != nil {
		fmt.Println(res)
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	data, err := json.Marshal(entries)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	fmt.Fprintln(w, string(data))
}
