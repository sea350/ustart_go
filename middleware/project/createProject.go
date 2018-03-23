package project

import (
	"fmt"
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//CreateProject ... takes in a form for a new project and the id from session
//creates a new project and updates associated user arrays
func CreateProject(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	r.ParseForm()
	title := r.FormValue("project_title")
	description := []rune(r.FormValue("project_desc"))
	category := r.FormValue("category")
	college := r.FormValue("universityName")
	customURL := r.FormValue("curl")

	id, err := uses.CreateProject(client.Eclient, title, description, session.Values["DocID"].(string), category, college, customURL)
	if err != nil {
		fmt.Println("This is an error middleware/project/createproject")
		fmt.Println(err)
	}

	http.Redirect(w, r, "/Projects/"+id, http.StatusFound)

}
