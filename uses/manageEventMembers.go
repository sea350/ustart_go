package uses

import (
	post "github.com/sea350/ustart_go/post/event"
	"github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//ManageEventMembers ... UPDATES THE FULL MEMBER'S ARRAY
//Requires the target events docID and the potential new url
//Returns an error if the url is taken or a database error
func ManageEventMembers(eclient *elastic.Client, eventID string, newMemberConfig []types.EventMembers) error {
	post.EventMemberLock.Lock()
	defer post.EventMemberLock.Unlock()
	err := post.UpdateEvent(eclient, eventID, "Members", newMemberConfig)
	return err
}
