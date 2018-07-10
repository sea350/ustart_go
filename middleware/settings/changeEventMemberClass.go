package settings

import (
	"log"
	"net/http"
	"os"

	get "github.com/sea350/ustart_go/get/event"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/event"
	uses "github.com/sea350/ustart_go/uses"
)

//ChangeEventMemberClass ...
func ChangeEventMemberClass(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	memberID := r.FormValue("memberID")
	eventID := r.FormValue("eventID")
	newRank := r.FormValue("newRank")

	event, err := get.EventByID(client.Eclient, eventID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	var isCreator, _ = uses.IsEventLeader(client.Eclient, eventID, test1.(string))

	if isCreator {
		for i, member := range event.Members {
			if member.MemberID == test1.(string) && member.Role <= 0 {
				isCreator = true
			}

			if member.MemberID == memberID {
				switch newRank {
				case "Moderator":
					event.Members[i].Role = 1
					event.Members[i].Title = "Admin"

				default:
					event.Members[i].Role = 2
					event.Members[i].Title = "Member"
				}

				if member.Role != 0 && newRank != "Creator" {
					err = post.UpdateEvent(client.Eclient, eventID, "Members", event.Members)
				} else {
					log.Println("You do not have permission to change member class of this event")
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
