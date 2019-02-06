package post

import (
	"context"
	"log"

	getFollow "github.com/sea350/ustart_go/get/follow"
	get "github.com/sea350/ustart_go/get/project"
	getUser "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	postChat "github.com/sea350/ustart_go/post/chat"
	postFollow "github.com/sea350/ustart_go/post/follow"
	postUser "github.com/sea350/ustart_go/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//InvisProject ... sets a project to invisible and removes all dependancies
//Requires project docID and a type member
//Returns an error
func InvisProject(eclient *elastic.Client, projID string) error {

	ctx := context.Background()

	//LOCK ALL HERE

	proj, err := get.ProjectByID(eclient, projID)
	if err != nil {
		return err
	}

	//for each member, remove them from all dependancies
	//chat, usr.Projects, etc
	for _, member := range proj.Members {
		usr, err := getUser.UserByID(eclient, member.MemberID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			continue
		}
		var usrIdx int
		for idx := range usr.Projects {
			if usr.Projects[idx].ProjectID == projID {
				usrIdx = idx
				break
			}
		}

		if usrIdx < len(usr.Projects)-1 {
			err = postUser.UpdateUser(eclient, member.MemberID, "Projects", append(usr.Projects[:usrIdx], usr.Projects[usrIdx+1:]...))
		} else {
			err = postUser.UpdateUser(eclient, member.MemberID, "Projects", usr.Projects[:usrIdx])
		}
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}

		for _, subchat := range proj.Subchats {
			err = postChat.RemoveEavesFromConversation(eclient, subchat.ConversationID, member.MemberID)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
		}
	}

	//removing external dependancies here
	_, follow, err := getFollow.ByID(eclient, projID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	} else {

		for followerID := range follow.UserFollowers {
			err = postFollow.RemoveUserFollow(eclient, followerID, "following", projID, "project")
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
		}
	}

	//INSERT REMOVE EVENT DEPENDANCIES HERE

	_, _ = eclient.Update().
		Index(globals.ProjectIndex).
		Type(globals.ProjectType).
		Id(projID).
		Doc(map[string]interface{}{"Visible": false}).
		Do(ctx)

	_, err = eclient.Update().
		Index(globals.ProjectIndex).
		Type(globals.ProjectType).
		Id(projID).
		Doc(map[string]interface{}{"URLName": ""}).
		Do(ctx)

	return err

}
