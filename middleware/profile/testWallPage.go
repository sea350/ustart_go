package profile

import (
	"net/http"

	"github.com/sea350/ustart_go/uses"

	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
)

//TestWallPage ... a page dedicated to testing only wall code
func TestWallPage(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	wallUsername := r.URL.Path[10:]
	wallID, err := get.IDByUsername(client.Eclient, wallUsername)
	if err != nil {
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
	user, err := get.UserByID(client.Eclient, wallID)
	if err != nil {
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
	entries, err := uses.LoadEntries(client.Eclient, user.EntryIDs, docID.(string))
	if err != nil {
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	viewer, err := get.UserByID(client.Eclient, docID.(string))
	if err != nil {
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	cs := client.ClientSide{UserInfo: user, Page: wallID, Wall: entries, ImageCode: viewer.Avatar}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "profile-wall", cs)
}
