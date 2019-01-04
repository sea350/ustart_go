package project

import (
	"errors"

	"net/http"
	"time"

	"github.com/microcosm-cc/bluemonday"
	getFollow "github.com/sea350/ustart_go/get/follow"
	get "github.com/sea350/ustart_go/get/user"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"

	client "github.com/sea350/ustart_go/middleware/client"
)

//ProjectsPage ... Displays the projects page
func ProjectsPage(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values[`DocID`]
	if docID == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	var cs client.ClientSide

	url := r.URL.Path[10:]
	if url == `` {

		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + `NO URL PASSED`)
	}
	if url == `_blank` {
		return
	}

	project, err := uses.AggregateProjectData(client.Eclient, r.URL.Path[10:], docID.(string))
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		cs.ErrorStatus = true
		cs.ErrorOutput = err
		client.RenderSidebar(w, r, "template2-nil")
		client.RenderSidebar(w, r, "leftnav-nil")
		client.RenderTemplate(w, r, "projectsF", cs)
		return
	}
	if !project.ProjectData.Visible {
		http.Redirect(w, r, "/404/", http.StatusFound)
		return
	}

	widgets, errs := uses.LoadWidgets(client.Eclient, project.ProjectData.Widgets)
	if len(errs) > 0 {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "there were one or more errors loading widgets")
		for _, eror := range errs {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", eror)
		}
	}
	userstruct, err := get.UserByID(client.Eclient, docID.(string))
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		cs.ErrorStatus = true
		cs.ErrorOutput = err
	}

	_, follDoc, err := getFollow.ByID(client.Eclient, project.DocID)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
	_, followingState := follDoc.UserFollowers[docID.(string)]

	numberFollowers := len(follDoc.UserFollowers)
	// numberFollowing := len(follDoc.UserFollowing) + len(follDoc.ProjectFollowing) + len(follDoc.EventFollowing)
	cs = client.ClientSide{UserInfo: userstruct, DOCID: docID.(string), Username: session.Values["Username"].(string), Followers: int(numberFollowers), FollowingStatus: followingState, Project: project, Widgets: widgets}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "projectsF", cs)
}

//MyProjects ... ManageProject
func MyProjects(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values[`DocID`]
	if docID == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	var heads []types.FloatingHead
	var cs client.ClientSide
	userstruct, err := get.UserByID(client.Eclient, docID.(string))
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		cs.ErrorStatus = true
		cs.ErrorOutput = err
		client.RenderSidebar(w, r, "template2-nil")
		client.RenderSidebar(w, r, "leftnav-nil")
		client.RenderTemplate(w, r, "manageprojects-Nil", cs)
		return
	}
	for _, projectInfo := range userstruct.Projects {
		head, err := uses.ConvertProjectToFloatingHead(client.Eclient, projectInfo.ProjectID)
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
		heads = append(heads, head)
	}
	cs = client.ClientSide{UserInfo: userstruct, DOCID: docID.(string), Username: session.Values["Username"].(string), ListOfHeads: heads}
	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "manageprojects-Nil", cs)
}

//CreateProjectPage ...
func CreateProjectPage(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values[`DocID`]
	if docID == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	userstruct, err := get.UserByID(client.Eclient, docID.(string))
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
	cs := client.ClientSide{UserInfo: userstruct, DOCID: docID.(string), Username: session.Values["Username"].(string)}

	p := bluemonday.UGCPolicy()
	title := r.FormValue("project_title")
	cleanTitle := p.Sanitize(title)

	description := []rune(r.FormValue("project_desc"))
	cleanDesc := p.Sanitize(string(description))

	cleanCat := p.Sanitize(r.FormValue("category"))
	if len(cleanCat) == 0 {
		//		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"Cannot leave category blank")
		client.RenderSidebar(w, r, "template2-nil")
		client.RenderSidebar(w, r, "leftnav-nil")
		client.RenderTemplate(w, r, "createProject-Nil", cs)
		return
	}

	cleanCollege := p.Sanitize(r.FormValue("universityName"))
	if len(cleanCollege) == 0 {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Cannot leave college blank")
		client.RenderSidebar(w, r, "template2-nil")
		client.RenderSidebar(w, r, "leftnav-nil")
		client.RenderTemplate(w, r, "createProject-Nil", cs)
		return
	}

	cleanURL := p.Sanitize(r.FormValue("curl"))
	if len(cleanURL) == 0 {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Cannot leave custom URL blank")
		client.RenderSidebar(w, r, "template2-nil")
		client.RenderSidebar(w, r, "leftnav-nil")
		client.RenderTemplate(w, r, "createProject-Nil", cs)
		return
	}

	cleanCountry := p.Sanitize(r.FormValue("country"))

	cleanState := p.Sanitize(r.FormValue("state"))

	cleanCity := p.Sanitize(r.FormValue("city"))

	cleanZip := p.Sanitize(r.FormValue("zip"))

	cleanStreet := p.Sanitize(r.FormValue("street"))

	var projLocation types.LocStruct
	projLocation.Street = cleanStreet
	projLocation.City = cleanCity
	projLocation.Country = cleanCountry
	projLocation.Zip = cleanZip
	projLocation.State = cleanState
	projLocation.Street = cleanStreet

	if title != `` {
		//proper URL
		if !uses.ValidUsername(cleanURL) {

			client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Invalid custom project URL")
			cs.ErrorStatus = true
			cs.ErrorOutput = errors.New("Invalid custom project URL")
			client.RenderSidebar(w, r, "template2-nil")
			client.RenderSidebar(w, r, "leftnav-nil")
			client.RenderTemplate(w, r, "createProject-Nil", cs)
			return

		}
		url, err := uses.CreateProject(client.Eclient, cleanTitle, []rune(cleanDesc), docID.(string), cleanCat, cleanCollege, cleanURL, projLocation)
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			cs.ErrorStatus = true
			cs.ErrorOutput = err
		} else {
			time.Sleep(2 * time.Second)
			http.Redirect(w, r, "/Projects/"+url, http.StatusFound)
			return
		}
	}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "createProject-Nil", cs)
}
