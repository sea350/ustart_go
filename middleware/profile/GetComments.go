package profile 

import (
    "net/http"
    uses "github.com/sea350/ustart_go/uses"
    "fmt"
)


func GetComments(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
    if (test1 == nil){
        //No docid in session
        http.Redirect(w, r, "/~", http.StatusFound)
    }

	r.ParseForm()
	postID := r.FormValue("PostID")
	postaid := postID[9:]
	postactual := postID[10:]
	pika := r.FormValue("Pikachu")
	// journal entry, err 
	parentPost, arrayofComments, err4 := uses.LoadComments(eclient, postactual, 0, -1)
	if (err4 != nil){
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
                                                        <img style="height:40px;" class="img-rounded" src=d`+arrayofComments[i].Image+`>
                                                    </a>
                                                    <div class="media-body">
                                                        <h5 class="media-heading user_name" style="color:cadetblue;">`+arrayofComments[i].FirstName+" "+arrayofComments[i].LastName+`</h5>
                                                        <p>`+commentBody+`</p>
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
		fmt.Println(arrayofComments[i].FirstName)
		sum += i
	}
	username := session.Values["Username"].(string)
	s := string(parentPost.Element.Content)

	output += `
	 <div class="modal fade" id=main-moda`+postaid+` role="dialog">
                                <div class="modal-dialog">
                                    <!-- Modal content-->
                                    <div class="modal-content">
                                        <div class="modal-header">
                                            <button type="button" class="close" data-dismiss="modal">&times;</button>
                                            <div class="media">
                                                <a class="pull-left" href="#">
                                                    <img class="media-object img-rounded" src=d`+parentPost.Image+`>
                                                </a>
                                                <div class="media-body">
                                                    <h6 class="pull-right text-muted time"></h6>
                                                    <h5 class="mt-0" style="color:cadetblue;">`+parentPost.FirstName +" "+parentPost.LastName+`</h5>
                                                    <p>`+s+`</p>
                                                </div>
                                            </div>
                                        </div>
                                        <div class="modal-body">
                                            <div class="input-group">
                                                <form class="commentform" method="POST" action="/AddComment">
                                                    <input name="commentz" class="form-control" placeholder="Add a comment" type="text">
                                                      <input type="hidden" name="followstat" value=`+postaid+`>
                                                      <input type="hidden" name = "id" value=`+pika+`>
                                                      <input type ="hidden" name="username" value=`+username+`>
                                                </form>
                                                <span class="input-group-addon">
                                                    <a onclick="document.getElementByClass('commentform').submit();">
                                                    <script>
                                                    console.log('inside the its not gonna work because it's just hml stuff so put inside script')
                                                    </script>
                                                        <i class="fa fa-edit"></i>
                                                    </a>
                                                </span>
                                            </div>
                                            <br>
                                            <div class="comments-list">
                                                `+commentoutputs+`
                                                </div>    `

	fmt.Fprintln(w, output) 

}
