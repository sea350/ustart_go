package event

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	client "github.com/sea350/ustart_go/middleware/client"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"

	userGet "github.com/sea350/ustart_go/get/user"
)

//ViewEvent ... rendering the event
func ViewEvent(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	event, err := uses.AggregateEventData(client.Eclient, r.URL.Path[7:], test1.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	userstruct, err := userGet.UserByID(client.Eclient, session.Values["DocID"].(string))
	if err != nil {
		panic(err)
	}
	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string), Event: event}
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

	startDateOfEvent := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)
	startDate := r.FormValue("startDate")
	if len(startDate) > 15 {
		startDateOfEvent.Year, _ = strconv.Atoi(startDate[6:10])
		startDateOfEvent.Month, _ = time.Month(strconv.Atoi(startDate[0:2]))
		startDateOfEvent.Day, _ = strconv.Atoi(startDate[3:5])
		startDateOfEvent.Hour, _ = strconv.Atoi(startDate[11:13])
		startDateOfEvent.Minute, _ = strconv.Atoi(startDate[14:16])
	} else {
		log.Println("DateStart is Less than 15 Characters: ", startDate)
	}
	endDateOfEvent := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := r.FormValue("startDate")
	if len(endDate) > 15 {
		endDateOfEvent.Year, _ = strconv.Atoi(endDate[6:10])
		endDateOfEvent.Month, _ = time.Month(strconv.Atoi(endDate[0:2]))
		endDateOfEvent.Day, _ = strconv.Atoi(endDate[3:5])
		endDateOfEvent.Hour, _ = strconv.Atoi(endDate[11:13])
		endDateOfEvent.Minute, _ = strconv.Atoi(endDate[14:16])
	} else {
		log.Println("DateStart is Less than 15 Characters: ", endDate)
	}

	country := r.FormValue("country")
	state := r.FormValue("state")
	city := r.FormValue("city")
	zip := r.FormValue("zip")
	desc := []rune(r.FormValue("event_desc"))
	category := r.FormValue("category")

	var eventLocation types.LocStruct
	eventLocation.City = city
	eventLocation.Country = country
	eventLocation.Zip = zip
	eventLocation.State = state

	id, err := uses.CreateEvent(client.Eclient, title, desc, test1.(string), category, eventLocation, startDateOfEvent, endDateOfEvent)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
		http.Redirect(w, r, "/StartEvent/"+id, http.StatusFound)
	}

	http.Redirect(w, r, "/Event/"+id, http.StatusFound)
}
