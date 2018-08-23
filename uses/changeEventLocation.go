package uses

import (
	post "github.com/sea350/ustart_go/post/event"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChangeEventLocation ... CHANGES EVENT THE EVENT'S LISTED LOCATION
//Requires the atarget event's docID all aspects of a types.LocStruct
//Returns an error if there was a problem with database submission
func ChangeEventLocation(eclient *elastic.Client, eventID string, country string, state string, city string, street string, zip string) error {
	var newLoc types.LocStruct
	newLoc.Country = country
	newLoc.State = state
	newLoc.City = city
	newLoc.Street = street
	newLoc.Zip = zip

	err := post.UpdateEvent(eclient, eventID, "Location", newLoc)
	return err

}
