package uses

import (
	userPost "github.com/sea350/ustart_go/post/user"

	eventPost "github.com/sea350/ustart_go/post/event"
	elastic "gopkg.in/olivere/elastic.v5"
)

//RequestEventMember ...
func RequestEventMember(eclient *elastic.Client, eventID string, userID string) error {
	err := userPost.AppendSentEventReq(eclient, userID, eventID)
	if err != nil {
		return err
	}

	err = eventPost.AppendMemberReqSent(eclient, eventID, userID)
	return err
}
