package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/project"
	getUser "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	postChat "github.com/sea350/ustart_go/post/chat"
	postUser "github.com/sea350/ustart_go/post/user"
	elastic "github.com/olivere/elastic"
)

//DeleteMember ... REMOVES A SPECIFIC MEMBER FROM AN ARRAY
//Requires project docID and a type member
//Returns an error
func DeleteMember(eclient *elastic.Client, projID string, userID string) error {

	ctx := context.Background()

	ModifyMemberLock.Lock()
	defer ModifyMemberLock.Unlock()

	usr, err := getUser.UserByID(eclient, userID)
	if err != nil {
		return err
	}
	proj, projErr := get.ProjectByID(eclient, projID)
	if err != nil {
		return projErr
	}

	var usrIdx = -1
	for idx := range usr.Projects {
		if usr.Projects[idx].ProjectID == projID {
			usrIdx = idx
			break
		}
	}

	if usrIdx > len(usr.Projects) {
		return errors.New("Index out of bounds")
	}
	if usrIdx < 0 {
		return errors.New("Project not found")
	}

	if usrIdx < len(usr.Projects)-1 {
		updatedProjects := append(usr.Projects[:usrIdx], usr.Projects[usrIdx+1:]...)
		err = postUser.UpdateUser(eclient, userID, "Projects", updatedProjects)

		// _, err = eclient.Update().
		// 	Index(globals.UserIndex).
		// 	Type(globals.UserType).
		// 	Id(projID).
		// 	Doc(map[string]interface{}{"Projects": updatedProjects}).
		// 	Do(ctx)

		if err != nil {
			return err
		}
	} else if usrIdx == len(usr.Projects)-1 {
		err = postUser.UpdateUser(eclient, userID, "Projects", usr.Projects[:usrIdx])

	} else if usrIdx == len(usr.Projects) {
		err = postUser.UpdateUser(eclient, userID, "Projects", nil)
		// _, err = eclient.Update().
		// 	Index(globals.UserIndex).
		// 	Type(globals.UserType).
		// 	Id(projID).
		// 	Doc(map[string]interface{}{"Projects": nil}).
		// 	Do(ctx)
		if err != nil {
			return err
		}
	}

	index := -1
	for i := range proj.Members {
		if proj.Members[i].MemberID == userID {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("Member not found")
	}

	proj.Members = append(proj.Members[:index], proj.Members[index+1:]...)

	for _, subchat := range proj.Subchats {
		err = postChat.RemoveEavesFromConversation(eclient, subchat.ConversationID, userID)
		if err != nil {
			return err
		}
	}

	_, err = eclient.Update().
		Index(globals.ProjectIndex).
		Type(globals.ProjectType).
		Id(projID).
		Doc(map[string]interface{}{"Members": proj.Members}).
		Do(ctx)

	return err

}
