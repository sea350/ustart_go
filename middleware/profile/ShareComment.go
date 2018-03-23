package profile

import (
	"fmt"
	"net/http"

	"github.com/sea350/ustart_go/middleware/stringHTML"
	uses "github.com/sea350/ustart_go/uses"
)

//ShareComments ...  purpose unknown
func ShareComments(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	r.ParseForm()
	postid := r.FormValue("PostID")
	//postaid := postid[9:]
	postactual := postid[11:]
	pika := r.FormValue("Pikachu")
	fmt.Println("This is debug text, ShareComment.go: 21")
	fmt.Println(pika)
	parentPost, arrayofComments, err4 := uses.LoadComments(eclient, postactual, 0, -1)
	if err4 != nil {
		fmt.Println(err4)
	}
	var sum int
	var output string
	var commentoutputs string

	for i := 0; i < len(arrayofComments); i++ {

		commentoutputs += stringHTML.CommentEntry(arrayofComments[i].Image, arrayofComments[i].FirstName, arrayofComments[i].LastName, string(arrayofComments[i].Element.Content),postactual)
		sum += i
	}
	username := session.Values["Username"].(string)

	output += stringHTML.OutputShare(postactual, parentPost.Image, parentPost.FirstName, parentPost.LastName, string(parentPost.Element.Content), pika, username)

	fmt.Fprintln(w, output)

}

/* This function might not be used anymore. */

//ShareComment2 ... pupose unknown
func ShareComment2(w http.ResponseWriter, r *http.Request) {
	// If followingStatus = no
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
	}

	r.ParseForm()
	docid := r.FormValue("id")
	postid := r.FormValue("postid")
	msg := r.FormValue("msg")
	username := r.FormValue("username")
	content := []rune(msg)

	err := uses.UserShareEntry(eclient, docid, postid, content)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/profile/"+username, http.StatusFound)

}
