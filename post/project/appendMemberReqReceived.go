package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/project"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendMemberReqReceived ... APPENDS A USER ID TO A PROJECTS MemberReqReceived ARRAY
//Requires project's docID and the user's docID
//Returns an error
func AppendMemberReqReceived(eclient *elastic.Client, projectID string, userID string) error {

	ctx := context.Background()

	proj, err := get.ProjectByID(eclient, projectID)
	if err != nil {
		return errors.New("Project does not exist")
	} //return error if nonexistent

	proj.MemberReqReceived = append(proj.MemberReqReceived, userID) //retreive project

	_, err = eclient.Update().
		Index(globals.ProjectIndex).
		Type(globals.ProjectType).
		Id(projectID).
		Doc(map[string]interface{}{"MemberReqReceived": proj.MemberReqReceived}). //update project doc with new link appended to array
		Do(ctx)

	return err
}
