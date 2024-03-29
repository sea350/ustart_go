package settings

import (
	"net/http"

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
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	r.ParseForm()
	p := bluemonday.UGCPolicy()
	country := p.Sanitize(r.FormValue("country"))
	state := p.Sanitize(r.FormValue("state"))
	city := p.Sanitize(r.FormValue("city"))
	street := p.Sanitize(r.FormValue("street"))
	zip := p.Sanitize(r.FormValue("zip"))

	evnt, err := get.EventByID(client.Eclient, r.FormValue("eventID"))

	//TODO: DocID
	err = uses.ChangeEventLocation(client.Eclient, r.FormValue("eventID"), country, state, city, street, zip)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
	return

}
