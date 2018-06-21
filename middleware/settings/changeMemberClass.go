package settings

import (
	"fmt"
	"net/http"

	get "github.com/sea350/ustart_go/get/project"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/project"
	uses "github.com/sea350/ustart_go/uses"
)

//ChangeMemberClass ...
func ChangeMemberClass(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	memberID := r.FormValue("memberID")
	projectID := r.FormValue("projectID")
	newRole := r.FormValue("newRole")

	// var roleName string
	var roleInt int
	switch newRole {
	case "Member":
		// roleName = "Member"
		roleInt = 2
	case "Moderator":
		// roleName = "Admin"
		roleInt = 1
	}

	project, err := get.ProjectByID(client.Eclient, projectID)
	if err != nil {
		fmt.Println("error: middleware/project/changememberclass line 25")
		fmt.Println(err)
	}

	var isCreator, _ = uses.IsLeader(client.Eclient, projectID, test1.(string))
	fmt.Println("IS CREATOR:", isCreator)
	if isCreator {
		fmt.Println("IS CREATOR")
		for i, member := range project.Members {
			// if member.MemberID == test1.(string) && member.Role <= 0 {
			// 	isCreator = true
			// }

			if member.MemberID == memberID {
				fmt.Println(member.Role)
				if err != nil {
					fmt.Println("error: middleware/project/changememberclass line 38")
					fmt.Println(err)
				} else if member.Role != 0 {

					project.Members[i].Role = roleInt

					err = post.UpdateProject(client.Eclient, projectID, "Members", project.Members)
					if err != nil {
						fmt.Println("error: middleware/project/changememberclass line 49")
						fmt.Println(err)
					}
				}
			}
		}
		// if isCreator {
		// 	err = post.UpdateProject(client.Eclient, projectID, "Members", project.Members)
		// 	if err != nil {
		// 		fmt.Println("error: middleware/project/changememberclass line 49")
		// 		fmt.Println(err)
		// 	}
		//}
	}
}
