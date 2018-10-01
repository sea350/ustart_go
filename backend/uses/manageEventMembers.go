package uses

import (
	post "github.com/sea350/ustart_go/backend/post/event"
	"github.com/sea350/ustart_go/backend/types"
	elastic "gopkg.in/olivere/elastic.v5"
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
