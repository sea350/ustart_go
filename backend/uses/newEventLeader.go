package uses

import (
	get "github.com/sea350/ustart_go/backend/get/event"
	post "github.com/sea350/ustart_go/backend/post/event"
	elastic "gopkg.in/olivere/elastic.v5"
)

//NewEventLeader ...
//ENables current leader to set a new leader
func NewEventLeader(eclient *elastic.Client, eventID string, currentLeaderID string, newLeaderID string) error {
	evnt, err := get.EventByID(eclient, eventID)
	if err != nil {
		panic(err)
	}
	isLeader, idx := IsEventLeader(eclient, eventID, currentLeaderID)

	for i := range evnt.Members {
		if evnt.Members[i].MemberID == newLeaderID && isLeader {
			evnt.Members[i].Role = 0
			evnt.Members[idx].Role = 1 //can be any role value
		}
	}

	updateErr := post.UpdateEvent(eclient, eventID, "Members", evnt.Members)
	return updateErr

}
