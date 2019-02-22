package uses

import (
	eventGet "github.com/sea350/ustart_go/get/event"
	projGet "github.com/sea350/ustart_go/get/project"
	eventPost "github.com/sea350/ustart_go/post/event"
	projPost "github.com/sea350/ustart_go/post/project"
	elastic "github.com/olivere/elastic"
)

//ChangeEventHost ... allows a project to host (or even unhost, works both ways)
func ChangeEventHost(eclient *elastic.Client, eventID string, projectID string) error {
	event, err := eventGet.EventByID(eclient, eventID)
	if err != nil {
		return err
	}

	//Change from the project hosting it to the creator
	if event.IsProjectHost {
		projectID = event.Host
		for idx := range event.Members {
			if event.Members[idx].Role == 0 {
				event.Host = event.Members[idx].MemberID
				break
			}
		}
		event.IsProjectHost = false
		proj, err := projGet.ProjectByID(eclient, projectID)
		if err != nil {
			return err
		}

		var eventToRemoveIdx int
		for idx := range proj.EventIDs {
			if proj.EventIDs[idx] == eventID {
				eventToRemoveIdx = idx
				break
			}
		}

		if eventToRemoveIdx < len(proj.EventIDs)-1 {
			err = projPost.UpdateProject(eclient, projectID, "EventIDs", append(proj.EventIDs[:eventToRemoveIdx], proj.EventIDs[eventToRemoveIdx+1:]...))
		}
		if err != nil {
			return err
		}

		if eventToRemoveIdx == len(proj.EventIDs) {
			err = projPost.UpdateProject(eclient, projectID, "EventIDs", nil)
		}
		if err != nil {
			return err
		}

		return eventPost.UpdateEvent(eclient, eventID, "Host", event.Host)
	}

	//If the if statement did not reach then we do the opposite, set the Host from the creator to the project
	event.Host = projectID
	event.IsProjectHost = true
	err = eventPost.UpdateEvent(eclient, eventID, "Host", event.Host)
	if err != nil {
		return err
	}

	return eventPost.UpdateEvent(eclient, eventID, "IsProjectHost", event.IsProjectHost)

}
