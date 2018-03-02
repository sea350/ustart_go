package project

import (
	"fmt"
	"net/http"

	uses "github.com/sea350/ustart_go/uses"
)

//CreateProject ... takes in a form for a new project and the id from session
//creates a new project and updates associated user arrays
func CreateProject(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	r.ParseForm()
	title := r.FormValue("UNKOWN")
	description := []rune(r.FormValue("UNKNOWN"))

	err := uses.CreateProject(eclient, title, description, session.Values["DocID"].(string))
	if err != nil {
		fmt.Println("This is an error middleware/project/createproject")
		fmt.Println(err)
	}

	http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)

}
