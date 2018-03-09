package stringHTML

//OutputShare ... use isnt exactly known WARNING: has some unknown variables
func OutputShare(postactual string, image string, fName string, lName string, content string, pika string, username string) string {
	return `
	<div class="modal fade" id=share-modal` + postactual + ` role="dialog">
		<div class="modal-dialog">
			<!-- Modal content-->
			<div class="modal-content">
				<div class="modal-header">
					<button type="button" class="close" data-dismiss="modal">&times;</button>
					<h4 class="modal-title">Share On Your Profile shareComment.go</h4>
				</div>
				<div class="modal-body">
					<div class="media">
						<a class="pull-left" href="#">
							<img class="media-object img-rounded" src=d` + image + `>
						</a>
						<div class="media-body">
							<h6 class="pull-right text-muted time"></h6>
							<h5 class="mt-0" style="color:cadetblue;">` + fName + " " + lName + `</h5>
							<p>` + content + `</p>
						</div>
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
			</div>
		</div>
	</div>
</div>
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
}
