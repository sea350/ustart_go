package profile

import (
	"fmt"
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
	// "github.com/sea350/ustart_go/middleware/stringHTML"
)

//DeleteWallPost ... Iunno
func DeleteWallPost(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	postid := r.FormValue("postid")

	fmt.Println(postid + "IS THE DELETE")
	err := uses.HideEntry(client.Eclient, postid)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)
	return

}

//GenerateDeleteModal ...
func GenerateDeleteModal(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	postid := r.FormValue("PostID")
	//postaid := postid[9:]
	postactual := postid[14:]
	// IF I change output to <p> LOOOOOOL </p>, it works. problem finding modal shit
	output := `                              

							<p> postactual nani` + postactual + ` </p>
							<div class="modal fade" id=confirm-delete` + postactual + `  role="dialog" >
                                <div class="modal-dialog">
                                    <div class="modal-content">
                                        <div class="modal-header">
                                            <span style="font-size:20px;">Confirm Deletion</span>
                                        </div>
                                        <div class="modal-body">
                                            <span style="font-size:15px;">Are you sure you want to delete this post?</span>
                                        </div>
                                        <div class="modal-footer">
                                            <button type="button" class="btn btn-default" data-dismiss="modal">Cancel</button>
                                            <form action ="http://ustart.today:5000/deletePost" method ="GET"> 
                                                 <input type="hidden" name="postid" value=` + postactual + `>  
                                             <button type="submit"><a class="btn btn-danger btn-ok">Delete</a></button>
                                        </form>
                                        </div> 
                                    </div>
                                </div>
                            </div>`
	fmt.Fprintln(w, output)
}
