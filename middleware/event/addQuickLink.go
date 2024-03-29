package event

import (
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	get "github.com/sea350/ustart_go/get/event"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/event"
	"github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
)

//AddEventQuickLink ...
func AddEventQuickLink(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	ID := r.FormValue("eventID")
	if ID == `` {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Crucial data was not passed in, now exiting")
		return
	}

	evnt, err := get.EventByID(client.Eclient, ID)
	if err != nil {
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
	p := bluemonday.UGCPolicy()

	var exists bool
	var member types.EventMembers
	for _, mem := range evnt.Members {
		if mem.MemberID == session.Values["DocID"].(string) {
			exists = true
			member = mem
			break
		}
	}

	if !exists {
		return
	}

	hasPermission := uses.HasEventPrivilege("links", evnt.PrivilegeProfiles, member)

	if !hasPermission {
		return
	}

	var newArr []types.Link

	if len(evnt.QuickLinks) <= 1 {
		err := post.UpdateEvent(client.Eclient, ID, "QuickLinks", newArr)
		if err != nil {
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
		return
	}

	evnt.QuickLinks = append(evnt.QuickLinks, types.Link{Name: p.Sanitize(r.FormValue("eventLinkDesc")), URL: p.Sanitize(r.FormValue("eventLink"))})

	err = post.UpdateEvent(client.Eclient, ID, "QuickLinks", evnt.QuickLinks)
	if err != nil {
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
}
