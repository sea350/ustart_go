package settings

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/microcosm-cc/bluemonday"

	get "github.com/sea350/ustart_go/get/event"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//EventLocation ...
//For Events Location
func EventLocation(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	r.ParseForm()
	p := bluemonday.UGCPolicy()
	country := p.Sanitize(r.FormValue("country"))
	state := p.Sanitize(r.FormValue("state"))
	city := p.Sanitize(r.FormValue("city"))
	street := p.Sanitize(r.FormValue("street"))
	zip := p.Sanitize(r.FormValue("zip"))
	//   fmt.Println(blob)

	evnt, err := get.EventByID(client.Eclient, r.FormValue("eventID"))
	//fmt.Println(reflect.TypeOf(blob))
	//TODO: DocID
	err = uses.ChangeEventLocation(client.Eclient, r.FormValue("eventID"), country, state, city, street, zip)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
	return

}
