package uses

import (
	getProject "github.com/sea350/ustart_go/get/project"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AggregateProjectData ...
//Adds a new widget to the UserWidgets array
func AggregateProjectData(eclient *elastic.Client, url string) (types.ProjectAggregate, error) {
	var projectData types.ProjectAggregate

	data, err := getProject.ByURL(eclient, url)
	if err != nil {
		panic(err)
	}
	projectData.ProjectData = data

	id, err := getProject.IDByURL(eclient, url)
	if err != nil {
		panic(err)
	}
	projectData.DocID = id

	//Remember to load widgets seperately

	return projectData, err

}
