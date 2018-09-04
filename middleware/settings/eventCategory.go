package settings

import (
	"log"
	"net/http"

	"github.com/microcosm-cc/bluemonday"

	get "github.com/sea350/ustart_go/get/event"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/event"
)

//EventCategory ...
func EventCategory(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	p := bluemonday.UGCPolicy()
	newCategory := p.Sanitize(r.FormValue("type_select"))
	if len(newCategory) == 0 {
		log.Println("Invalid category")
		return
	}

	evntID := r.FormValue("eventID")
	proj, err := get.EventByID(client.Eclient, evntID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	err = post.UpdateEvent(client.Eclient, evntID, "Category", newCategory)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	http.Redirect(w, r, "/EventSettings/"+proj.URLName, http.StatusFound)
	return

}
