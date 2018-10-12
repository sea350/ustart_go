package widget

import (
	"encoding/json"
	"log"
	"net/http"

	getEvent "github.com/sea350/ustart_go/get/event"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/event"
	"github.com/sea350/ustart_go/uses"
)

//SortEventWidgets ... gets new array of widget ids from project page and updates project struct in ES
func SortEventWidgets(w http.ResponseWriter, r *http.Request) {
	// If followingStatus = no
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	r.ParseForm()
	sortedWidgets := r.FormValue("sortedWidgets")
	eventURL := r.FormValue("pageID")
	if eventURL == `` {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("Event URL not passed in")
		http.Redirect(w, r, "/404/", http.StatusFound)
		return
	}

	defer http.Redirect(w, r, "/Events/"+eventURL, http.StatusFound)

	arr := []string{}
	err := json.Unmarshal([]byte(sortedWidgets), &arr)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	id, err := getEvent.EventIDByURL(client.Eclient, eventURL)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	event, member, err := getEvent.EventAndMember(client.Eclient, id, docID.(string))

	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	if uses.HasEventPrivilege("widget", event.PrivilegeProfiles, member) {
		err = post.UpdateEvent(client.Eclient, id, "Widgets", arr)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
	} else {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("You do not have the privilege to add a widget to this event. Check your privilege.")
	}
}
