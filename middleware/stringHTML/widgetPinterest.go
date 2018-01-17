package stringHTML

//WidgetPinterest ... strigifies HTML for Pinterest widget
func WidgetPinterest(url string) string {
	output := `
<li class="ui-state-default widgetListItem sortable">
	<div class="projectsWidgetCont">
		<div class="widgetTitle">
			<span class="pull-right fa fa-2x fa-sort"></span>
			<span class="pull-right fa fa-2x fa-trash"></span>
			<span class="pull-right fa fa-2x fa-pencil" id="pin-edit" data-toggle="modal" data-target="#pin-modal"></span>
			<h4 name="textEditorHeader">
				Pinterest
			</h4>
		</div>
		<div class="modal fade" id="pin-modal" role="dialog">
			<div class="modal-dialog">
				<div class="modal-content">
					<div class="modal-header">
						<button type="button" class="close" data-dismiss="modal">&times;</button>
						<h4 class="modal-title">Embed Pinterest</h4>
					</div>
					<div class="modal-body">
						<span id="pin-list-title">Removable Pins</span>
						<ul id="pin-edit-list"></ul>
						<p>Please paste the Pinterest Pin URL here:</p>
						<div class="form-group">
							<input type="text" class="form-control" id="pin-embed-input">

						</div>
					</div>
					<div class="modal-footer">
						<button type="submit" id="pin-submit-btn" class="btn btn-default" data-dismiss="modal">Submit</button>

						<button type="button" class="btn btn-primary" data-dismiss="modal">Close</button>
					</div>
				</div>
			</div>
		</div> 
		<div id="widgetBodyPin" class="widgetBody">
			` + url + `
		</div>
	</div>
</li>
`
	return output
}
