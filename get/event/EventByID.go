package get

import (
	"context"
	"encoding/json"
	"fmt"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//EventByID ... Pulls from ES an Event (and an error)
func EventByID(eclient *elastic.Client, eventID string) (types.Events, error) {
	var evnt types.Events

	ctx := context.Background()

	log.Println("Event ID: " + eventID)

	searchResult, err := eclient.Get().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		Do(ctx)
	if err != nil {
		fmt.Printf("Error From EventByID.SearchResult.Get(): %s\n", err.Error())
	}
	
	err = json.Unmarshal(*searchResult.Source, &evnt)
	if err != nil {
		fmt.Printf("Err: %s\n", err.Error())
	}

	return evnt, err
}
