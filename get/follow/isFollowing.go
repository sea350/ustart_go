package get

import (
	"log"

	elastic "gopkg.in/olivere/elastic.v5"
	//post "github.com/sea350/ustart_go/post"
)

//IsFollowing ...
//Determines if specific doc id is being followed
func IsFollowing(eclient *elastic.Client, userID string, followID string, followType int) (bool, error) {
	_, follows, err := ByUserID(eclient, userID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return false, err
	}

	var exists bool
	switch followType {
	case 1:
		_, exists = follows.UserFollowing[followID]
	case 2:
		_, exists = follows.ProjectFollowing[followID]
	case 3:
		_, exists = follows.EventFollowing[followID]
	}

	return exists, err
}
