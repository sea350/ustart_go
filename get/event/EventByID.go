package get

import (
	"context"
	"encoding/json"
	"log"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//EventByID ... Pulls from ES an Event (and an error)
func EventByID(eclient *elastic.Client, eventID string) (types.Events, error) {
	var evnt types.Events

	ctx := context.Background()

	searchResult, err := eclient.Get().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		Do(ctx)

	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return evnt, err
	}

	err = json.Unmarshal(*searchResult.Source, &evnt)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return evnt, err
	}

	if len(evnt.GuestReqSent) == 0 {
		evnt.GuestReqSent = make(map[string]int)
	}
	if len(evnt.GuestReqReceived) == 0 {
		evnt.GuestReqReceived = make(map[string]int)
	}
	return evnt, err
}
