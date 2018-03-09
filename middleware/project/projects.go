package project

import (
	"net/http"

	"fmt"

	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"

	client "github.com/sea350/ustart_go/middleware/client"
)

//ProjectsPage ... Displays the projects page
func ProjectsPage(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	project, err := uses.AggregateProjectData(client.Eclient, r.URL.Path[10:])
	if err != nil {
		panic(err)
	}
	//WIP
	//numFollowers:=len(project.ProjectData.Followers)

	widgets, errs := uses.LoadWidgets(client.Eclient, project.ProjectData.Widgets)
	if len(errs) > 0 {
		fmt.Println("there were one or more errors loading widgets")
		for _, eror := range errs {
			fmt.Println(eror)
		}
	}
	userstruct, err := get.UserByID(client.Eclient, session.Values["DocID"].(string))
	if err != nil {
		panic(err)
	}
	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string), Project: project, Widgets: widgets}
	client.RenderTemplate(w, "template2-nil", cs)
	client.RenderTemplate(w, "projectsF", cs)
}

//MyProjects ...
func MyProjects(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
	}

	var heads []types.FloatingHead

	userstruct, err := get.UserByID(client.Eclient, session.Values["DocID"].(string))
	if err != nil {
		panic(err)
	}
	for _, projectInfo := range userstruct.Projects {
		head, err := uses.ConvertProjectToFloatingHead(client.Eclient, projectInfo.ProjectID)
		if err != nil {
			panic(err)
		}
		heads = append(heads, head)
	}
	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string), ListOfHeads: heads}
	client.RenderTemplate(w, "template2-nil", cs)
	client.RenderTemplate(w, "manageprojects-Nil", cs)
}

//CreateProjectPage ...
func CreateProjectPage(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	userstruct, err := get.UserByID(client.Eclient, session.Values["DocID"].(string))
	if err != nil {
		panic(err)
	}
	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string)}
	client.RenderTemplate(w, "template2-nil", cs)
	client.RenderTemplate(w, "createProject-Nil", cs)
}
