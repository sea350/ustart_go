package profile

import (
	"fmt"
	"net/http"

	"github.com/sea350/ustart_go/middleware/stringHTML"
	"github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

//AddWidget ... Iunno
func AddWidget(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	username := test1.(string)

	r.ParseForm()
	if r.FormValue("formValue") == `0` {
		title := r.FormValue("title")
		description := r.FormValue("description")
		newWidget := types.Widget{UserID: session.Values["DocID"].(string), Title: title, Position: 0, Description: description, Classification: 0}
		err := uses.AddWidget(eclient, session.Values["DocID"].(string), newWidget)
		fmt.Println(err)
		fmt.Println("this is an error: middleware/profile/addWidget.go 27")
		output := stringHTML.WidgetText(title, description)
		fmt.Fprintln(w, output)
	}

	//contentArray := []rune(comment)
	//username := r.FormValue("username")

	http.Redirect(w, r, "/profile/"+username, http.StatusFound)
}
