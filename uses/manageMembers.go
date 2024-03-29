package uses

import (
	post "github.com/sea350/ustart_go/post/project"
	"github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
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
