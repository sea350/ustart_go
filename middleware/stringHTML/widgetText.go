package stringHTML

//WidgetText ... strigifies HTML for a text widget
func WidgetText(text string) string {
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
					Description
				</h4>
			</div>
			<div class="widgetBody">
				<div id="text-box" class="text-box" name="textEditorBody">
					` + text + `
				</div>
			</div>
		</div>
	</li>
	`
}
