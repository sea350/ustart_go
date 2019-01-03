package settings

import (
	"fmt"
	
	"net/http"
	

	"github.com/microcosm-cc/bluemonday"

	"github.com/sea350/ustart_go/middleware/client"

	get "github.com/sea350/ustart_go/get/event"
	uses "github.com/sea350/ustart_go/uses"
)

//EventCustomURL ... pushes a new banner image into ES
func EventCustomURL(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	r.ParseForm()
	p := bluemonday.UGCPolicy()
	newURL := p.Sanitize(r.FormValue("purl"))
	if len(newURL) < 1 {
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"Cannot have a blank URL!")
		return

	}

	evntID := r.FormValue("eventID")

	inUse, err := get.EventURLInUse(client.Eclient, newURL)
	if err != nil {
		

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	evnt, err := get.EventByID(client.Eclient, evntID)
	if err != nil {
		

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	if inUse {
		
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"URL IS IN USE, ERROR NOT PROPERLY HANDLED REDIRECTING TO EVENT PAGE")
		http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
		return
	}

	err = uses.ChangeEventURL(client.Eclient, evntID, newURL)
	if err != nil {
		

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
	return

}
