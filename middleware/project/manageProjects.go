package project

import (
	"fmt"
	"net/http"

	getProj "github.com/sea350/ustart_go/get/project"
	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/uses"

	"github.com/sea350/ustart_go/middleware/client"
	types "github.com/sea350/ustart_go/types"
)

//ManageProjects ...
func ManageProjects(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
	}

	var heads []types.FloatingHead

	userstruct, err := get.UserByID(client.Eclient, session.Values["DocID"].(string))
	if err != nil {
		fmt.Println(err)
		fmt.Println("err: middleware/project/manageprojects Line 26")
	}
	for _, projectInfo := range userstruct.Projects {
		head, err := uses.ConvertProjectToFloatingHead(client.Eclient, projectInfo.ProjectID)
		if err != nil {
			fmt.Println(err)
			fmt.Println("err: middleware/project/manageprojects Line 32")
		}

		proj, err := getProj.ProjectByID(client.Eclient, projectInfo.ProjectID)
		if err != nil {
			fmt.Println(err)
			fmt.Println("err: middleware/project/manageprojects Line 39")
		}

		for _, memberInfo := range proj.Members {
			if memberInfo.MemberID == test1.(string) && memberInfo.Role == 0 {
				//finds user in the list of members and also checks if they have creator rank
				head.Followed = true
				//head.Followed in this case expresses whether or not they have edit permissions
			}
		}
		if !head.Followed {
			continue
		}
		heads = append(heads, head)
	}
	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string), ListOfHeads: heads}
	client.RenderTemplate(w, "template2-nil", cs)
	//client.RenderTemplate(w, "manageprojects-Nil", cs)
}