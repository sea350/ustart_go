package stringHTML

//WidgetText ... strigifies HTML for a text widget
func WidgetText(title string, description string) string {
	return `
	<li class="ui-state-default">
		<div class="projectsWidgetCont">
			<div class="widgetTitle">
				<span class="pull-right fa fa-2x fa-sort"></span>
				<script>
					function focusDesc() {
						document.getElementById('text-box').focus();
					}
				</script>
				<span class="pull-right fa fa-2x fa-pencil" onclick="focusDesc()"></span>
				<h4 id="text-title" name="textEditorHeader">
					` + title + `
				</h4>
			</div>
			<div class="widgetBody">
				<div id="text-box" class="text-box" name="textEditorBody">
					` + description + `
				</div>
			</div>
		</div>
	</li>
	`
}
