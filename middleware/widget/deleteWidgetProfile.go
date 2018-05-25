package widget

import (
	"fmt"
	"net/http"

	"github.com/sea350/ustart_go/get/widget"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//DeleteWidgetProfile ... Deletes a widget and redirects to profile page
func DeleteWidgetProfile(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	username := test1.(string)
	r.ParseForm()

	_, err := get.WidgetByID(client.Eclient, r.FormValue("deleteID"))
	err = uses.RemoveWidget(client.Eclient, r.FormValue("deleteID"), false)
	if err != nil {
		fmt.Println("This is an err, deleteWidgetProfile line24")
		fmt.Println(err)
	}

	http.Redirect(w, r, "/profile/"+username, http.StatusFound)
	return
}
