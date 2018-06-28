package get

import (
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//isEventMember ...
func isEventMember(eclient *elastic.Client, memberID string, event types.Events) bool {
	if len(event.Members) < 1 {
		return false
	}

	for mem, _ := range event.Members {
		if event.Members[mem].MemberID == memberID {
			return true
		}
	}

	return false

}
