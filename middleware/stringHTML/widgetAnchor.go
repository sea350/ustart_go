package stringHTML

//WidgetAnchor ... strigifies HTML for widget
func WidgetAnchor(url string) string {
	output := `
<li class="ui-state-default widgetListItem sortable">
	<div class="projectsWidgetCont">
		<div class="widgetTitle">
		<span class="pull-right fa fa-2x fa-sort"></span>
		<span class="pull-right fa fa-2x fa-trash"></span>
		<span class="pull-right fa fa-2x fa-pencil" id="ar-edit" data-toggle="modal" data-target="#ar-modal"></span>
			<h4>Anchor</h4>
		</div>

		<div class="modal fade" id="ar-modal" role="dialog">
			<div class="modal-dialog">
				<div class="modal-content">
					<div class="modal-header">
						<button type="button" class="close" data-dismiss="modal">&times;</button>
						<h4 class="modal-title">Embed Anchors</h4>
					</div>
					<div class="modal-body">
						<span id="ar-list-title">Removable Anchors</span>
						<ul id="ar-edit-list"></ul>
						<p>Insert the Anchor Code in the URL:</p>
						<div class="form-group">
							<input type="text" class="form-control" id="ar-embed-input">

						</div>
					</div>
					<div class="modal-footer">
						<button type="submit" id="ar-submit-btn" class="btn btn-default" data-dismiss="modal">Submit</button>

						<button type="button" class="btn btn-primary" data-dismiss="modal">Close</button>
					</div>
				</div>
			</div>
		</div> 

		<div id= "widgetBodyAR" class="widgetBody">
			<div class="anchor-feed"><iframe src="https://anchor.fm/e/' + ` + url + ` + '" height="270px" width="400px" frameborder="0" scrolling="no"></iframe></div>
		</div>
	</div>
</li>
`
	return output
}
