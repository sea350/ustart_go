package widget

import (
	"fmt"
	"html/template"
	"net/http"

	get "github.com/sea350/ustart_go/get/widget"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//EditWidgetInstaDelete ... deletes a link on an instagram widget
func EditWidgetInstaDelete(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	username := test1.(string)
	deletedURL := template.HTML(r.FormValue("instaURL"))
	widget, err := get.WidgetByID(client.Eclient, r.FormValue("editID"))
	if err != nil {
		fmt.Println(err)
		fmt.Println("this is an err, editInstaAdd line 24")
	}

	var target int
	for index, link := range widget.Data {
		if link == deletedURL {
			target = index
			break
		}
	}

	var newArr []template.HTML
	if target+1 < len(widget.Data) {
		newArr = append(widget.Data[:target], widget.Data[target+1:]...)
	} else {
		newArr = widget.Data[:target]
	}

	err = uses.EditWidget(client.Eclient, r.FormValue("editID"), newArr)
	http.Redirect(w, r, "/profile/"+username, http.StatusFound)
}
