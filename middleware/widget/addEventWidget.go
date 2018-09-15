package widget

import (
	"fmt"
	"log"
	"net/http"

	getEvnt "github.com/sea350/ustart_go/get/event"

	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/widget"
	"github.com/sea350/ustart_go/uses"
)

//AddEventWidget ... After widget form submission adds a widget to database
func AddEventWidget(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		//No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	evnt, member, err := getEvnt.EventAndMember(client.Eclient, r.FormValue("eventWidget"), test1.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	// newWidget, err := ProcessWidgetForm(r)
	// if err != nil {
	// 	fmt.Println("this is an error: middleware/widget/addEventWidget 31")
	// 	fmt.Println(err)
	// 	http.Redirect(w, r, "/Events/"+evnt.URLName, http.StatusFound)
	// 	return
	// }

	// newWidget.UserID = r.FormValue("eventWidget")

	if uses.HasEventPrivilege("widget", evnt.PrivilegeProfiles, member) {
		// if r.FormValue("editID") == `0` {
		// 	fmt.Println("this is debug text middeware/widget/addEventWidget.go")
		// 	fmt.Println(r.FormValue("eventWidget"))
		// err := uses.AddWidget(client.Eclient, r.FormValue("eventWidget"), newWidget, true)
		newWidget, err := ProcessWidgetForm(r)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			http.Redirect(w, r, "/Events/"+evnt.URLName, http.StatusFound)
		}

		newWidget.UserID = r.FormValue("eventWidget")

		if r.FormValue("editID") == `0` {
			// fmt.Println("this is debug text middeware/widget/addeventwidget.go")
			// fmt.Println(r.FormValue("eventWidget"))
			// fmt.Println(newWidget.Data)

			err := uses.AddWidget(client.Eclient, r.FormValue("eventWidget"), newWidget, false, true)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
		} else {
			err := post.ReindexWidget(client.Eclient, r.FormValue("editID"), newWidget)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
		}

		http.Redirect(w, r, "/Events/"+evnt.URLName, http.StatusFound)
	} else {
		fmt.Println("You do not have the privilege to add a widget to this event. Check your privilege. ")
	}
	return
}
