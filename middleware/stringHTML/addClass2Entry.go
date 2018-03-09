package stringHTML

//AddClass2Entry ... Used to add a Class 0 entry to a display html string
func AddClass2Entry(comment string, image string, fName string, lName string, content string, elementID string, numLikes string, numReplies string, numShares string) string {
	return `
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
				<a class="dropdown-item" data-toggle="modal" data-target=#confirm-delete` + elementID + `>
					<H6>Delete1</H6>
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
						<img class="media-object img-rounded" src=d` + image + ` alt="40x40">
					</a>
					<div class="media-body">
						<h6 class="pull-right text-muted time">X hours ago</h6>
						<h5 class="mt-0" style="color:cadetblue;">` + fName + " " + lName + `</h5>
						<p>` + content + `</p>
					</div>
				</div>
			</div>
		</div>
		<ul>
			<li>
			<a class="btn btn-sm like-btn" id =main-modal` + elementID + `>
				<img class="like-btn-ico" src="/ustart_front/ico/like.png">
				<p class="mt-0" style="color:cadetblue; display:inline;">` + numLikes + `</p>
			</a>
			</li>
			<li>
				<a class="btn btn-sm comment-btn" id =main-modal` + elementID + `>
					<img class="coment-btn-ico" src="/ustart_front/ico/no comment.png">
					<p class="mt-0" style="color:cadetblue; margin-left:1px; display:inline;">` + numReplies + `</p>
				</a>
			</li>
			<li>
				<a class="btn btn-sm share-btn" id=share-modal` + elementID + `>
					<span>
						<img class="share-btn-ico" src="/ustart_front/ico/not share.png">
						<p class="mt-0" style="margin-left:1px; color:cadetblue; display:inline;">` + numShares + `</p>
					</span>
				</a>
			</li>
		</ul>

	</div>


	
	`
}
