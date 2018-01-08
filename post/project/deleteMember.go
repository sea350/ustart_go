package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/project"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteMember ... REMOVES A SPECIFIC MEMBER FROM AN ARRAY
//Requires project docID and a type member
//Returns an error
func DeleteMember(eclient *elastic.Client, projID string, member types.Member) error {

	ctx := context.Background()

	ModifyMemberLock.Lock()

	proj, err := get.ProjectByID(eclient, projID)
	if err != nil {
		return errors.New("Project does not exist")
	}

	index := -1
	for i := range proj.Members {
		if proj.Members[i] == member {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("Member not found")
	}

	proj.Members = append(proj.Members[:index], proj.Members[index+1:]...)

	_, err = eclient.Update().
		Index(globals.ProjectIndex).
		Type(globals.ProjectType).
		Id(projID).
		Doc(map[string]interface{}{"Members": proj.Members}).
		Do(ctx)

	defer ModifyMemberLock.Unlock()
	return err

}
