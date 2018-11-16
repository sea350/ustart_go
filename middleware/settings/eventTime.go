package settings

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/microcosm-cc/bluemonday"

	get "github.com/sea350/ustart_go/get/event"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//EventTime ...
//For Events Time
func EventTime(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	r.ParseForm()

	p := bluemonday.UGCPolicy()
	sClean := p.Sanitize(r.FormValue("startDate"))
	Syear, _ := strconv.Atoi(sClean[0:4])
	Smonth, _ := strconv.Atoi(sClean[5:7])
	Sday, _ := strconv.Atoi(sClean[8:10])
	Shour, _ := strconv.Atoi(sClean[11:13])
	Sminute, _ := strconv.Atoi(sClean[14:16])
	Sdate := time.Date(Syear, time.Month(Smonth), Sday, Shour, Sminute, 0, 0, time.UTC)

	eClean := p.Sanitize(r.FormValue("endDate"))
	Eyear, _ := strconv.Atoi(eClean[0:4])
	Emonth, _ := strconv.Atoi(eClean[5:7])
	Eday, _ := strconv.Atoi(eClean[8:10])
	Ehour, _ := strconv.Atoi(eClean[11:13])
	Eminute, _ := strconv.Atoi(eClean[14:16])
	Edate := time.Date(Eyear, time.Month(Emonth), Eday, Ehour, Eminute, 0, 0, time.UTC)

	evnt, err := get.EventByID(client.Eclient, r.FormValue("eventID"))
	//TODO: DocID
	err = uses.ChangeEventTime(client.Eclient, r.FormValue("eventID"), Sdate, Edate)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	//TODO: Add in right URL
	http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
	return

}
