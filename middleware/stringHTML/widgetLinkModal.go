package stringHTML

//WidgetLinkModal ... strigifies HTML for links WARNING INCOMPLETE
func WidgetLinkModal(skills []string) string {
	output := `
<li class="ui-state-default widgetListItem sortable">
	<div id="linksWidget" class="projectsWidgetCont">
		<div class="widgetTitle">
			<span class="pull-right fa fa-2x fa-sort"></span>
			<span class="pull-right fa fa-2x fa-pencil" data-toggle="modal" data-target="#addLinkModal"></span>
			<h4>Links</h4>
		</div>
		<div class="widgetBody">
			<div class="links-container"></div>
			<div class="numLinks">
				<span id="linkCountIndicator">16 Links Remaining</span>
			</div>
		</div>

		<!-- Add Link Modal -->
		<div class="modal fade" id="addLinkModal" role="dialog">
			<div class="modal-dialog">
				<!-- Modal content-->
				<div class="modal-content">
					<div class="modal-header">
						<button type="button" class="close" data-dismiss="modal">&times;</button>
						<h4 class="modal-title">Add Link</h4>
					</div>
					<form id="addLinkForm">
						<div class="modal-body">
							<p>
								<span class="modal-cell">Title</span>
								<input type="text" class="link-input-read" name="webTitle" placeholder="example: 'My LinkedIn'" spellcheck="false" required autofocus />
							</p>
							<p>
								<span class="modal-cell">URL</span>
								<input type="url" class="link-input-read" name="webURL" placeholder="example: 'https://www.linkedin.com'" spellcheck="false" required/>
							</p>
						</div>
						<div class="modal-footer">
							<div class="btn-group">
								<input type="submit" class="btn btn-primary btn-add-link" />
								<button type="button" class="btn btn-info" data-dismiss="modal">Close</button>
							</div>
						</div>
					</form>
				</div>
			</div>
		</div>
		<!-- End of add link Modal -->
		<script src="js/jquery.validate.min.js"></script>
		<script src="js/layout_links.js"></script>
	</div>
</li>
	`
	return output
}
