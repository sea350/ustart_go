package profile

import (
	"fmt"
	"net/http"

	// "github.com/sea350/ustart_go/middleware/stringHTML"

)


func DeleteWallPost(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
	}

	//err := uses.RemoveEntry(eclient,session.Values["DocID"].(string),)

}

func GenerateDeleteModal(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	// IF I change output to <p> LOOOOOOL </p>, it works. problem finding modal shit 
	output := `                              <div class="modal fade" id="confirm-delete" role="dialog" >
                                <div class="modal-dialog">
                                    <div class="modal-content">
                                        <div class="modal-header">
                                            <span style="font-size:20px;">Confirm Deletion</span>
                                        </div>
                                        <div class="modal-body">
                                            <span style="font-size:15px;">Are you sure you want to delete this post?</span>
                                        </div>
<!--                                         <div class="modal-footer">
                                            <button type="button" class="btn btn-default" data-dismiss="modal">Cancel</button>
                                            <form action ="http://ustart.today:5000/deletePost" method ="POST"> 
                                                 <input type="text" name="postid" value=(this).attr("id")>  
                                             <button type="submit"><a class="btn btn-danger btn-ok">Delete</a></button>
                                        </form>
                                        </div> -->
                                    </div>
                                </div>
                            </div>`
fmt.Fprintln(w, output)
}