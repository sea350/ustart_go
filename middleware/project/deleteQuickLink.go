package project

import (
	"fmt"
	"net/http"

	get "github.com/sea350/ustart_go/get/project"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/project"
	"github.com/sea350/ustart_go/types"
)

//DeleteQuickLink ...
func DeleteQuickLink(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
	}

	ID := r.FormValue("projectID")

	proj, err := get.ProjectByID(client.Eclient, ID)
	if err != nil {
		fmt.Println(err)
		fmt.Println("this is an err: middleware/project/deleteQuickLink line 25")
	}

	deleteTitle := r.FormValue("deleteProjectLinkDesc")
	deleteURL := r.FormValue("deleteProjectLink")

	var newArr []types.Link

	if len(proj.QuickLinks) == 1 {
		err := post.UpdateProject(client.Eclient, ID, "QuickLinks", newArr)
		if err != nil {
			fmt.Println(err)
			fmt.Println("this is an err: middleware/project/deleteQuickLink line 39")
		}
		http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
		return
	}

	target := -1
	for index, link := range proj.QuickLinks {

		if link.Name == deleteTitle && link.URL == deleteURL {
			target = index
			fmt.Println(target)
			break
		}
	}

	if target == -1 {
		fmt.Println("deleted object not found")
		fmt.Println("this is an err, middleware/profile/deleteQuickLink line 57")
		newArr = proj.QuickLinks
	} else if (target + 1) < len(proj.QuickLinks) {
		newArr = append(proj.QuickLinks[:target], proj.QuickLinks[(target+1):]...)
	} else {
		newArr = proj.QuickLinks[:target]
	}

	err = post.UpdateProject(client.Eclient, ID, "QuickLinks", newArr)
	if err != nil {
		fmt.Println(err)
		fmt.Println("this is an err: middleware/profile/deleteQuickLink line 68")
	}

	http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
}
