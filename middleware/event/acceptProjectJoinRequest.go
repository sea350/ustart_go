package event

import (
	"fmt"
	"log"
	"net/http"
	"os"

	client "github.com/sea350/ustart_go/middleware/client"
	evntPost "github.com/sea350/ustart_go/post/event"
	projPost "github.com/sea350/ustart_go/post/project"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

//AcceptProjectJoinRequest ...
func AcceptProjectJoinRequest(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	eventID := r.FormValue("eventID")
	projectID := r.FormValue("projectID")

	newNumRequests, err := uses.RemoveProjectEventRequest(client.Eclient, eventID, projectID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	err = projPost.AppendEventID(client.Eclient, projectID, eventID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	var newProjectGuest types.EventProjectGuests
	newProjectGuest.Status = 0
	newProjectGuest.Visible = true
	newProjectGuest.ProjectID = projectID

	err = evntPost.AppendProjectGuest(client.Eclient, eventID, newProjectGuest)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	fmt.Fprintln(w, newNumRequests)

}
