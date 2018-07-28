package widget

import (
	"log"
	"net/http"
	"os"

	"github.com/sea350/ustart_go/get/widget"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//DeleteWidgetProject ... Deletes a widget and redirects to projects page
func DeleteWidgetProject(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	projectURL := r.FormValue("deleteProjectURL")

	_, err := get.WidgetByID(client.Eclient, r.FormValue("deleteID"))
	err = uses.RemoveWidget(client.Eclient, r.FormValue("deleteID"), true, false)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	http.Redirect(w, r, "/Projects/"+projectURL, http.StatusFound)
	return
}
