package event

import (
	"log"
	"net/http"

	get "github.com/sea350/ustart_go/get/event"
	client "github.com/sea350/ustart_go/middleware/client"
	evntPost "github.com/sea350/ustart_go/post/event"
	userPost "github.com/sea350/ustart_go/post/user"
)

//MemberRequestToJoin ...
func MemberRequestToJoin(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	id := r.FormValue("eventID") //event docID
	if id == `` {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("Event ID not passed in")
		return
	}

	evnt, err := get.EventByID(client.Eclient, id)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	for _, memberInfo := range evnt.Members {
		if memberInfo.MemberID == test1.(string) {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println("user is already a member")
			http.Redirect(w, r, "/Events/"+evnt.URLName, http.StatusFound)
			return
		}
	}
	for _, receivedReq := range evnt.MemberReqReceived {
		if receivedReq == test1.(string) {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println("user's request already received")
			http.Redirect(w, r, "/Events/"+evnt.URLName, http.StatusFound)
			return
		}
	}

	err = userPost.AppendSentEventReq(client.Eclient, test1.(string), id)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	err = evntPost.AppendMemberReqReceived(client.Eclient, id, test1.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	http.Redirect(w, r, "/Events/"+evnt.URLName, http.StatusFound)
	return
}
