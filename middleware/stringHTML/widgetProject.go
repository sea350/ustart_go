package stringHTML

import "github.com/sea350/ustart_go/types"

//WidgetProject ... strigifies HTML for a text widget WARNING WIP need to update projectLink
func WidgetProject(data []types.ProjectSubWidgets) string {
	projectLink := "https://media.discordapp.net/attachments/281949538506506241/402950154400956424/wtf.jpg?width=702&height=468"
	output := `
<li class="ui-state-default">
    <div id="projectWidget" class="projectsWidgetCont">
        <div class="widgetTitle">
            <span class="pull-right fa fa-2x fa-sort"></span>
            <span class="pull-right fa fa-2x fa-pencil" data-toggle="modal" data-target="#projModal"></span>
            <h4>Projects</h4>
        </div>
        <div class="widgetBody">
			<!--- THIS IS ONE INSTANCE OF THE PROJECT probably want a for loop here -->`

	for _, subWidget := range data {
		output += `
            <button class="btn btn-default projectsColumn" id="project` + subWidget.ID + `">
                <div class="columnImage">
                    <a href="` + projectLink + `">
						<img src=" ` + subWidget.Avatar + `" alt="Featured Project Image" />
                        <div class="` + subWidget.Name + `">
							` + subWidget.Name + `
                        </div>
                    </a>
                </div>
            </button>
			<!-- END for loop here -->`
	}

	output += `
        </div>
    </div>
    <div class="modal fade" id="projModal" role="dialog">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal">&times;</button>
                    <h3>Projects</h3>
                </div>
                <div class="modal-body">
                    List of projects (sortable) + toggle to show each individual proj
                    <ul id="projSortable">
                    </ul>
                </div>
            </div>
        </div>
    </div>
</li>
	`
	return output
}
