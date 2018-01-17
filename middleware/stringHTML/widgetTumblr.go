package stringHTML

//WidgetTumblr ... strigifies HTML for tumblr widget
func WidgetTumblr(url string) string {
	output := `
<li class="ui-state-default widgetListItem sortable">
	<div class="projectsWidgetCont">
		<div class="widgetTitle">
			<span class="pull-right fa fa-2x fa-sort"></span>
			<span class="pull-right fa fa-2x fa-trash"></span>
			<span class="pull-right fa fa-2x fa-pencil" id="tumblr-edit" data-toggle="modal" data-target="#tumblr-modal"></span>
			<h4 name="textEditorHeader">
				Tumblr
			</h4>
		</div>
		<div class="modal fade" id="tumblr-modal" role="dialog">
			<div class="modal-dialog">
				<div class="modal-content">
					<div class="modal-header">
						<button type="button" class="close" data-dismiss="modal">&times;</button>
						<h4 class="modal-title">Embed Tumblr</h4>
					</div>
					<div class="modal-body">
						<span id="tumblr-list-title">Removable Tumblr Posts</span>
						<ul id="tumblr-edit-list"></ul>
						<p>Please paste the Tumblr embed code in here:</p>
						<div class="form-group">
							<input type="text" class="form-control" id="tumblr-embed-input">
						</div>
					</div>
					<div class="modal-footer">
						<button type="submit" id="tumblr-submit-btn" class="btn btn-default" data-dismiss="modal">Submit</button>

						<button type="button" class="btn btn-primary" data-dismiss="modal">Close</button>
					</div>
				</div>
			</div>
		</div>
		<div id= "widgetBodyTumblr" class="widgetBody">
			<!-- embeded code goes inside here -->
		</div>
	</div>
</li>
`
	return output
}
