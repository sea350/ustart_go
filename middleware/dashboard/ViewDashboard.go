package event

import (
	"log"
	"net/http"
	"os"

	userGet "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
)

//ViewDashboard ... rendering the dashboard
func ViewDashboard(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	userstruct, err := userGet.UserByID(client.Eclient, session.Values["DocID"].(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	// fmt.Println("URL IS ", r.URL.Path[7:])
	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string)}
	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "dash", cs)
}
