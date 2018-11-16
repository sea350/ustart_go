package project

import (
	"log"
	"net/http"

	getFollow "github.com/sea350/ustart_go/get/follow"
	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

//AjaxLoadProjectFollowers ... Shows the page for followers
func AjaxLoadProjectFollowers(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	_, followDoc, err := getFollow.ByID(client.Eclient, r.URL.Path[10:])
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	project, err := uses.AggregateProjectData(client.Eclient, r.URL.Path[10:], test1.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	userstruct, err := get.UserByID(client.Eclient, test1.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	heads := []types.FloatingHead{}

	//_ for bell follows
	for idKey := range followDoc.UserFollowers {
		head, err := uses.ConvertUserToFloatingHead(client.Eclient, idKey)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			log.Println("index " + idKey)
			continue
		}
		isFollowing, _ := followDoc.UserFollowing[idKey]
		head.Followed = isFollowing
		heads = append(heads, head)
	}

	for idKey := range followDoc.ProjectFollowers {
		head, err := uses.ConvertProjectToFloatingHead(client.Eclient, idKey)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			log.Println("index " + idKey)
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
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println("index " + idKey)
			log.Println(err)
			continue
		}
		heads2 = append(heads2, head)
	}

	for idKey := range followDoc.ProjectFollowing {
		head, err := uses.ConvertProjectToFloatingHead(client.Eclient, idKey)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println("index " + idKey)
			log.Println(err)
			continue
		}
		heads2 = append(heads2, head)
	}

	cs := client.ClientSide{UserInfo: userstruct, Page: test1.(string), ListOfHeads: heads, Project: project, ListOfHeads2: heads2}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "followerlist-nil", cs)
}
