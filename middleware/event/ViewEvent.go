package event

import (
	"log"
	"net/http"
	"os"
	"time"

	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/event"
	types "github.com/sea350/ustart_go/types"

	userGet "github.com/sea350/ustart_go/get/user"
	userPost "github.com/sea350/ustart_go/post/user"
)

//ViewEvent ... rendering the event
func ViewEvent(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	cs := client.ClientSide{}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "events", cs)
}

//StartEvent ... rendering the event form
func StartEvent(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	cs := client.ClientSide{}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "eventStart", cs)
}

//AddEvent ... append event to database
func AddEvent(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()

	title := r.FormValue("title")
	dateStart := r.FormValue("dateStart")
	dateEnd := r.FormValue("dateEnd")
	country := r.FormValue("country")
	state := r.FormValue("state")
	city := r.FormValue("city")
	zip := r.FormValue("zip")

	var eventLocation types.LocStruct
	eventLocation.City = city
	eventLocation.Country = country
	eventLocation.Zip = zip
	eventLocation.State = state

	layout := "2006-01-02T15:04:05.000Z"

	usr, err := userGet.UserByID(client.Eclient, test1.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	var newEvent types.Events
	newEvent.Name = title
	newEvent.EventDateStart, _ = time.Parse(layout, dateStart)
	newEvent.EventDateEnd, _ = time.Parse(layout, dateEnd)
	newEvent.Location = eventLocation
	newEvent.CreationDate = time.Now()
	newEvent.Host = test1.(string)
	newEvent.Visible = true

	id, err := post.IndexEvent(client.Eclient, newEvent)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	var newEventInfo types.EventInfo
	newEventInfo.EventID = id
	newEventInfo.Visible = true

	err = userPost.UpdateUser(client.Eclient, test1.(string), "Events", append(usr.Events, newEventInfo))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	//cs := client.ClientSide{ErrorStatus: false}

	http.Redirect(w, r, "/Event/"+id, http.StatusFound)
}
