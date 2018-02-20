package stringHTML

//AddClass0Entry ... Used to add a Class 0 entry to a display html string
func AddClass0Entry(image string, fName string, content string, elementID string, numLikes string, numReplies string, numShares string) string {
	return `<div id="wallPosts" class="panel-body">
	<!-- regular post sample -->
	<div class="media">
		<a class="pull-left" href="#">
			<img style="height:40px;" class="media-object img-rounded" src=d` + image + `>
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
		<h5 class="post-name mt-0" style="color:cadetblue;"><a href="#">` + fName + `</a></h5>
		<p class="post-message" style="word-spacing: 0px;">` + content + `</p>
									</div>
		<ul>
			<li>
				 <a class="btn btn-sm like-btn" id =main-modal` + elementID + `><img class="like-btn-ico" src="/ustart_front/ico/like.png">  <p class="mt-0" style="color:cadetblue; display:inline;">` + numLikes + `</p></a>
			</li>
			<li>
				 <a class="btn btn-sm comment-btn" id =main-modal` + elementID + `><img class="coment-btn-ico" src="/ustart_front/ico/no comment.png">  <p class="mt-0" style="color:cadetblue; margin-left:1px; display:inline;">` + numReplies + `</p></a>
			</li>
			<li>
				 <a class="btn btn-sm share-btn" id=share-modal` + elementID + `><span><img class="share-btn-ico" src="/ustart_front/ico/not share.png"> <p class="mt-0" style="margin-left:1px; color:cadetblue; display:inline;">` + numShares + `</p></span></a>
			</li>
		</ul>
	</div>
	<!-- end regular post -->
	
	<hr>
	</div>
	`
}
