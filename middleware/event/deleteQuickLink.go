package event

import (
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	get "github.com/sea350/ustart_go/get/event"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/event"
	"github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

//DeleteEventQuickLink ...
func DeleteEventQuickLink(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		//No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	p := bluemonday.UGCPolicy()

	ID := r.FormValue("eventID")
	deleteTitle := p.Sanitize(r.FormValue("deleteEventLinkDesc"))
	deleteURL := p.Sanitize(r.FormValue("deleteEventLink"))
	if ID == `` || deleteURL == `` {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Crucial data was not passed in, now exiting")
		return
	}

	evnt, err := get.EventByID(client.Eclient, ID)
	if err != nil {
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

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

	if len(evnt.QuickLinks) == 1 {
		err := post.UpdateEvent(client.Eclient, ID, "QuickLinks", newArr)
		if err != nil {
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
		http.Redirect(w, r, "/Events/"+evnt.URLName, http.StatusFound)
		return
	}

	target := -1
	for index, link := range evnt.QuickLinks {
		if link.Name == deleteTitle && link.URL == deleteURL {
			target = index
			break
		}
	}

	if target == -1 {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "deleted object not found")
		newArr = evnt.QuickLinks
	} else if (target + 1) < len(evnt.QuickLinks) {
		newArr = append(evnt.QuickLinks[:target], evnt.QuickLinks[(target+1):]...)
	} else {
		newArr = evnt.QuickLinks[:target]
	}

	err = post.UpdateEvent(client.Eclient, ID, "QuickLinks", newArr)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	http.Redirect(w, r, "/Events/"+evnt.URLName, http.StatusFound)
	return
}
