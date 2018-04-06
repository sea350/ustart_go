/*package stringHTML

//CommentEntry ... creates a string of html that displays a comment
func CommentEntry(image string, fName string, lName string, content string, postID string) string {
	return `
	<div class="media">
		<h6 class="pull-right text-muted time">2 hours ago</h6>
		<a class="media-left" href="#">
			<img style="height:40px;" class="img-rounded" src=d` + image + `>
		</a>
		<div class="media-body">
			<h5 class="media-heading user_name" style="color:cadetblue;">` + fName + `comment.go` + " " + lName + postID + `</h5>
			<p>` + content + `</p>
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
				$(".` + postID + `").css("display","none");
				});
			$('.view-replies').click(function(e) {
        e.preventDefault();
        var postId = "` + postID + `"
        var modified ="#` + postID + `"
        console.log(modified);
        $.ajax({
            type: 'GET',
            url: 'http://ustart.today:5000/getPostComments/',
            contentType: "application/json; charset=utf-8",
            data: {PostID:postId},
            success: function(data) {
				console.log("succcerosed");
				console.log(` + postID + `);
         	   $(".commentOfComment").html(data);

            }//end success
        });
    }); // end view-replies click
			</script>


			</p>
				<div class="commentOfComment" id="replies">

					<div class="media">
						<a class="media-left" href="#">
							<img class="media-object img-rounded" src="https://scontent-lga3-1.xx.fbcdn.net/v/t31.0-8/12514060_499384470233859_6798591731419500290_o.jpg?oh=329ea2ff03ab981dad7b19d9172152b7&oe=5A2D7F0D">
						</a>
						<div class="media-body">
							<h5 class="media-heading user_name" style="color:cadetblue;">Bryan Brosbyani</h5>
							<p> Hell No!</p>
						</div>
					</div>

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
				<form id="innercommentform" method="POST" action="/AddComment2">
				    <input type = "hidden" name="postID" value= "` + postID + `">
					<input class="form-control" placeholder="Add a reply" type="text" name="body">
				</form>
				<span class="input-group-addon">
					<a onclick="document.getElementById('innercommentform').submit();">
						<i class="fa fa-edit"></i>
					</a>
				</span>
		</div>
	</div>`
}*/

package stringHTML

//CommentEntry ... creates a string of html that displays a comment
func CommentEntry(image string, fName string, lName string, content string, postID string) string {
	return `
	<div class="media">
		<h6 class="pull-right text-muted time">2 hours ago</h6>
		<a class="media-left" href="#">
			<img style="height:40px;" class="img-rounded" src=d` + image + `>
		</a>
		<div class="media-body">
			<h5 class="media-heading user_name" style="color:cadetblue;">` + fName + `comment.go` + " " + lName + postID + `</h5>
			<p>` + content + `</p>
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
			$('.view-replies').click(function(e) {
        e.preventDefault();
        var postId = "` + postID + `"
        var modified ="#` + postID + `"
        console.log(modified);
        $.ajax({
            type: 'GET',  
            url: 'http://ustart.today:5000/getPostComments/',
            contentType: "application/json; charset=utf-8",
            data: {PostID:postId},
            success: function(data) {
            	console.log("succcerosed");
         	   $(".commentOfComment").html(data);
            }//end success
        });
    }); // end view-replies click 
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
				<form id="innercommentform" method="POST" action="/AddComment2">
				    <input type = "hidden" name="postID" value= "` + postID + `">
					<input class="form-control" placeholder="Add a reply" type="text" name="body">
				</form>
				<span class="input-group-addon">
					<a onclick="document.getElementById('innercommentform').submit();">
						<i class="fa fa-edit"></i>
					</a>
				</span>
		</div>
	</div>`
}
