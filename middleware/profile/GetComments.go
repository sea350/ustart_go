package profile

import (
	"fmt"
	"net/http"

	"github.com/sea350/ustart_go/middleware/stringHTML"
	uses "github.com/sea350/ustart_go/uses"
)

//GetComments ... gets comments???
func GetComments(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		//No docid in session
		http.Redirect(w, r, "/~", http.StatusFound)
	}

	r.ParseForm()
	postID := r.FormValue("PostID")
	postaid := postID[9:]
	postactual := postID[10:]
	// need to trim beginning of postID
	pika := r.FormValue("Pikachu")
	fmt.Println("This is debug text, GetComments.go: 23")
	fmt.Println(pika)
	// journal entry, err
	parentPost, arrayofComments, err4 := uses.LoadComments(eclient, postactual, 0, -1)
		fmt.Println(parentPost);

	fmt.Println("ARRAY OF COMMENTS");
	fmt.Println(arrayofComments);
	if err4 != nil {
		fmt.Println("This is debug text, GetComments.go: 29")
		fmt.Println(err4)
	}

	var sum int
	var commentoutputs string
	/*
		The following is how AJAX for loading comments is handled on the server side.
	*/
	for i := 0; i < len(arrayofComments); i++ {
		commentoutputs += stringHTML.CommentEntry(arrayofComments[i].Image, arrayofComments[i].FirstName, arrayofComments[i].LastName, string(arrayofComments[i].Element.Content))
		fmt.Println(arrayofComments[i].Element.Content);
		sum += i
	}
	username := session.Values["Username"].(string)

	output := stringHTML.ParentEntry(postaid, parentPost.Image, parentPost.FirstName, parentPost.LastName, string(parentPost.Element.Content), pika, username, commentoutputs)

	fmt.Fprintln(w, output)

}
