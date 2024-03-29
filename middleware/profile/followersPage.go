package profile

import (
	"net/http"

	getFollow "github.com/sea350/ustart_go/get/follow"
	get "github.com/sea350/ustart_go/get/user"
	getUser "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

//FollowersPage ... Shows the page for followers
func FollowersPage(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if r.URL.Path[11:] == `_blank` {
		return
	}

	id, err := getUser.IDByUsername(client.Eclient, r.URL.Path[11:])
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
	_, followDoc, err := getFollow.ByID(client.Eclient, id)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	userstruct, err := get.UserByID(client.Eclient, id)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	heads := []types.FloatingHead{}

	//_ for bell follows
	for idKey := range followDoc.UserFollowers {
		head, err := uses.ConvertUserToFloatingHead(client.Eclient, idKey)
		if err != nil {

			client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + idKey)
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			continue
		}
		_, isFollowing := followDoc.UserFollowing[idKey]
		head.Followed = isFollowing
		heads = append(heads, head)
	}

	for idKey := range followDoc.ProjectFollowers {
		head, err := uses.ConvertProjectToFloatingHead(client.Eclient, idKey)
		if err != nil {

			client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + idKey)
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			continue
		}
		isFollowing, _ := followDoc.ProjectFollowing[idKey]
		head.Followed = isFollowing
		heads = append(heads, head)
	}

	heads2 := []types.FloatingHead{}
	for idKey := range followDoc.UserFollowing {
		head, err := uses.ConvertUserToFloatingHead(client.Eclient, idKey)
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + idKey)
			continue
		}

		head.Followed, err = getFollow.IsFollowing(client.Eclient, test1.(string), idKey, "user")
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + idKey)
			continue
		}

		heads2 = append(heads2, head)
	}

	heads3 := []types.FloatingHead{}
	for idKey := range followDoc.ProjectFollowing {
		projHead, err := uses.ConvertProjectToFloatingHead(client.Eclient, idKey)
		if err != nil {
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + idKey)
			continue
		}

		projHead.Followed, err = getFollow.IsFollowing(client.Eclient, test1.(string), idKey, "project")
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + idKey)
			continue
		}
		heads3 = append(heads3, projHead)
	}

	isFollowing, err := getFollow.IsFollowing(client.Eclient, test1.(string), id, "user")

	numberFollowers := len(followDoc.UserFollowers) + len(followDoc.ProjectFollowers) + len(followDoc.EventFollowers)
	userFoll := len(followDoc.UserFollowing)
	projFoll := len(followDoc.ProjectFollowing)
	eventFoll := len(followDoc.EventFollowing)
	cs := client.ClientSide{UserInfo: userstruct, Page: test1.(string), Followers: numberFollowers, FollowingStatus: isFollowing, UserFollowing: userFoll, ProjFollowing: projFoll, EventFollowing: eventFoll, ListOfHeads: heads, ListOfHeads2: heads2, ListOfHeads3: heads3}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "followerlist-nil", cs)
}
