package project

import (
	"fmt"
	"net/http"

	get "github.com/sea350/ustart_go/get/project"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/project"
	"github.com/sea350/ustart_go/types"
)

//AddQuickLink ...
func AddQuickLink(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	ID := r.FormValue("projectID")
	fmt.Println("debug text. middleware/project/addquicklink. projectID is:")
	fmt.Println(r.FormValue("projectID"))
	proj, err := get.ProjectByID(client.Eclient, ID)
	if err != nil {
		fmt.Println(err)
		fmt.Println("this is an err: middleware/project/addQuickLink line 25")
	}
	fmt.Println("debug text. middleware/project/addquicklink. line 29")
	proj.QuickLinks = append(proj.QuickLinks, types.Link{Name: r.FormValue("projectLinkDesc"), URL: r.FormValue("projectLink")})

	err = post.UpdateProject(client.Eclient, ID, "QuickLinks", proj.QuickLinks)
	if err != nil {
		fmt.Println(err)
		fmt.Println("this is an err: middleware/project/addQuickLink line 33")
	}

}
