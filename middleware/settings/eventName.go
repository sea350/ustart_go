package settings

import (
	"net/http"

	"github.com/microcosm-cc/bluemonday"

	get "github.com/sea350/ustart_go/get/event"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//EventChangeNameAndDescription ...
//For Events
func EventChangeNameAndDescription(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	r.ParseForm()
	p := bluemonday.UGCPolicy()
	evntName := p.Sanitize(r.FormValue("pname"))

	cleanDesc := p.Sanitize(r.FormValue("inputDesc"))
	evntDesc := []rune(cleanDesc)

	evnt, err := get.EventByID(client.Eclient, r.FormValue("eventID"))
	//TODO: DocID
	err = uses.ChangeEventNameAndDescription(client.Eclient, r.FormValue("eventID"), evntName, evntDesc)
	if err != nil {
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}
	//TODO: Add in right URL
	http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
	return

}
