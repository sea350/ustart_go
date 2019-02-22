package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/project"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//AppendMember ... APPENDS A NEW MEMBER TO AN EXISTING PROJECT DOC
//Requires the project docID and a type Member
//Returns an error
func AppendMember(eclient *elastic.Client, projectID string, member types.Member) error {

	ctx := context.Background()

	ModifyMemberLock.Lock()
	defer ModifyMemberLock.Unlock()

	proj, err := get.ProjectByID(eclient, projectID)
	if err != nil {
		return errors.New("Project does not exist")
	}

	proj.Members = append(proj.Members, member)

	_, err = eclient.Update().
		Index(globals.ProjectIndex).
		Type(globals.ProjectType).
		Id(projectID).
		Doc(map[string]interface{}{"Members": proj.Members}).
		Do(ctx)

	return err

}
