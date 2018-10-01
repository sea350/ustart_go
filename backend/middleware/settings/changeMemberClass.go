package settings

import (
	"log"
	"net/http"
	"os"

	get "github.com/sea350/ustart_go/backend/get/project"
	client "github.com/sea350/ustart_go/backend/middleware/client"
	post "github.com/sea350/ustart_go/backend/post/project"
	uses "github.com/sea350/ustart_go/backend/uses"
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
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	var isCreator, _ = uses.IsLeader(client.Eclient, projectID, test1.(string))

	if isCreator {
		for i, member := range project.Members {
			if member.MemberID == test1.(string) && member.Role <= 0 {
				isCreator = true
			}

			if member.MemberID == memberID {
				switch newRank {
				case "Moderator":
					project.Members[i].Role = 1
					project.Members[i].Title = "Admin"

				default:
					project.Members[i].Role = 2
					project.Members[i].Title = "Member"
				}

				if member.Role != 0 && newRank != "Creator" {
					err = post.UpdateProject(client.Eclient, projectID, "Members", project.Members)
				} else {
					log.Println("You do not have permission to change member class of this project")
				}
			}

			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				dir, _ := os.Getwd()
				log.Println(dir, err)

			}
		}
	}
}
