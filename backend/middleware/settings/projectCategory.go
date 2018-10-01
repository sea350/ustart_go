package settings

import (
	"log"
	"net/http"
	"os"

	"github.com/microcosm-cc/bluemonday"

	get "github.com/sea350/ustart_go/backend/get/project"
	client "github.com/sea350/ustart_go/backend/middleware/client"
	uses "github.com/sea350/ustart_go/backend/uses"
)

//ProjectCategory ...
func ProjectCategory(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	p := bluemonday.UGCPolicy()
	newCategory := p.Sanitize(r.FormValue("type_select"))

	projID := r.FormValue("projectID")
	proj, err := get.ProjectByID(client.Eclient, projID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	err = uses.ChangeProjectCategory(client.Eclient, projID, newCategory)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
	return

}
