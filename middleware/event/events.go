package event

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"

	client "github.com/sea350/ustart_go/middleware/client"
)

//EventsPage ... Displays the events page
func EventsPage(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	event, err := uses.AggregateEventData(client.Eclient, r.URL.Path[10:], test1.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	widgets, errs := uses.LoadWidgets(client.Eclient, event.EventData.Widgets)
	if len(errs) > 0 {
		log.Println("there were one or more errors loading widgets")
		for _, eror := range errs {
			fmt.Println(eror)
		}
	}
	userstruct, err := get.UserByID(client.Eclient, session.Values["DocID"].(string))
	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string), Event: event, Widgets: widgets}
	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "eventsF", cs)
}

//MyEvents ...
func MyEvents(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	var heads []types.FloatingHead

	userstruct, err := get.UserByID(client.Eclient, session.Values["DocID"].(string))
	if err != nil {
		panic(err)
	}

	for _, eventInfo := range userstruct.Events {
		head, err := uses.ConvertEventToFloatingHead(client.Eclient, eventInfo.EventID)
		if err != nil {
			panic(err)
		}
		heads = append(heads, head)
	}

	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string), ListOfHeads: heads}
	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "manageevents-Nil", cs)
}

//CreateEventPage ...
func CreateEventPage(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	userstruct, err := get.UserByID(client.Eclient, session.Values["DocID"].(string))
	if err != nil {
		panic(err)
	}
	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string)}

	layout := "2006-01-02T15:04:05.000Z"

	title := r.FormValue("event_title")
	description := []rune(r.FormValue("event_desc"))
	category := r.FormValue("category")
	//location := r.FormValue("location")
	customURL := r.FormValue("curl")
	eventStartForm := r.FormValue("event_start")
	eventEndForm := r.FormValue("event_end")
	country := r.FormValue("country")
	state := r.FormValue("state")
	city := r.FormValue("city")
	zip := r.FormValue("zip")

	eventStart, err := time.Parse(layout, eventStartForm)
	if err != nil {
		panic(err)
	}
	eventEnd, err := time.Parse(layout, eventEndForm)
	if err != nil {
		panic(err)
	}

	var eventLocation types.LocStruct
	eventLocation.City = city
	eventLocation.Country = country
	eventLocation.Zip = zip
	eventLocation.State = state

	if title != `` {
		//proper URL
		if !uses.ValidUsername(customURL) {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, "Check me please")
			cs.ErrorStatus = true
			cs.ErrorOutput = err
		}

		url, err := uses.CreateEvent(client.Eclient, title, description, session.Values["DocID"].(string), category, customURL, eventLocation, eventStart, eventEnd)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
			cs.ErrorStatus = true
			cs.ErrorOutput = err
		} else {
			time.Sleep(5000)
			http.Redirect(w, r, "/Events/"+url, http.StatusFound)
			return
		}
	}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "createEvent-Nil", cs)
}
