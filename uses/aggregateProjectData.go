package uses

import (
	get "github.com/sea350/ustart_go/get/project"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AggregateProjectData ...
//Adds a new widget to the UserWidgets array
func AggregateProjectData(eclient *elastic.Client, url string) (types.ProjectAggregate, error) {
	var projectData types.ProjectAggregate

	data, err := get.ProjectByURL(eclient, url)
	if err != nil {
		panic(err)
	}
	projectData.ProjectData = data

	id, err := get.ProjectIDByURL(eclient, url)
	if err != nil {
		panic(err)
	}
	projectData.DocID = id

	//Remember to load widgets seperately
	//Remember to load wall posts seperately

	for _, member := range data.Members {
		id := member.MemberID
		mem, err := ConvertToFloatingHead(eclient, id)
		if err != nil {
			panic(err)
		}
		projectData.MemberData = append(projectData.MemberData, mem)
	}

	return projectData, err
}
