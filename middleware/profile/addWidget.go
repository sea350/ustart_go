package profile

import (
	"fmt"
	"net/http"

	"github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
)

//AddWidget ... After widget form submission adds a widget to database
func AddWidget(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	username := test1.(string)
	r.ParseForm()

	if r.FormValue("widgetSubmit") == `0` {
		title := r.FormValue("customHeader")
		description := r.FormValue("customContent")
		data := []string{title, description}
		//call fuction that gets the next available slot in user's widgets
		//position == len(user.widgets)
		newWidget := types.Widget{UserID: session.Values["DocID"].(string), Data: data, Position: 0, Classification: 0}
		err := uses.AddWidget(eclient, session.Values["DocID"].(string), newWidget)
		if err != nil {
			fmt.Println(err)
			fmt.Println("this is an error: middleware/profile/addWidget.go 34")
		}
	}

	//contentArray := []rune(comment)
	//username := r.FormValue("username")

	http.Redirect(w, r, "/profile/"+username, http.StatusFound)
}
