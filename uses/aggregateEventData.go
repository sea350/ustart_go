package uses

import (
	"log"
	"os"
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
	
	fmt.Println("1")
	data, err := getEvent.EventByURL(eclient, url)
	fmt.Println("2")
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	fmt.Println("3")
	eventData.EventData = data

	fmt.Println("4")
	id, err := getEvent.EventIDByURL(eclient, url)
	fmt.Println("5")
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	fmt.Println("6")
	eventData.DocID = id

	fmt.Println("7")
	//Remember to load widgets seperately
	//Remember to load wall posts seperately

	for _, member := range data.Members {
		id := member.MemberID
		mem, err := ConvertUserToFloatingHead(eclient, id)
		fmt.Println("8")
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
	fmt.Println("9")
	for _, receivedReq := range eventData.EventData.MemberReqReceived {
		fmt.Println("10")
		if receivedReq == viewerID {
			eventData.RequestAllowed = false
		}
	}
	fmt.Println("11")

	return eventData, err
}
