package uses

import (
	post "github.com/sea350/ustart_go/post/event"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChangeEventCategory ... CHANGES EVENT THE EVENT'S CATEGORY
//Requires the target event's docID, all aspects of a types.LocStruct
//Returns an error if there was a problem with database submission
func ChangeEventCategory(eclient *elastic.Client, eventID string, category string) error {
	err := post.UpdateEvent(eclient, eventID, "Category", category)
	return err
}
