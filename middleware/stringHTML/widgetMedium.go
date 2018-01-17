package stringHTML

//WidgetMedium ... strigifies HTML for widget WARNING INCOMPLETE
func WidgetMedium(html string) string {
	output := `
<li class="ui-state-default widgetListItem sortable">
	<div class="projectsWidgetCont">
		<div class="widgetTitle">
		<span class="pull-right fa fa-2x fa-sort"></span>
		<span class="pull-right fa fa-2x fa-trash"></span>
		<span class="pull-right fa fa-2x fa-pencil" id="med-edit" data-toggle="modal" data-target="#med-modal"></span>
			<h4>Medium</h4>
		</div>

		<div class="modal fade" id="med-modal" role="dialog">
			<div class="modal-dialog">
				<div class="modal-content">
					<div class="modal-header">
						<button type="button" class="close" data-dismiss="modal">&times;</button>
						<h4 class="modal-title">Embed Medium</h4>
					</div>
					<div class="modal-body">
						<p>Select a Medium import option, either your username, or your publication name</p>
						<div class="input-group">
							<div class="input-group-btn">
								<button id="med-setting" type="button" class="btn btn-secondary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
									Username <i class="glyphicon glyphicon-menu-down"></i>
								</button>
								<div class="dropdown-menu">
									<a id="med-setting-user" class="dropdown-item" style="display:block;">Username</a>
									<a id="med-setting-pub" class="dropdown-item" style="display:block;">Publication</a>
								</div>
							</div>
							<input type="text" class="form-control" id="med-embed-username"/>
							<input type="text" class="form-control" id="med-embed-publication" disabled/>
						</div>
						<br/>
						<div class="input-group">
							<span class="input-group-addon">Publication Tag</span>
							<input type="text" class="form-control" id="med-embed-tag" disabled placeholder="(Optional)"/>
						</div>
						<br/>
						<div class="input-group">
							<span class="input-group-addon">Article Count</span>
							<input type="number" class="form-control" id="med-count" min='1' max='12' value='4'/>
						</div>
					</div>
					<div class="modal-footer">
						<button type="submit" id="med-submit-btn" class="btn btn-default" data-dismiss="modal">Submit</button>
						<button type="button" class="btn btn-primary" data-dismiss="modal">Close</button>
					</div>
				</div>
			</div>
		</div> 

		<div id="widgetBodyMed" class="widgetBody"></div>
	</div>
</li>
`
	return output
}
