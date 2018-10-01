package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/backend/get/event"
	globals "github.com/sea350/ustart_go/backend/globals"
	types "github.com/sea350/ustart_go/backend/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendEventLink ... ADDS A LINK TYPE TO A Event'S QUICKLINKS
//Requires Event's docID and a type link
//Returns an error
func AppendEventLink(eclient *elastic.Client, eventID string, link types.Link) error {
	ctx := context.Background()

	evnt, err := get.EventByID(eclient, eventID)
	if err != nil {
		return errors.New("Event does not exist")
	} //return error if nonexistent

	evnt.QuickLinks = append(evnt.QuickLinks, link) //retreive Event

	_, err = eclient.Update().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		Doc(map[string]interface{}{"QuickLinks": evnt.QuickLinks}). //update Event doc with new link appended to array
		Do(ctx)

	return err

}
