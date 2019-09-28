package project

import (
	"net/http"
	"sync"

	get "github.com/sea350/ustart_go/get/project"
	client "github.com/sea350/ustart_go/middleware/client"
	projPost "github.com/sea350/ustart_go/post/project"
	userPost "github.com/sea350/ustart_go/post/user"
)

var sktchyLck sync.Mutex

//RequestToJoin ...
func RequestToJoin(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	ID := r.FormValue("projID") //project docID

	sktchyLck.Lock()
	defer sktchyLck.Unlock()

	proj, err := get.ProjectByID(client.Eclient, ID)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	for _, memberInfo := range proj.Members {
		if memberInfo.MemberID == test1.(string) {
			http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
			return
		}
	}
	for _, receivedReq := range proj.MemberReqReceived {
		if receivedReq == test1.(string) {
			http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
			return
		}
	}
	err = userPost.AppendSentProjReq(client.Eclient, test1.(string), ID)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
	err = projPost.AppendMemberReqReceived(client.Eclient, ID, test1.(string))
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
	return
}
