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
	fmt.Println("test1", test1)
	if test1 == nil {
		//No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	evnt, member, err := getEvnt.EventAndMember(client.Eclient, r.FormValue("eventWidget"), test1.(string))
	fmt.Println("evnt", evnt)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	if uses.HasEventPrivilege("widget", evnt.PrivilegeProfiles, member) {

		/* empty eventWidget */
		if r.FormValue("eventWidget") == `` {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println("Event ID Missing")
			http.Redirect(w, r, "/Events/"+evnt.URLName, http.StatusFound)
			return
		}

		newWidget, err := ProcessWidgetForm(r)
		fmt.Println("newWidget", newWidget)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			http.Redirect(w, r, "/Events/"+evnt.URLName, http.StatusFound)
		}

		newWidget.UserID = r.FormValue("eventWidget")
		fmt.Println("newWidget.UserID", newWidget.UserID)
		if r.FormValue("editID") == `0` {
			err := uses.AddWidget(client.Eclient, r.FormValue("eventWidget"), newWidget, false, true)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
		} else {
			err := post.ReindexWidget(client.Eclient, r.FormValue("editID"), newWidget)

			fmt.Println("line 60, ", err)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
		}
		fmt.Println("CAN YOU COME HERE?")
		http.Redirect(w, r, "/Events/"+evnt.URLName, http.StatusFound)
	} else {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("You do not have the privilege to add a widget to this event. Check your privilege. ")
		http.Redirect(w, r, "/Events/"+evnt.URLName, http.StatusFound)
	}
}
