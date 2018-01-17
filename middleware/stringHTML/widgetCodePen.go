package stringHTML

//WidgetCodePen ... strigifies HTML for a code pen widget
func WidgetCodePen(codePenID string) string {
	output := `
<li class="ui-state-default widgetListItem sortable">
	<div class="projectsWidgetCont">
		<div class="widgetTitle">
			<span class="pull-right fa fa-2x fa-sort"></span>
			<span class="pull-right fa fa-2x fa-trash"></span>
			<span class="pull-right fa fa-2x fa-pencil id="code-edit" data-toggle="modal" data-target="#code-modal""></span>
			<h4 name="textEditorHeader">
				CodePen
			</h4>
		</div>

		<div class="modal fade" id="code-modal" role="dialog">
			<div class="modal-dialog">
				<div class="modal-content">
					<div class="modal-header">
						<button type="button" class="close" data-dismiss="modal">&times;</button>
						<h4 class="modal-title">Embed CodePen</h4>
					</div>
					<div class="modal-body">
						<p>Please paste the CodePen embed code in here:</p>
						<div class="form-group">
							<input type="text" class="form-control" id="code-embed-input">

						</div>
					</div>
					<div class="modal-footer">
						<button type="submit" id="code-submit-btn" class="btn btn-default" data-dismiss="modal">Submit</button>

						<button type="button" class="btn btn-primary" data-dismiss="modal">Close</button>
					</div>
				</div>
			</div>
		</div>
		
		<div id="widgetBodyCode" class="widgetBody cpEmbeded">
			<!-- this is the default text to show if no link is uploaded
			the actual codepen will need to be sent by ajax after submit
			-->
			<p data-height="265" data-theme-id="dark" data-slug-hash="` + codePenID + `" data-default-tab="result" data-user="short" data-embed-version="2" data-pen-title="campfire" class="codepen">See the Pen <a href="https://codepen.io/short/pen/gGWbQB/">campfire</a> by Short (<a href="https://codepen.io/short">@short</a>) on <a href="https://codepen.io">CodePen</a>.</p>
			<!-- end default text -->
			<script async src="https://production-assets.codepen.io/assets/embed/ei.js"></script>
		</div>
	</div>
</li>
	`
	return output
}
