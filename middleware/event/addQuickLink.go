package event

import (
	"log"
	"net/http"

	get "github.com/sea350/ustart_go/get/event"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/event"
	"github.com/sea350/ustart_go/types"
)

//AddEventQuickLink ...
func AddEventQuickLink(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	ID := r.FormValue("eventID")

	evnt, err := get.EventByID(client.Eclient, ID)
	if err != nil {
		log.Println("middleware/event/addQuickLink line 24")
		log.Println(err)
	}

	evnt.QuickLinks = append(evnt.QuickLinks, types.Link{Name: r.FormValue("eventLinkDesc"), URL: r.FormValue("eventLink")})

	err = post.UpdateEvent(client.Eclient, ID, "QuickLinks", evnt.QuickLinks)
	if err != nil {
		log.Println("middleware/event/addQuickLink line 32")
		log.Println(err)
	}
}
