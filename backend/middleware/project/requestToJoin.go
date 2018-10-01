package project

import (
	"log"
	"net/http"

	get "github.com/sea350/ustart_go/backend/get/project"
	client "github.com/sea350/ustart_go/backend/middleware/client"
	projPost "github.com/sea350/ustart_go/backend/post/project"
	userPost "github.com/sea350/ustart_go/backend/post/user"
)

//RequestToJoin ...
func RequestToJoin(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	ID := r.FormValue("projID") //project docID

	proj, err := get.ProjectByID(client.Eclient, ID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
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
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	err = projPost.AppendMemberReqReceived(client.Eclient, ID, test1.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
	return
}
