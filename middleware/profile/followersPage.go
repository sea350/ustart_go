package profile

import (
	"fmt"
	"log"
	"net/http"
	"os"

	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

//FollowersPage ... Shows the page for followers
func FollowersPage(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	docID, err := get.IDByUsername(client.Eclient, r.URL.Path[11:])
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	// docID := session.Values["DocID"].(string)
	userstruct, err := get.UserByID(client.Eclient, docID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	heads := []types.FloatingHead{}

	for index, followerID := range userstruct.Followers {
		head, err := uses.ConvertUserToFloatingHead(client.Eclient, followerID)
		if err != nil {
			fmt.Println(fmt.Sprintf("err middleware/profile/followerspage: line 36, index %d", index))
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
			continue
		}
		for _, element := range userstruct.Following {
			if element == followerID {
				head.Followed = true
				break
			}
		}
		heads = append(heads, head)
	}

	heads2 := []types.FloatingHead{}
	for index, following := range userstruct.Following {
		head, err := uses.ConvertUserToFloatingHead(client.Eclient, following)
		if err != nil {
			fmt.Println(fmt.Sprintf("err middleware/profile/followerspage: line 36, index %d", index))
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
			continue
		}
		heads2 = append(heads2, head)
	}

	cs := client.ClientSide{UserInfo: userstruct, Page: docID, ListOfHeads: heads, ListOfHeads2: heads2}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "followerlist-nil", cs)
}
