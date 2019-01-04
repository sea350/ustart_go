package event

import (
	
	"net/http"

	get "github.com/sea350/ustart_go/get/event"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/event"
)

//GuestRequestToJoin ... PUBLIC no need for request received
func GuestRequestToJoin(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	id := r.FormValue("eventID") //event docID
	if id == `` {
		
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"Event ID not passed in")
		return
	}

	evnt, err := get.EventByID(client.Eclient, id)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		return
	}

	// 
	// 		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"Debug check 1/2")

	for _, guestInfo := range evnt.Guests {
		if guestInfo.GuestID == test1.(string) {
			
					client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"user is already a guest")
			http.Redirect(w, r, "/Event/"+evnt.URLName, http.StatusFound)
			return
		}
	}
	for receivedReq := range evnt.GuestReqReceived {
		if receivedReq == test1.(string) {
			
					client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"user has already sent a request")
			http.Redirect(w, r, "/Event/"+evnt.URLName, http.StatusFound)
			return
		}
	}

	err = post.AppendGuestReqReceived(client.Eclient, id, test1.(string), 1)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	// 
	// 		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"Debug check 2/2")

	http.Redirect(w, r, "/Event/"+evnt.URLName, http.StatusFound)

}
