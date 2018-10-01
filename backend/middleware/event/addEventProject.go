package event

import (
	"log"
	"net/http"
	"os"

	client "github.com/sea350/ustart_go/backend/middleware/client"
	postEvent "github.com/sea350/ustart_go/backend/post/event"
	types "github.com/sea350/ustart_go/backend/types"
)

//AddEventProject ... append project to event
func AddEventProject(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	projectID := r.FormValue("projectID")
	eventID := r.FormValue("eventID")

	project := types.EventProjects{ProjectID: projectID, Title: "Project Title Undefined (addEventProject)", Visible: true}

	err := postEvent.AppendProject(client.Eclient, eventID, project)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
} //end of AddEventProject
