package profile

import (
	"fmt"
	"net/http"

	uses "github.com/sea350/ustart_go/uses"
)

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
	parentPost, arrayofComments, err4 := uses.LoadComments(eclient, postactual, 0, -1)
	if err4 != nil {
		fmt.Println(err4)
	}
	var sum int
	var output string
	var commentoutputs string

	for i := 0; i < len(arrayofComments); i++ {
		commentBody := string(arrayofComments[i].Element.Content)
		commentoutputs += `    <div class="media">
                                                    <h6 class="pull-right text-muted time">2 hours ago</h6>
                                                    <a class="media-left" href="#">
                                                        <img style="height:40px;" class="img-rounded" src=d` + arrayofComments[i].Image + `>
                                                    </a>
                                                    <div class="media-body">
                                                        <h5 class="media-heading user_name" style="color:cadetblue;">` + arrayofComments[i].FirstName + " " + arrayofComments[i].LastName + `</h5>
                                                        <p>` + commentBody + `</p>
                                                        <p>
                                                            <small>
                                                                <a class="comment-like">Like</a>
                                                            </small> -
                                                            <small>
                                                                <a class="confirmation-callback">Remove</a>
                                                            </small>
                                                        </p>
                                                        <p>
                                                             <small>
                                                                <a class="view-replies" onclick="document.getElementById('replies').style.display = 'block'; this.style.display = 'none'">View 2 Replies</a>
                                                            </small>
                                                            <script>
                                                               $(document).ready(function (){
                                                                   $(".commentOfComment").css("display","none");
                                                                });
                                                            </script>
                                                        </p>
                                                         <div class="commentOfComment" id="replies">
                                                             <!-- first reply of comment-->
                                                             <div class="media">
                                                                 <a class="media-left" href="#">
                                                                    <img class="media-object img-rounded" src="https://scontent-lga3-1.xx.fbcdn.net/v/t31.0-8/12514060_499384470233859_6798591731419500290_o.jpg?oh=329ea2ff03ab981dad7b19d9172152b7&oe=5A2D7F0D">
                                                                </a>
                                                                <div class="media-body">
                                                                    <h5 class="media-heading user_name" style="color:cadetblue;">Bryan Brosbyani</h5>
                                                                    <p> Hell No!</p>
                                                                </div>
                                                            </div>
                                                             <!-- second reply of comment-->
                                                             <div class="media">
                                                                <a class="media-left" href="#">
                                                                    <img class="media-object img-rounded" src="http://engineering.nyu.edu/files/imagecache/img_col_3_140/pictures/picture-310.jpg">
                                                                </a>
                                                                <div class="media-body">
                                                                    <h5 class="media-heading user_name" style="color:cadetblue;">Phyllis Frankyl</h5>
                                                                    <p> Naughty boii</p>
                                                                </div>
                                                            </div>
                                                        </div>
                                                        <div class="input-group pull-right">
                                                        <form id="innercommentform">
                                                            <input class="form-control" placeholder="Add a reply" type="text">
                                                        </form>
                                                        <span class="input-group-addon">
                                                            <a onclick="document.getElementById('innercommentform').submit();">
                                                                <i class="fa fa-edit"></i>
                                                            </a>
                                                        </span>
                                                        </div>
                                                    </div>`
		sum += i
	}
	username := session.Values["Username"].(string)
	s := string(parentPost.Element.Content)

	output += `
	 <div class="modal fade" id=share-modal` + postactual + ` role="dialog">
                                <div class="modal-dialog">
                                    <!-- Modal content-->
                                    <div class="modal-content">
                                        <div class="modal-header">
                                            <button type="button" class="close" data-dismiss="modal">&times;</button>
                                            <h4 class="modal-title">Share On Your Profile</h4>
                                            </div>
                                            <div class="modal-body">
                                            <div class="media">
                                                <a class="pull-left" href="#">
                                                    <img class="media-object img-rounded" src=d` + parentPost.Image + `>
                                                </a>
                                                <div class="media-body">
                                                    <h6 class="pull-right text-muted time"></h6>
                                                    <h5 class="mt-0" style="color:cadetblue;">` + parentPost.FirstName + " " + parentPost.LastName + `</h5>
                                                    <p>` + s + `</p> </div>
                                                      <div class="form-group">
                                                <form id="shareCommentForm" method="POST" action="/ShareComment">
                                                    <input type="text" class="form-control" id="comment-msg" name="msg" placeholder="Say Something about this..."></input>
                                                    <!--What is 'odom-submit'? If it's not used, remove it-->
                                                    <input type="hidden" name="postid" value=` + postactual + `>
                                                      <input type="hidden" name = "id" value=` + pika + `>
                                                      <input type ="hidden" name="username" value=` + username + `>
                                                    <button class="btn btn-primary odom-submit">Post</button>
                                                </form>
 
                                                </div>
                                            </div>
                                        </div>
                                         
                                           </div> </div> </div> </div>
                                            </div>
                                                                    <!-- delete confirmation modal -->
                            <div class="modal fade" id="confirm-delete" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
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
                                            <a class="btn btn-danger btn-ok">Delete</a>
                                        </div>
                                    </div>
                                </div>
                            </div>
  


	`

	fmt.Fprintln(w, output)

}

/* This function might not be used anymore. */
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
