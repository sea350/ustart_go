package uses

import (
	"log"

	getProject "github.com/sea350/ustart_go/get/project"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AggregateProjectData ...
//Adds a new widget to the UserWidgets array
func AggregateProjectData(eclient *elastic.Client, url string, viewerID string) (types.ProjectAggregate, error) {
	var projectData types.ProjectAggregate
	projectData.RequestAllowed = true

	data, err := getProject.ProjectByURL(eclient, url)
	if err != nil {
		log.Println("Error: uses/aggregateProjectData line 18")
		log.Println(err)
	}
	projectData.ProjectData = data

	id, err := getProject.ProjectIDByURL(eclient, url)
	if err != nil {
		log.Println("Error: uses/aggregateProjectData line 25")
		log.Println(err)
	}
	projectData.DocID = id

	//Remember to load widgets seperately
	//Remember to load wall posts seperately

	for _, member := range data.Members {
		id := member.MemberID
		mem, err := ConvertUserToFloatingHead(eclient, id)
		if err != nil {
			log.Println("Error: uses/aggregateProjectData line 37")
			log.Println(err)
		}
		mem.Classification = member.Role
		projectData.MemberData = append(projectData.MemberData, mem)
		if viewerID == member.MemberID {
			projectData.RequestAllowed = false
		}
	}
	for _, receivedReq := range projectData.ProjectData.MemberReqReceived {
		if receivedReq == viewerID {
			projectData.RequestAllowed = false
		}
	}

	return projectData, err
}
