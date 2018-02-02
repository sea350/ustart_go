package widget

import (
	"fmt"
	"html/template"
	"net/http"

	get "github.com/sea350/ustart_go/get/widget"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//EditWidgetInstaAdd ... adds a new link to an instagram widget
func EditWidgetInstaAdd(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	username := test1.(string)
	newURL := template.HTML(r.FormValue("UNKNOWN"))
	widget, err := get.WidgetByID(client.Eclient, r.FormValue("editID"))
	if err != nil {
		fmt.Println(err)
		fmt.Println("this is an err, editInstaAdd line 24")
	}

	newArr := append(widget.Data, newURL)
	err = uses.EditWidget(client.Eclient, r.FormValue("editID"), newArr)
	http.Redirect(w, r, "/profile/"+username, http.StatusFound)
}
