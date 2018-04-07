package project

import (
	"fmt"
	"net/http"
	"time"

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
	project, err := uses.AggregateProjectData(client.Eclient, r.URL.Path[10:], test1.(string))
	if err != nil {
		fmt.Println(err)
		fmt.Println("this is an err: middleware/project/projects.go line 23")
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
	client.RenderTemplate(w, r, "template2-nil", cs)
	client.RenderTemplate(w, r, "leftnav-nil", cs)
	client.RenderTemplate(w, r, "projectsF", cs)
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
	client.RenderTemplate(w, r, "template2-nil", cs)
	client.RenderTemplate(w, r, "leftnav-nil", cs)
	client.RenderTemplate(w, r, "manageprojects-Nil", cs)
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

	title := r.FormValue("project_title")
	description := []rune(r.FormValue("project_desc"))
	category := r.FormValue("category")
	college := r.FormValue("universityName")
	customURL := r.FormValue("curl")

	if title != `` {
		url, err := uses.CreateProject(client.Eclient, title, description, session.Values["DocID"].(string), category, college, customURL)
		if err != nil {
			fmt.Println("This is an error middleware/project/createproject")
			fmt.Println(err)
			cs.ErrorStatus = true
			cs.ErrorOutput = err
		} else {
			time.Sleep(5000)
			http.Redirect(w, r, "/Projects/"+url, http.StatusFound)
			return
		}
	}

	client.RenderTemplate(w, r, "template2-nil", cs)
	client.RenderTemplate(w, r, "leftnav-nil", cs)
	client.RenderTemplate(w, r, "createProject-Nil", cs)
}