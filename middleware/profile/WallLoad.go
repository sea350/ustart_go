package profile

import (
	"fmt"
	"net/http"
	"strings"

	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

func WallLoad(w http.ResponseWriter, r *http.Request) {
	// If followingStatus = no
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	fmt.Println(test1)
	fmt.Println("This is debug text, WallLoad.go: 20")
	r.ParseForm()
	entryIDs := r.FormValue("entryIDs")
	fmt.Println(entryIDs)
	var jEntries []types.JournalEntry

	pageID := r.FormValue("pageID")
	if strings.Compare("null", entryIDs) != 0 {

		actualIDs := strings.Split(entryIDs, ",")
		//jEntriesPointer := &jEntries
		jEntries, err5 := uses.LoadEntries(eclient, actualIDs)
		fmt.Println(jEntries[0].ElementID)
		if err5 != nil {
			fmt.Println(err5)
			fmt.Println("This is an error, WallLoad.go: 34")
		}
	}
	var output string
	DocID := session.Values["DocID"].(string)

	output +=
		`
	<script>
							 $(".comment-btn").hover(function (e) {
                                                    var cmtBtnImg = $(this).find('img');
                                                    cmtBtnImg.attr('src', "/www/ustart.tech/ico/comment.png");     
                                                 },function (e) {
                                                    var cmtBtnImg = $(this).find('img');
                                                    cmtBtnImg.attr('src', "/www/ustart.tech/ico/no comment.png");     
                                                 });   
                                                $(".share-btn").hover(function (e) {
                                                    var shrBtnImg = $(this).find('img');
                                                    shrBtnImg.attr('src', "/www/ustart.tech/ico/share.png");     
                                                 },function (e) {
                                                    var shrBtnImg = $(this).find('img');
                                                    shrBtnImg.attr('src', "/www/ustart.tech/ico/not share.png");     
                                                 });
                                                  $(".like-btn").hover(function (e) {
                                                    var likeBtnImg = $(this).find('img');
                                                    if (likeBtnImg.attr('src') === "/www/ustart.tech/ico/like.png") {
                                                        likeBtnImg.attr('src', "/www/ustart.tech/ico/liked.png");
                                                    } else {
                                                        likeBtnImg.attr('src', "/www/ustart.tech/ico/like.png");
                                                    }
                                                    return false;
                                                });
                                                $(".comment-like").click(function (e) {
                                                    if ($(this).html() == "Like") {
                                                        $(this).html('Liked');
                                                    } else {
                                                        $(this).html('Like');
                                                    }
                                                    return false;
                                                });
                                                  $('body').on('click', '.odom-submit', function (e) {
                                                        $('#shareCommentForm').submit();
                                                    });
                              $('.comment-btn').click(function(e) {
                                        var postId= $(this).attr("id");
                                        var modified ="#"+postId;
                                   //     console.log(modified);
                                        var Pikachu = "` + DocID + `";
                                        //e.preventDefault();
                                        $.ajax({
                                            type: 'GET',  
                                            url: 'http://ustart.today:5000/getComments/',
                                            contentType: "application/json; charset=utf-8",
                                            data: {userID:"` + pageID + `", PostID:postId,Pikachu:Pikachu},
                                            success: function(data) {
                                            	$("#commentnil").html(data);
                                             //   console.log(data);
                                                $(modified).modal('show');
                                            }
                                        });
                                    });

                                    $('.share-btn').click(function(e) {
                                        var postId= $(this).attr("id");
                                        var modified ="#"+postId;
                                        console.log(modified);
                                        var Pikachu = "` + DocID + `";
                                        //e.preventDefault();
                                        $.ajax({
                                            type: 'GET',  
                                            url: 'http://ustart.today:5000/shareComments/',
                                            contentType: "application/json; charset=utf-8",
                                            data: {userID:"` + pageID + `", PostID:postId,Pikachu:Pikachu},
                                            success: function(data) {
                                                $("#sharenil").html(data);
                                                console.log("share clicked");
                                                $(modified).modal('show');
                                            }
                                        });
                                    });    

                                        $('.like-btn').click(function(e) {
                                        var postId= $(this).attr("id");
                                        var modified ="#"+postId;
                                        console.log(modified);
                                        var selfDoc = "` + DocID + `";
                                        //e.preventDefault();
                                        $.ajax({
                                            type: 'GET',  
                                            url: 'http://ustart.today:5000/Like',
                                            contentType: "application/json; charset=utf-8",
                                            data: {userID:"` + pageID + `", PostID:postId,selfDoc:selfDoc},
                                            success: function(data) {
                                                    var likeBtnImg = $(this).find('img');
                                                    if (likeBtnImg.attr('src') === "/www/ustart.tech/ico/like.png") {
                                                        likeBtnImg.attr('src', "/www/ustart.tech/ico/liked.png");
                                                    } else {
                                                        likeBtnImg.attr('src', "/www/ustart.tech/ico/like.png");
                                                    }
                                                console.log("like clicked");
                                            }
                                        });
                                    }); 

      

	</script>
	`

	if strings.Compare("null", entryIDs) != 0 {
		//fmt.Println("ROOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOAR")
		sum := 0
		class0 := `<div class="panel panel-default wallAppend">`

		for i := len(jEntries) - 1; i >= 0; i-- {
			sum++
			bodyText := string(jEntries[i].Element.Content)
			if jEntries[i].Element.Classification == 0 {
				//fmt.Println("classifcation is 0")
				likes := string(jEntries[i].NumLikes)
				fmt.Println("this post has" + string(jEntries[i].NumLikes) + " likes")
				class0 += `<div id="wallPosts" class="panel-body">
                                            <!-- regular post sample -->
                                            <div class="media">
                                                <a class="pull-left" href="#">
                                                    <img style="height:40px;" class="media-object img-rounded" src=d` + jEntries[i].Image + `>
                                                </a>
                                                <!--edit dropdown -->
                                                <div class="dropdown pull-right">
                                                    <a class="dropdown-toggle" data-toggle="dropdown"><span class="glyphicon glyphicon-cog"></span><span class="caret"></span></a>
                                                    <ul class="dropdown-menu" style="min-width: 0px !important; padding:0px !important;">
                                                        <li>
                                                            <a class="dropdown-item " data-toggle="modal" data-target="#EditModal">
                                                                <H6>Edit</H6>
                                                            </a>
                                                        </li>
                                                        <li>
                                                            <a class="dropdown-item " data-toggle="modal" data-target="#confirm-delete">
                                                                <H6>Delete</H6>
                                                            </a>
                                                        </li>
                                                    </ul>
                                                </div>
                                                <!--end edit dropdown -->
                                            <div class="media-body">
                                                <h6 class="post-time pull-right text-muted time" style="padding-right:4px;"> X minutes ago` + /*jEntries[i].Element.TimeStamp+*/ `</h6>
                                                <h5 class="post-name mt-0" style="color:cadetblue;"><a href="#">` + jEntries[i].FirstName + `</a></h5>
                                                <p class="post-message" style="word-spacing: 0px;">` + bodyText + `</p>
                                                                            </div>
                                                <ul>
                                                    <li>
                                                         <a class="btn btn-sm like-btn" id =main-modal` + jEntries[i].ElementID + `><img class="like-btn-ico" src="/www/ustart.tech/ico/like.png">  <p class="mt-0" style="color:cadetblue; display:inline;">` + likes + `</p></a>
                                                    </li>
                                                    <li>
                                                         <a class="btn btn-sm comment-btn" id =main-modal` + jEntries[i].ElementID + `><img class="coment-btn-ico" src="/www/ustart.tech/ico/no comment.png">  <p class="mt-0" style="color:cadetblue; margin-left:1px; display:inline;">` + string(jEntries[i].NumReplies) + `</p></a>
                                                    </li>
                                                    <li>
                                                         <a class="btn btn-sm share-btn" id=share-modal` + jEntries[i].ElementID + `><span><img class="share-btn-ico" src="/www/ustart.tech/ico/not share.png"> <p class="mt-0" style="margin-left:1px; color:cadetblue; display:inline;">` + string(jEntries[i].NumShares) + `</p></span></a>
                                                    </li>
                                                </ul>
                                            </div>
                                            <!-- end regular post -->
                                            
                                            <hr>
                                            </div>
                                            `

			}
			if jEntries[i].Element.Classification == 2 {
				postIDArray := []string{jEntries[i].Element.ReferenceEntry} // just an array with 1 entry
				jEntry, err5 := uses.LoadEntries(eclient, postIDArray)
				if err5 != nil {
					fmt.Println(err5)
					fmt.Println("This is an error, WallLoad.go: 207")
				}
				bodyText := string(jEntry[0].Element.Content)
				comment := string(jEntries[i].Element.Content)

				class0 += `
						 <div class="dropdown pull-right">
                                            <a class="dropdown-toggle" data-toggle="dropdown">
                                                <span class="glyphicon glyphicon-cog"></span>
                                                <span class="caret"></span>
                                            </a>
                                            <ul class="dropdown-menu" style="min-width: 0px !important; padding:0px !important;">
                                                <li>
                                                    <a class="dropdown-item " data-toggle="modal" data-target="#EditModal">
                                                        <H6>Edit</H6>
                                                    </a>
                                                </li>
                                                <li>
                                                    <a class="dropdown-item" data-toggle="modal" data-target="#confirm-delete">
                                                        <H6>Delete</H6>
                                                    </a>
                                                </li>
                                            </ul>
                                        </div>
                                        <!--end edit dropdown -->
                                        <h6 class="pull-right text-muted time" style="padding-right:4px;">X hours ago</h6>
                                        <h5 class="mt-0" style="color:cadetblue">You shared a post:</h5>
                                           <p style="margin-left:2em">` + comment + `</p>
                                        <div class="media">
                                            <div class="panel panel-default">
                                                <div class="panel-body">
                                                    <div class="media">
                                                        <a class="pull-left" href="#">
                                                            <img class="media-object img-rounded" src=d` + jEntry[0].Image + ` alt="40x40">
                                                        </a>
                                                        <div class="media-body">
                                                            <h6 class="pull-right text-muted time">X hours ago</h6>
                                                            <h5 class="mt-0" style="color:cadetblue;">` + jEntry[0].FirstName + " " + jEntry[0].LastName + `</h5>
                                                            <p>` + bodyText + `</p>
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                                <ul>
                                                    <li>
                                                         <a class="btn btn-sm like-btn" id =main-modal` + jEntry[0].ElementID + `><img class="like-btn-ico" src="/www/ustart.tech/ico/like.png">  <p class="mt-0" style="color:cadetblue; display:inline;">` + string(jEntry[0].NumLikes) + `</p></a>
                                                    </li>
                                                    <li>
                                                         <a class="btn btn-sm comment-btn" id =main-modal` + jEntry[0].ElementID + `><img class="coment-btn-ico" src="/www/ustart.tech/ico/no comment.png">  <p class="mt-0" style="color:cadetblue; margin-left:1px; display:inline;">` + string(jEntry[0].NumReplies) + `</p></a>
                                                    </li>
                                                    <li>
                                                         <a class="btn btn-sm share-btn" id=share-modal` + jEntry[0].ElementID + `><span><img class="share-btn-ico" src="/www/ustart.tech/ico/not share.png"> <p class="mt-0" style="margin-left:1px; color:cadetblue; display:inline;">` + string(jEntry[0].NumShares) + `</p></span></a>
                                                    </li>
                                                </ul>
                                        </div>



`

			}

		}
		class0 += "</div>"
		output += class0
	}
	//	output += "</div>" // should be last line, closes main id
	output += `
								<script>
	                                 $('#new-postSubmit').click(function(e) {
                                        //e.preventDefault();
                                        var docID = "` + DocID + `";
                                        var text = $("#post-msg").val();
                                        console.log(text);
                                        $.ajax({
                                            type: 'GET',  
                                            url: 'http://ustart.today:5000/addPost/',
                                            contentType: "application/json; charset=utf-8",
                                            data: {docID:docID,text:text},
                                            success: function(data) {
                                            //	console.log(data);
                                                $(".wallAppend").prepend(data);

                                                console.log('hello m8');

                                                							 $(".comment-btn").hover(function (e) {
                                                    var cmtBtnImg = $(this).find('img');
                                                    cmtBtnImg.attr('src', "/www/ustart.tech/ico/comment.png");     
                                                 },function (e) {
                                                    var cmtBtnImg = $(this).find('img');
                                                    cmtBtnImg.attr('src', "/www/ustart.tech/ico/no comment.png");     
                                                 });   
                                                $(".share-btn").hover(function (e) {
                                                    var shrBtnImg = $(this).find('img');
                                                    shrBtnImg.attr('src', "/www/ustart.tech/ico/share.png");     
                                                 },function (e) {
                                                    var shrBtnImg = $(this).find('img');
                                                    shrBtnImg.attr('src', "/www/ustart.tech/ico/not share.png");     
                                                 });
                                                  $(".like-btn").hover(function (e) {
                                                    var likeBtnImg = $(this).find('img');
                                                    if (likeBtnImg.attr('src') === "/www/ustart.tech/ico/like.png") {
                                                        likeBtnImg.attr('src', "/www/ustart.tech/ico/liked.png");
                                                    } else {
                                                        likeBtnImg.attr('src', "/www/ustart.tech/ico/like.png");
                                                    }
                                                    return false;
                                                });
                                                $(".comment-like").click(function (e) {
                                                    if ($(this).html() == "Like") {
                                                        $(this).html('Liked');
                                                    } else {
                                                        $(this).html('Like');
                                                    }
                                                    return false;
                                                });
                                                  $('body').on('click', '.odom-submit', function (e) {
                                                        $('#shareCommentForm').submit();
                                                    });
                              $('.comment-btn').click(function(e) {
                                        var postId= $(this).attr("id");
                                        var modified ="#"+postId;
                                        console.log(modified);
                                        var Pikachu = "` + DocID + `";
                                        //e.preventDefault();
                                        $.ajax({
                                            type: 'GET',  
                                            url: 'http://ustart.today:5000/getComments/',
                                            contentType: "application/json; charset=utf-8",
                                            data: {userID:"` + pageID + `", PostID:postId,Pikachu:Pikachu},
                                            success: function(data) {
                                            	$("#commentnil").html(data);
                                                console.log(data);
                                                $(modified).modal('show');
                                            }
                                        });
                                    });

                                         
                                        $('.share-btn').click(function(e) {
                                        var postId= $(this).attr("id");
                                        var modified ="#"+postId;
                                        console.log(modified);
                                        var Pikachu = "` + DocID + `";
                                        //e.preventDefault();
                                        $.ajax({
                                            type: 'GET',  
                                            url: 'http://ustart.today:5000/shareComments/',
                                            contentType: "application/json; charset=utf-8",
                                            data: {userID:"` + pageID + `", PostID:postId,Pikachu:Pikachu},
                                            success: function(data) {
                                                $("#sharenil").html(data);
                                                console.log("share clicked ");
                                                $(modified).modal('show');
                                            }
                                        });
                                    });  

                                        $('.like-btn').click(function(e) {
                                        var postId= $(this).attr("id");
                                        var modified ="#"+postId;
                                        console.log(modified);
                                        var selfDoc = "` + DocID + `";
                                        //e.preventDefault();
                                        $.ajax({
                                            type: 'GET',  
                                            url: 'http://ustart.today:5000/Like',
                                            contentType: "application/json; charset=utf-8",
                                            data: {userID:"` + pageID + `", PostID:postId,selfDoc:selfDoc},
                                            success: function(data) {
                                                    var likeBtnImg = $(this).find('img');
                                                    if (likeBtnImg.attr('src') === "/www/ustart.tech/ico/like.png") {
                                                        likeBtnImg.attr('src', "/www/ustart.tech/ico/liked.png");
                                                    } else {
                                                        likeBtnImg.attr('src', "/www/ustart.tech/ico/like.png");
                                                    }
                                                console.log("like clicked");
                                            }
                                        });
                                    }); 

                                           }
                                        });
                                    });      


      
                                </script> `

	//var responseHtml string
	fmt.Fprintln(w, output) // this line sends data back to the ajax call on the front end side
}
