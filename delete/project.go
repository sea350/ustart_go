package delete

import (
	"log"

	getFollow "github.com/sea350/ustart_go/get/follow"
	getProject "github.com/sea350/ustart_go/get/project"
	post "github.com/sea350/ustart_go/post/event"
	postFollow "github.com/sea350/ustart_go/post/follow"
	elastic "github.com/olivere/elastic"
)

//Project ... Removes all traces of the project
//Requires projectID
//Returns an error
func Project(eclient *elastic.Client, projectID string) error {
	//ctx := context.Background()

	project, err := getProject.ProjectByID(eclient, projectID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	//Delete members from project and project from members
	for _, element := range project.Members {
		err = post.DeleteMember(eclient, projectID, element.MemberID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
	}

	//Make entries and child entries invisible
	//arrayEntryIDs, err := getEntry.EntryIDsByProjectID(eclient, projectID)

	//Delete Followers who are users
	_, follow, err := getFollow.ByID(eclient, projectID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	for key := range follow.UserFollowers {
		err = postFollow.RemoveUserFollow(eclient, key, "following", projectID, "project")
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
	}
	return err
}
