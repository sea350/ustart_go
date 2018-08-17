package event

import (
	"log"
	"net/http"
	"os"

	getEvent "github.com/sea350/ustart_go/get/event"
	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/uses"

	"github.com/sea350/ustart_go/middleware/client"
	types "github.com/sea350/ustart_go/types"
)

//ManageEvents ...
func ManageEvents(w http.ResponseWriter, r *http.Request) {
	//log.Println("ManageEvents?")
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	var heads []types.FloatingHead

	userstruct, err := get.UserByID(client.Eclient, session.Values["DocID"].(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	for _, eventInfo := range userstruct.Events {
		if eventInfo.EventID == "" {
			continue //Missing EventID, meaning BAD ID
		}
		evnt, err := getEvent.EventByID(client.Eclient, eventInfo.EventID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
		for _, memberInfo := range evnt.Members {
			if memberInfo.MemberID == test1.(string) && memberInfo.Role > 1 {
				//finds user in the list of members and also checks if they have creator rank
				continue
				//head.Followed in this case expresses whether or not they have edit permissions
			}
		}
		head, err := uses.ConvertEventToFloatingHead(client.Eclient, eventInfo.EventID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
		heads = append(heads, head)

		//fmt.Println(heads)

	}

	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string), ListOfHeads: heads}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "eventManager", cs)
}
