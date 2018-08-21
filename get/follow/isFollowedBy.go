package get

import (
	"errors"
	"log"
	"strings"

	elastic "gopkg.in/olivere/elastic.v5"
	//post "github.com/sea350/ustart_go/post"
)

//IsFollowedBy ...
//Determines if specific doc id is being followed
func IsFollowedBy(eclient *elastic.Client, userID string, followID string, followType string) (bool, error) {
	_, follows, err := ByID(eclient, userID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return false, err
	}

	var exists bool
	switch strings.ToLower(followType) {
	case "user":
		_, exists = follows.UserFollowers[followID]
	case "project":
		_, exists = follows.ProjectFollowers[followID]
	case "event":
		_, exists = follows.EventFollowers[followID]
	default:
		return false, errors.New("Invalid field")
	}

	return exists, err
}
