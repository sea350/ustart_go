package project

import (
	"fmt"
	"net/http"

	get "github.com/sea350/ustart_go/get/project"
	client "github.com/sea350/ustart_go/middleware/client"
	projPost "github.com/sea350/ustart_go/post/project"
	userPost "github.com/sea350/ustart_go/post/user"
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
	fmt.Println(ID)
	fmt.Println("debug text requesttojoin line 23")

	proj, err := get.ProjectByID(client.Eclient, ID)
	if err != nil {
		fmt.Println("err middleware/project/requesttojoin line25")
		fmt.Println(err)
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
		fmt.Println("err middleware/project/requesttojoin line42")
		fmt.Println(err)
	}
	err = projPost.AppendMemberReqReceived(client.Eclient, ID, test1.(string))
	if err != nil {
		fmt.Println("err middleware/project/requesttojoin line47")
		fmt.Println(err)
	}

	http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
}
