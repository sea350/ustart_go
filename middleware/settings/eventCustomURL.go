package settings

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
		log.Println("Cannot have a blank URL!")
		return

	}

	evntID := r.FormValue("eventID")

	inUse, err := get.EventURLInUse(client.Eclient, newURL)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	evnt, err := get.EventByID(client.Eclient, evntID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	if inUse {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("URL IS IN USE, ERROR NOT PROPERLY HANDLED REDIRECTING TO EVENT PAGE")
		http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
		return
	}

	err = uses.ChangeEventURL(client.Eclient, evntID, newURL)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
	return

}
