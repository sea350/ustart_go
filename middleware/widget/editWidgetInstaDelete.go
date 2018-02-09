package widget

import (
	"fmt"
	"html/template"
	"net/http"

	get "github.com/sea350/ustart_go/get/widget"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//EditWidgetDataDelete ... deletes a link in a class 4 or 5 widget widget
func EditWidgetDataDelete(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	username := test1.(string)

	deletedURL := template.HTML(r.FormValue("instaURL"))
	oldWidget, err := get.WidgetByID(client.Eclient, r.FormValue("editID"))
	if err != nil {
		fmt.Println(err)
		fmt.Println("this is an err, editInstaAdd line 26")
	}

	if len(oldWidget.Data) == 1 {
		err = uses.RemoveWidget(client.Eclient, r.FormValue("editID"))
		if err != nil {
			fmt.Println(err)
			fmt.Println("this is an err, editInstaAdd line 34")
		}
		http.Redirect(w, r, "/profile/"+username, http.StatusFound)
		return
	}

	var target int
	for index, link := range oldWidget.Data {
		if link == deletedURL {
			fmt.Println(link)
			fmt.Println(index)
			fmt.Println("debug text: middleware/widget/editWidgetDelete line 45")
			target = index
			break
		}
	}

	var newArr []template.HTML
	if target+1 < len(oldWidget.Data) {
		newArr = append(oldWidget.Data[:target], oldWidget.Data[target+1:]...)
	} else {
		newArr = oldWidget.Data[:target]
	}

	err = uses.EditWidget(client.Eclient, r.FormValue("editID"), newArr)
	if err != nil {
		fmt.Println(err)
		fmt.Println("this is an err, editInstaAdd line 58")
	}
	http.Redirect(w, r, "/profile/"+username, http.StatusFound)
}
