package uses

import (
	"log"

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
		log.Println(err)
	}
	eventData.EventData = data

	id, err := getEvent.EventIDByURL(eclient, url)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	eventData.DocID = id

	//Remember to load widgets seperately
	//Remember to load wall posts seperately

	for _, member := range data.Members {
		id := member.MemberID
		mem, err := ConvertUserToFloatingHead(eclient, id)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		mem.Classification = member.Role
		eventData.MemberData = append(eventData.MemberData, mem)

		if viewerID == member.MemberID {
			eventData.RequestAllowed = false
		}
	}
	for _, guest := range data.Guests {
		id := guest.GuestID
		guest, err := ConvertUserToFloatingHead(eclient, id)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		//guest.Classification = guest.Role
		eventData.GuestData = append(eventData.GuestData, guest)
		if viewerID == id {
			eventData.RequestAllowed = false
		}

	}

	for _, project := range data.Projects {
		id := project.ProjectID
		proj, err := ConvertProjectToFloatingHead(eclient, id)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		eventData.ProjectData = append(eventData.ProjectData, proj)
	}

	if eventData.RequestAllowed {
		for guestID := range data.GuestReqReceived {
			if guestID == viewerID {
				eventData.RequestAllowed = false
				break
			}
		}
	}

	if eventData.RequestAllowed {
		for _, memReq := range data.MemberReqReceived {
			if memReq == viewerID {
				eventData.RequestAllowed = false
				break
			}
		}
	}

	return eventData, err
}
