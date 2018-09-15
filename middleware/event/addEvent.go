package event

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/microcosm-cc/bluemonday"
	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

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
	p := bluemonday.UGCPolicy()

	title := p.Sanitize(r.FormValue("title"))
	college := p.Sanitize(r.FormValue("universityName"))
	customURL := p.Sanitize(r.FormValue("eventurl"))
	category := r.FormValue("eventCategory")

	//endDateOfEvent
	//eventLocation
	country := r.FormValue("country")
	state := r.FormValue("state")
	city := p.Sanitize(r.FormValue("city"))
	zip := p.Sanitize(r.FormValue("zip"))
	street := p.Sanitize(r.FormValue("street"))
	lat := r.FormValue("lat")
	lng := r.FormValue("lng")

	desc := []rune(p.Sanitize(r.FormValue("event_desc")))

	Syear, _ := strconv.Atoi(r.FormValue("startDate")[0:4])
	Smonth, _ := strconv.Atoi(r.FormValue("startDate")[5:7])
	Sday, _ := strconv.Atoi(r.FormValue("startDate")[8:10])
	Shour, _ := strconv.Atoi(r.FormValue("startDate")[11:13])
	Sminute, _ := strconv.Atoi(r.FormValue("startDate")[14:16])
	startDateOfEvent := time.Date(Syear, time.Month(Smonth), Sday, Shour, Sminute, 0, 0, time.UTC)

	Eyear, _ := strconv.Atoi(r.FormValue("endDate")[0:4])
	Emonth, _ := strconv.Atoi(r.FormValue("endDate")[5:7])
	Eday, _ := strconv.Atoi(r.FormValue("endDate")[8:10])
	Ehour, _ := strconv.Atoi(r.FormValue("endDate")[11:13])
	Eminute, _ := strconv.Atoi(r.FormValue("endDate")[14:16])
	endDateOfEvent := time.Date(Eyear, time.Month(Emonth), Eday, Ehour, Eminute, 0, 0, time.UTC)

	var eventLocation types.LocStruct
	eventLocation.Street = street
	eventLocation.City = city
	eventLocation.Country = country
	eventLocation.Zip = zip
	eventLocation.State = state
	eventLocation.Street = street
	eventLocation.Lat = lat
	eventLocation.Lng = lng

	if title != `` {
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
			time.Sleep(2 * time.Second)
			http.Redirect(w, r, "/Event/"+url, http.StatusFound)
			return
		}
	}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "eventStart", cs)
} //end of AddEvent
