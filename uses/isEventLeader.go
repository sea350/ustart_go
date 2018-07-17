package uses

import (
	get "github.com/sea350/ustart_go/get/event"
	elastic "gopkg.in/olivere/elastic.v5"
)

//IsEventLeader ...
//Returns event to represent whether member is event leader
//ALso returns index of member
func IsEventLeader(eclient *elastic.Client, eventID string, memberID string) (bool, int) {
	evnt, err := get.EventByID(eclient, eventID)
	if err != nil {
		panic(err)
	}
	idx := -1
	for i := range evnt.Members {
		//idx = i
		if evnt.Members[i].MemberID == memberID && evnt.Members[i].Role == 0 {
			idx = i
			return true, idx
		}
	}

	return false, idx
}
