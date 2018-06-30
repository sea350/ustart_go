package uses

import (
	"fmt"

	getEvent "github.com/sea350/ustart_go/get/event"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AggregateEventData ...
//Adds a new widget to the UserWidgets array
func AggregateEventData(eclient *elastic.Client, url string, viewerID string) (types.EventAggregate, error) {
	var eventData types.EventAggregate
	eventData.RequestAllowed = true

	data, err := getEvent.EventByURL(eclient, url)
	if err != nil {
		fmt.Println(err)
		fmt.Println("error aggregateEventData.go 17")
	}
	eventData.EventData = data

	id, err := getEvent.EventIDByURL(eclient, url)
	if err != nil {
		fmt.Println(err)
		fmt.Println("error aggregateEventData.go 24")
	}
	eventData.DocID = id

	//Remember to load widgets seperately
	//Remember to load wall posts seperately

	for _, member := range data.Members {
		id := member.MemberID
		mem, err := ConvertUserToFloatingHead(eclient, id)
		if err != nil {
			fmt.Println(err)
			fmt.Println("error aggregateEventData.go 36")
		}
		mem.Classification = member.Role
		eventData.MemberData = append(eventData.MemberData, mem)
		if viewerID == member.MemberID {
			eventData.RequestAllowed = false
		}
	}
	for _, receivedReq := range eventData.EventData.MemberReqReceived {
		if receivedReq == viewerID {
			eventData.RequestAllowed = false
		}
	}

	return eventData, err
}
