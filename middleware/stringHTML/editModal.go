package stringHTML

//EditModal ... creates a string of html that displays a comment WARNING: has some unknown variables
func EditModal(postDocID string, image string, content string) string {
	return `
<div class="modal fade" id="EditModa` + postDocID + `" role="dialog">
<div class="modal-dialog">
	<!-- Modal content-->
	<div class="modal-content">
		<div class="modal-header">
			<button type="button" class="close" data-dismiss="modal">&times;</button>
			<h4 class="modal-title">Edit Post</h4>
		</div>
		<div class="modal-body">
			<div class="media">
				<a class="pull-left" href="#">
					<img class="media-object img-rounded" src="` + image + `"
						alt="64x64">
				</a>
				<div class="media-body">
					<div class="form-group">
						<form id="Edit-Post-Form">
							<textarea class="form-control" id="post-msg" style="resize:none;" placeholder="">` + content + `</textarea>
						</form>
					</div>
				</div>
			</div>
		</div>
		<div class="modal-footer">
			<button id="edit-postSubmit" class="btn btn-primary pull-right">Post</button>
		</div>
	</div>
</div>
</div>`
}
