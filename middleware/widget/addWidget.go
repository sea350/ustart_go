package widget

import (
	"log"
	"net/http"
	"os"

	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/widget"
	"github.com/sea350/ustart_go/uses"
)

//AddWidget ... After widget form submission adds a widget to database
func AddWidget(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	username := test1.(string)

	newWidget, err := ProcessWidgetForm(r)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
		http.Redirect(w, r, "/profile/"+username, http.StatusFound)
		return
	}

	newWidget.UserID = session.Values["DocID"].(string)

	if r.FormValue("editID") == `0` {
		err := uses.AddWidget(client.Eclient, session.Values["DocID"].(string), newWidget, false, false)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
	} else {
		err := post.ReindexWidget(client.Eclient, r.FormValue("editID"), newWidget)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
	}

	http.Redirect(w, r, "/profile/"+username, http.StatusFound)
	return
}
