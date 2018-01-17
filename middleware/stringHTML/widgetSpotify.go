package stringHTML

//WidgetSpotify ... strigifies HTML for widget
func WidgetSpotify(html string) string {
	output := `
<li class="ui-state-default widgetListItem sortable">
	<div class="projectsWidgetCont">
		<div class="widgetTitle">
			<span class="pull-right fa fa-2x fa-sort"></span>
			<span class="pull-right fa fa-2x fa-trash"></span>
			<span class="pull-right fa fa-2x fa-pencil" id="spot-edit" data-toggle="modal" data-target="#spot-modal"></span>

				<div class="modal fade" id="spot-modal" role="dialog">
					<div class="modal-dialog">
						<div class="modal-content">
							<div class="modal-header">
								<button type="button" class="close" data-dismiss="modal">&times;</button>
								<h4 class="modal-title">Embed Spotify</h4>
							</div>
							<div class="modal-body">
								<span id="spot-list-title">Removable Spotify Sounds</span>
								<ul id="spot-edit-list"></ul>
								<p>Please paste the Spotify embed code in here:</p>
								<div class="form-group">
									<input type="text" class="form-control" id="spot-embed-input">

								</div>
							</div>
							<div class="modal-footer">
								<button type="submit" id="spot-submit-btn" class="btn btn-default" data-dismiss="modal">Submit</button>

								<button type="button" class="btn btn-primary" data-dismiss="modal">Close</button>
							</div>
						</div>
					</div>
				</div> 


			<h4 name="textEditorHeader">
				Spotify
			</h4>
		</div>
		<div id="widgetBodySpot" class="widgetBody">
			` + html + `
		</div>
	</div>
</li>
`
	return output
}
