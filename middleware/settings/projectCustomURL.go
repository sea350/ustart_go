package settings

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/microcosm-cc/bluemonday"

	"github.com/sea350/ustart_go/middleware/client"

	get "github.com/sea350/ustart_go/get/project"
	uses "github.com/sea350/ustart_go/uses"
)

//ProjectCustomURL ... pushes a new banner image into ES
func ProjectCustomURL(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	p := bluemonday.UGCPolicy()
	newURL := p.Sanitize(r.FormValue("purl"))
	if len(newURL) < 1 {
		log.Println("URL cannot be blank!")
		return
	}
	projID := r.FormValue("projectID")

	inUse, err := get.URLInUse(client.Eclient, newURL)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	proj, err := get.ProjectByID(client.Eclient, projID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	if inUse {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("URL IS IN USE, ERROR NOT PROPERLY HANDLED REDIRECTING TO PROJECT PAGE")
		http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
		return
	}

	err = uses.ChangeProjectURL(client.Eclient, projID, newURL)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		time.Sleep(2 * time.Second)
		http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
		return
	}

	time.Sleep(2 * time.Second)
	http.Redirect(w, r, "/Projects/"+newURL, http.StatusFound)

}
