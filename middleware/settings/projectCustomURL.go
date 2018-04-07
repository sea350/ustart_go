package settings

import (
	"fmt"
	"net/http"

	"github.com/sea350/ustart_go/middleware/client"

	get "github.com/sea350/ustart_go/get/project"
	uses "github.com/sea350/ustart_go/uses"
)

//ProjectCustomURL ... pushes a new banner image into ES
func ProjectCustomURL(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	newURL := r.FormValue("purl")

	projID := r.FormValue("projectID")

	inUse, err := get.URLInUse(client.Eclient, newURL)
	if err != nil {
		fmt.Println("err: middleware/settings/projectcustomurl line 28")
		fmt.Println(err)
	}

	proj, err := get.ProjectByID(client.Eclient, projID)
	if err != nil {
		fmt.Println("err: middleware/settings/projectcustomurl line 34")
		fmt.Println(err)
	}

	if inUse {
		fmt.Println("URL IS IN USE, ERROR NOT PROPERLY HANDLED REDIRECTING TO PROJECT PAGE")
		http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
		return
	}

	err = uses.ChangeProjectURL(eclient, projID, newURL)
	if err != nil {
		fmt.Println("err: middleware/settings/projectcustomurl line 46")
		fmt.Println(err)
	}

	http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
	return

}
