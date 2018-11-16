package project

import (
	"log"
	"net/http"
	"os"

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
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	var heads []types.FloatingHead

	userstruct, err := get.UserByID(client.Eclient, session.Values["DocID"].(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	for _, projectInfo := range userstruct.Projects {
		var isAdmin = false
		proj, err := getProj.ProjectByID(client.Eclient, projectInfo.ProjectID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}

		for _, memberInfo := range proj.Members {
			if memberInfo.MemberID == test1.(string) && memberInfo.Role <= 1 {
				//finds user in the list of members and also checks if they have creator rank
				isAdmin = true
				//head.Followed in this case expresses whether or not they have edit permissions
			}
		}
		if !isAdmin {
			continue
		}
		head, err := uses.ConvertProjectToFloatingHead(client.Eclient, projectInfo.ProjectID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
		heads = append(heads, head)
	}

	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string), ListOfHeads: heads}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "ManageProjectMembersF", cs)
}
