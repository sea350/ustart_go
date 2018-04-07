package settings

import (
	"fmt"
	"net/http"
	"strconv"

	get "github.com/sea350/ustart_go/get/project"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/project"
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
	newRank := r.FormValue("newRank")

	project, err := get.ProjectByID(client.Eclient, projectID)
	if err != nil {
		fmt.Println("error: middleware/project/changememberclass line 25")
		fmt.Println(err)
	}

	var isCreator = false
	for i, member := range project.Members {
		if member.MemberID == test1.(string) && member.Role <= 0 {
			isCreator = true
		}

		if member.MemberID == memberID {
			rankInt, err := strconv.Atoi(newRank)
			if err != nil {
				fmt.Println("error: middleware/project/changememberclass line 38")
				fmt.Println(err)
			} else if member.Role != 0 && rankInt != 0 {
				project.Members[i].Role = rankInt
			}
		}
	}
	if isCreator {
		err = post.UpdateProject(client.Eclient, projectID, "Members", project.Members)
		if err != nil {
			fmt.Println("error: middleware/project/changememberclass line 49")
			fmt.Println(err)
		}
	}

}
