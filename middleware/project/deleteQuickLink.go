package project

import (
	
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
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	ID := r.FormValue("projectID")
	if ID == `` {
		
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"Crucial data was not passed in, now exiting")
		return
	}

	proj, err := get.ProjectByID(client.Eclient, ID)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	deleteTitle := r.FormValue("deleteProjectLinkDesc")
	deleteURL := r.FormValue("deleteProjectLink")

	// if deleteTitle == `` {
	// 	
	// 			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"WARNING: link title is blank")
	// }
	if deleteURL == `` {
		
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"Crucial data was not passed in, now exiting")
		return
	}

	var newArr []types.Link

	if len(proj.QuickLinks) <= 1 {
		err := post.UpdateProject(client.Eclient, ID, "QuickLinks", newArr)
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
		return
	}

	target := -1
	for index, link := range proj.QuickLinks {

		if link.Name == deleteTitle && link.URL == deleteURL {
			target = index
			break
		}
	}

	if target == -1 {
		
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"Deleted object not found")
		return
	} else if (target + 1) < len(proj.QuickLinks) {
		newArr = append(proj.QuickLinks[:target], proj.QuickLinks[(target+1):]...)
	} else {
		newArr = proj.QuickLinks[:target]
	}

	err = post.UpdateProject(client.Eclient, ID, "QuickLinks", newArr)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	return
}
