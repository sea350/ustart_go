package stringHTML

//ParentEntry ... creates a string of html that displays a comment WARNING: has some unknown variables
func ParentEntry(postaid string, image string, fName string, lName string, content string, pika string, username string, commentOutputs string) string {
	return `
	<div class="modal fade" id=main-moda` + postaid + ` role="dialog">
		<div class="modal-dialog">
			<!-- Modal content-->
				<div class="modal-content">
					<div class="modal-header">
						<button type="button" class="close" data-dismiss="modal">&times;</button>
						<div class="media">
							<a class="pull-left" href="#">
								<img class="media-object img-rounded" src=d` + image + `>
							</a>
							<div class="media-body">
								<h6 class="pull-right text-muted time"></h6>
								<h5 class="mt-0" style="color:cadetblue;">` + fName + " " + lName + `</h5>
								<p>` + content + `</p>
							</div>
						</div>
					</div>
						<div class="modal-body">
							<div class="input-group">
								<form class="commentform" method="POST" action="/AddComment">
									<input name="commentz" class="form-control" placeholder="Add a comment" type="text">
									<input type="hidden" name="followstat" value=` + postaid + `>
									<input type="hidden" name = "id" value=` + pika + `>
									<input type ="hidden" name="username" value=` + username + `>
									<span class="input-group-addon">
										<a onclick="this.submit();">
											<script>
												console.log('inside the its not gonna work because its just hml stuff so put inside script')
											</script>
											<i class="fa fa-edit"></i>
										</a>
									</span>
								</form>
							</div>
							<br>
							<div class="comments-list">
								` + commentOutputs + `
							</div>`
}
