package get

import (
	"errors"

	types "github.com/sea350/ustart_go/backend/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//EventAndMember ...
func EventAndMember(eclient *elastic.Client, eventID string, memberID string) (types.Events, types.EventMembers, error) {
	evnt, err := EventByID(eclient, eventID)
	if err != nil {
		return types.Events{}, types.EventMembers{}, err
	}

	if len(evnt.Members) < 1 {
		return types.Events{}, types.EventMembers{}, errors.New("Event has zero members")
	}

	var retMember types.EventMembers

	for mem := range evnt.Members {
		if evnt.Members[mem].MemberID == memberID {
			retMember = evnt.Members[mem]
		}
	}

	return evnt, retMember, err
}
