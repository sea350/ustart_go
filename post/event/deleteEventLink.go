package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/event"
	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteEventLink ... ADDS A LINK TYPE TO A EVENT'S QUICKLINKS
//Requires event's docID and a type link
//Returns an error
func DeleteEventLink(eclient *elastic.Client, eventID string, link types.Link) error {
	ctx := context.Background()

	GenericEventUpdateLock.Lock()
	defer GenericEventUpdateLock.Unlock()

	evnt, err := get.EventByID(eclient, eventID)
	if err != nil {
		return errors.New("Event does not exist")
	}

	//replace with universal.FindIndex when it works
	index := -1
	for i := range evnt.QuickLinks {
		if evnt.QuickLinks[i] == link {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("link not found")
	}

	evnt.QuickLinks = append(evnt.QuickLinks[:index], evnt.QuickLinks[index+1:]...)

	_, err = eclient.Update().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		Doc(map[string]interface{}{"Quicklinks": evnt.QuickLinks}).
		Do(ctx)

	return err

}
