package uses

import (
	post "github.com/sea350/ustart_go/backend/post/project"
	"github.com/sea350/ustart_go/backend/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ManageMembers ... UPDATES THE FULL MEMBER'S ARRAY
//Requires the target projects docID and the potential new url
//Returns an error if the url is taken or a database error
func ManageMembers(eclient *elastic.Client, projectID string, newMemberConfig []types.Member) error {
	post.ModifyMemberLock.Lock()
	defer post.ModifyMemberLock.Unlock()
	err := post.UpdateProject(eclient, projectID, "Members", newMemberConfig)
	return err
}
