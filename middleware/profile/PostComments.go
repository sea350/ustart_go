package profile

import (
	"fmt"
	"net/http"

	uses "github.com/sea350/ustart_go/uses"
)

//GetComments ... gets comments???
func PostComments(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		//No docid in session
		http.Redirect(w, r, "/~", http.StatusFound)
	}

	r.ParseForm()
	postID := r.FormValue("PostID")
	// postaid := postID[9:]
	// postactual := postID[10:]
	fmt.Println(postID + " IS THE POST 3 23 2018")
	// need to trim beginning of postID
	//	pika := r.FormValue("Pikachu")
	// fmt.Println("This is debug text, GetComments.go: 23")
	// fmt.Println(pika) // pika is your own doc id
	// journal entry, err
	_, arrayofComments, err4 := uses.LoadComments(eclient, postID, 0, -1)
	//	fmt.Println(parentPost);

	fmt.Println("ARRAY OF COMMENTS")
	//	fmt.Println(arrayofComments);
	if err4 != nil {
		fmt.Println("This is debug text, POSTComments.go: 29")
		fmt.Println(err4)
	}

	var sum int
	var commentoutputs string
	/*
		The following is how AJAX for loading comments is handled on the server side.
	*/
	for i := 0; i < len(arrayofComments); i++ {
		commentoutputs += `					<div class="media">
		 			<a class="media-left" href="#">
						<img class="media-object img-rounded" src="https://scontent-lga3-1.xx.fbcdn.net/v/t31.0-8/12514060_499384470233859_6798591731419500290_o.jpg?oh=329ea2ff03ab981dad7b19d9172152b7&oe=5A2D7F0D">
					</a>
					<div class="media-body">
						<h5 class="media-heading user_name ` + postID + `" style="color:cadetblue;">` + arrayofComments[i].FirstName + `</h5>
						<p> ` + string(arrayofComments[i].Element.Content) + `</p>
					</div>
				</div>`
		sum += i
	}
	// username := session.Values["Username"].(string)
	output := commentoutputs
	// output := stringHTML.ParentEntry(postaid, parentPost.Image, parentPost.FirstName, parentPost.LastName, string(parentPost.Element.Content), pika, username, commentoutputs)

	fmt.Fprintln(w, output)

}
