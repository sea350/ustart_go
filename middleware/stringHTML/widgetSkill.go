package stringHTML

//WidgetSkill ... strigifies HTML for a skill widget
func WidgetSkill(skills []string) string {
	output := `
	<li class="ui-state-default">
	<div class="projectsWidgetCont">
		<div class="widgetTitle">
			<span class="pull-right fa fa-2x fa-sort"></span>
			<span class="pull-right fa fa-2x fa-pencil" data-toggle="modal" data-target="#tagModal"></span>
			<h4>Skills</h4>
		</div>
		<div id="hashTagsBody" class="widgetBody">
			<!--- THIS IS ONE INSTANCE OF SKILL probably want a for loop here -->`
	for _, skill := range skills {
		output += `
		<button class="btn btn-default projectsColumn" id="skill-` + skill + `">
			<div class="columnImage">
				<div class="columnTitle">` + skill + `</div>
			</div>
		</button>
		<!-- END for loop here -->
		`
	}

	output += `
		</div>
	</div>
</li>
	`
	return output
}
