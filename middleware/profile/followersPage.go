package profile

import (
	"fmt"
	"net/http"

	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

//FollowersPage ... Shows the page for followers
func FollowersPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	userstruct, err := get.UserByID(client.Eclient, r.URL.Path[11:])
	if err != nil {
		fmt.Println("err middleware/profile/followerspage: line 20")
		fmt.Println(err)
	}

	heads := []types.FloatingHead{}

	for index, followerID := range userstruct.Followers {
		head, err := uses.ConvertUserToFloatingHead(client.Eclient, followerID)
		if err != nil {
			fmt.Println(fmt.Sprintf("err middleware/profile/followerspage: line 31, index %d", index))
			fmt.Println(err)
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

	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string), ListOfHeads: heads}

	client.RenderTemplate(w, "template2-nil", cs)
	client.RenderTemplate(w, "leftnav-nil", cs)
	client.RenderTemplate(w, "followerlist-nil", cs)
}
