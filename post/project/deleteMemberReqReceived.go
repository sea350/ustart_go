package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/project"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteMemberReqReceived ... DELETES A USER ID IN A PROJECT'S MemberReqReceived ARRAY
//Requires project's docID and the user's docID
//Returns an error
func DeleteMemberReqReceived(eclient *elastic.Client, projectID string, userID string) error {
	ctx := context.Background()
	proj, err := get.ProjectByID(eclient, projectID)
	if err != nil {
		return errors.New("Project does not exist")
	}

	//replace with universal.FindIndex when it works
	index := -1
	for i := range proj.MemberReqReceived {
		if proj.MemberReqReceived[i] == userID {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("link not found")
	}

	proj.MemberReqReceived = append(proj.MemberReqReceived[:index], proj.MemberReqReceived[index+1:]...)

	_, err = eclient.Update().
		Index(globals.ProjectIndex).
		Type(globals.ProjectType).
		Id(projectID).
		Doc(map[string]interface{}{"MemberReqReceived": proj.MemberReqReceived}).
		Do(ctx)

	return err

}
