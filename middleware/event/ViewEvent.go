package event

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

//ViewEvent ... rendering the event
//ProjectsPage
func ViewEvent(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	fmt.Println("Event", test1)

	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	event, err := uses.AggregateEventData(client.Eclient, r.URL.Path[7:], test1.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		//log.Println(err)
	}
	widgets, errs := uses.LoadWidgets(client.Eclient, event.EventData.Widgets)
	if len(errs) > 0 {
		log.Println("there were one or more errors loading widgets")
		for _, eror := range errs {
			fmt.Println(eror)
		}
	}

	userstruct, err := get.UserByID(client.Eclient, test1.(string))
	if err != nil {
		panic(err)
	}
	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string), Event: event, Widgets: widgets}
	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "events", cs)
}

//StartEvent ... rendering the event form
//equivalent of CreateProjectPage
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
//equivalent of CreateProjectPage
func AddEvent(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	userstruct, err := get.UserByID(client.Eclient, test1.(string))
	if err != nil {
		panic(err)
	}
	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string)}
	//r.ParseForm()

	title := r.FormValue("title")
	college := r.FormValue("universityName")
	customURL := r.FormValue("eventurl")
	category := r.FormValue("eventCategory")

	//endDateOfEvent
	//eventLocation
	country := r.FormValue("country")
	state := r.FormValue("state")
	city := r.FormValue("city")
	zip := r.FormValue("zip")
	street := r.FormValue("street")

	desc := []rune(r.FormValue("event_desc"))

	startDATE := r.FormValue("startDate")
	fmt.Println(startDATE)
	Syear, _ := strconv.Atoi(r.FormValue("startDate")[0:4])
	Smonth, _ := strconv.Atoi(r.FormValue("startDate")[5:7])
	Sday, _ := strconv.Atoi(r.FormValue("startDate")[8:10])
	Shour, _ := strconv.Atoi(r.FormValue("startDate")[11:13])
	Sminute, _ := strconv.Atoi(r.FormValue("startDate")[14:16])
	startDateOfEvent := time.Date(Syear, time.Month(Smonth), Sday, Shour, Sminute, 0, 0, time.UTC)
	/*
		startDate := r.FormValue("startDate")
		if len(startDate) > 15 {
			year, _ := strconv.Atoi(r.FormValue("startDate")[0:4])
			month, _ := strconv.Atoi(r.FormValue("startDate")[5:7])
			day, _ := strconv.Atoi(r.FormValue("startDate")[8:10])
		}
	*/
	/*
		if len(startDate) > 15 {
			year, _ := strconv.Atoi(startDate[6:10])
			month, _ := strconv.Atoi(startDate[0:2])
			day, _ := strconv.Atoi(startDate[3:5])
			hour, _ := strconv.Atoi(startDate[11:13])
			minute, _ := strconv.Atoi(startDate[14:16])
			startDateOfEvent = time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)
		}
	*/

	Eyear, _ := strconv.Atoi(r.FormValue("endDate")[0:4])
	Emonth, _ := strconv.Atoi(r.FormValue("endDate")[5:7])
	Eday, _ := strconv.Atoi(r.FormValue("endDate")[8:10])
	Ehour, _ := strconv.Atoi(r.FormValue("endDate")[11:13])
	Eminute, _ := strconv.Atoi(r.FormValue("endDate")[14:16])
	endDateOfEvent := time.Date(Eyear, time.Month(Emonth), Eday, Ehour, Eminute, 0, 0, time.UTC)
	/*
		endDateOfEvent := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)
		if len(endDate) > 15 {
			year, _ := strconv.Atoi(endDate[6:10])
			month, _ := strconv.Atoi(endDate[0:2])
			day, _ := strconv.Atoi(endDate[3:5])
			hour, _ := strconv.Atoi(endDate[11:13])
			minute, _ := strconv.Atoi(endDate[14:16])
			endDateOfEvent = time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)
		}
	*/

	var eventLocation types.LocStruct
	eventLocation.Street = street
	eventLocation.City = city
	eventLocation.Country = country
	eventLocation.Zip = zip
	eventLocation.State = state
	eventLocation.Street = street

	if title != `` {
		//proper URL
		if !uses.ValidUsername(customURL) {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println("Invalid custom event URL")
			cs.ErrorStatus = true
			cs.ErrorOutput = errors.New("Invalid custom event URL")
			client.RenderSidebar(w, r, "template2-nil")
			client.RenderSidebar(w, r, "leftnav-nil")
			client.RenderTemplate(w, r, "eventStart", cs)
			return
		}
		url, err := uses.CreateEvent(client.Eclient, title, desc, test1.(string), category, eventLocation, startDateOfEvent, endDateOfEvent, college, customURL)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		} else {
			http.Redirect(w, r, "/Event/"+url, http.StatusFound)
			return
		}
	}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "eventStart", cs)
} //end of AddEvent
