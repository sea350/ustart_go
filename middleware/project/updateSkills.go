package project

import (
	"encoding/json"
	
	"net/http"
	

	"github.com/microcosm-cc/bluemonday"

	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/project"
)

//UpdateSkills ...
func UpdateSkills(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	ID := r.FormValue("projectWidget")
	skills := r.FormValue("skillArray")

	p := bluemonday.UGCPolicy()

	var ss []string
	err := json.Unmarshal([]byte(skills), &ss)
	if err != nil {
		

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	for s := range ss {
		ss[s] = p.Sanitize(ss[s])
	}
	err = post.UpdateProject(client.Eclient, ID, "ListNeeded", ss)
	if err != nil {
		

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}
}
