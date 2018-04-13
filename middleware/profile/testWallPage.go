package profile

import (
	"fmt"
	"net/http"

	"github.com/sea350/ustart_go/uses"

	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
)

//TestWallPage ... a page dedicated to testing only wall code
func TestWallPage(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	wallUsername := r.URL.Path[10:]
	wallID, err := get.IDByUsername(client.Eclient, wallUsername)
	if err != nil {
		fmt.Println("err middleware/profile/testwallpage line 22")
		fmt.Println(err)
	}
	user, err := get.UserByID(client.Eclient, wallID)
	if err != nil {
		fmt.Println("err middleware/profile/testwallpage line 28")
		fmt.Println(err)
	}
	entries, err := uses.LoadEntries(client.Eclient, user.EntryIDs[0:5])
	if err != nil {
		fmt.Println("err middleware/profile/testwallpage line 28")
		fmt.Println(err)
	}

	cs := client.ClientSide{Page: wallID, Wall: entries}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "profile-wall", cs)
}
