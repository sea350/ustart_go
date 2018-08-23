package uses

import (
	post "github.com/sea350/ustart_go/post/event"
	elastic "gopkg.in/olivere/elastic.v5"
	"time"
)

//ChangeEventNameAndDescription ... CHANGES BOTH A EVENT NAME AND DESCRIPTION
//Requires the target event's docID, new name and description
//Returns an error if there was a problem with database submission
//NOTE: it is possible Name change goes through but not Description
func ChangeEventTime(eclient *elastic.Client, eventID string, Sdate time.Time, Edate time.Time) error {
	err := post.UpdateEvent(eclient, eventID, "EventDateStart", Sdate)
	if err != nil {
		return err
	}
	err = post.UpdateEvent(eclient, eventID, "EventDateEnd", Edate)
	return err
}
