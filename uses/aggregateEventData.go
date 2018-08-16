package uses

import (
	"log"
	"os"

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
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	eventData.EventData = data
	eventData.DocID = url

	//Remember to load widgets seperately
	//Remember to load wall posts seperately

	for _, member := range data.Members {
		id := member.MemberID
		mem, err := ConvertUserToFloatingHead(eclient, id)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
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
