package widget

import (
	"fmt"
	"net/http"

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

	widg, err := get.WidgetByID(client.Eclient, r.FormValue("deleteID"))
	if widg.Classification == 15 {
		http.Redirect(w, r, "/Projects/"+projectURL, http.StatusFound)
	}
	err = uses.RemoveWidget(client.Eclient, r.FormValue("deleteID"), true)
	if err != nil {
		fmt.Println("This is an err, deleteWidgetProject line29")
		fmt.Println(err)
	}

	http.Redirect(w, r, "/Projects/"+projectURL, http.StatusFound)
}
