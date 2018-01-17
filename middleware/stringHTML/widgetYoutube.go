package stringHTML

//WidgetYoutube ... strigifies HTML for a youtube widget
func WidgetYoutube(embedcode string) string {
	output := `
<li class="ui-state-default widgetListItem sortable">
	<div class="projectsWidgetCont">
		<div class="widgetTitle">
			<span class="pull-right fa fa-2x fa-sort"></span>
			<span class="pull-right fa fa-2x fa-trash"></span>
			<span class="pull-right fa fa-2x fa-pencil" code="` + embedcode + `"></span>
			<h4>Youtube</h4>
		</div>
		<div class="widgetBody">
			<div class="youtube-feed">
				<!-- don't actually copy this line below, instead pull the iframe code from database instead -->
				<iframe src="https://www.youtube.com/embed/` + embedcode + `" frameborder="0" allowfullscreen></iframe>
			</div>
		</div>
	</div>
</li>
	`
	return output
}
