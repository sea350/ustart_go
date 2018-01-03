package profile 

import (
    "net/http"
    uses "github.com/sea350/ustart_go/uses"
    "fmt"
)


func WallAdd(w http.ResponseWriter, r *http.Request){
	// If followingStatus = no 
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
    if (test1 == nil){
     	fmt.Println(test1)
    http.Redirect(w, r, "/~", http.StatusFound)
    }

	r.ParseForm()
	docID := r.FormValue("docID")
	text := r.FormValue("text")
	textRunes := []rune(text)
    postID, err := uses.UserNewEntry(eclient,docID,textRunes)
    if (err != nil){
    	fmt.Println(err);
    }
    postIDArray := []string{postID} // just an array with 1 entry 
    jEntry, err5 := uses.LoadEntries(eclient,postIDArray)
	if (err5 != nil){
		fmt.Println(err5);
	}


	var output string 
	// output += ` <div id="main"> `
                                  //class0 := `<div class="panel panel-default">` 

		bodyText := string(jEntry[0].Element.Content)
			class0 := `<div id="wallPosts" class="panel-body">
                                            <!-- regular post sample -->
                                            <div class="media">
                                                <a class="pull-left" href="#">
                                                    <img style="height:40px;" class="media-object img-rounded" src=d`+jEntry[0].Image+`>
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
                                                <h6 class="post-time pull-right text-muted time" style="padding-right:4px;"> X minutes ago`+/*jEntries[i].Element.TimeStamp+*/`</h6>
                                                <h5 class="post-name mt-0" style="color:cadetblue;"><a href="#">`+jEntry[0].FirstName+`</a></h5>
                                                <p class="post-message" style="word-spacing: 0px;">`+bodyText+`</p>
                                                                            </div>
                                                <ul>
                                                    <li>
                                                         <a class="btn btn-sm like-btn" id =main-modal`+jEntry[0].ElementID+`><img class="like-btn-ico" src="/www/ustart.tech/ico/like.png">  <p class="mt-0" style="color:cadetblue; display:inline;">`+string(jEntry[0].NumLikes)+`</p></a>
                                                    </li>
                                                    <li>
                                                         <a class="btn btn-sm comment-btn" id =main-modal`+jEntry[0].ElementID+`><img class="coment-btn-ico" src="/www/ustart.tech/ico/no comment.png">  <p class="mt-0" style="color:cadetblue; margin-left:1px; display:inline;">`+string(jEntry[0].NumReplies)+`</p></a>
                                                    </li>
                                                    <li>
                                                         <a class="btn btn-sm share-btn" id=share-modal`+jEntry[0].ElementID+`><span><img class="share-btn-ico" src="/www/ustart.tech/ico/not share.png"> <p class="mt-0" style="margin-left:1px; color:cadetblue; display:inline;">`+string(jEntry[0].NumShares)+`</p></span></a>
                                                    </li>
                                                </ul>
                                            </div>
                                            <!-- end regular post -->
                                            
                                            <hr>
                                            </div>
                                            `



		
//	class0 += "</div>"
	output += class0

//	output += "</div>" // should be last line, closes main id



	fmt.Println(output)
	//var responseHtml string 
	fmt.Fprintln(w, output) 
}