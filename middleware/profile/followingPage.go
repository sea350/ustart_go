package profile

import (
	"log"
	"net/http"

	getFollow "github.com/sea350/ustart_go/get/follow"
	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

//FollowingPage ... Shows the page for people the user is following
func FollowingPage(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	_, followDoc, err := getFollow.ByID(client.Eclient, test1.(string))
	userstruct, err := get.UserByID(client.Eclient, test1.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	heads := []types.FloatingHead{}

	for idKey := range followDoc.UserFollowing {
		head, err := uses.ConvertUserToFloatingHead(client.Eclient, idKey)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			log.Println(idKey)
			continue
		}
		heads = append(heads, head)
	}

	for idKey := range followDoc.ProjectFollowing {
		head, err := uses.ConvertProjectToFloatingHead(client.Eclient, idKey)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			log.Println(idKey)
			continue
		}
		heads = append(heads, head)
	}

	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string), ListOfHeads: heads}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "followerlist-nil", cs)
}
