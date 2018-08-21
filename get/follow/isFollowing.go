package get

import (
	"errors"
	"log"
	"strings"

	elastic "gopkg.in/olivere/elastic.v5"
)

//post "github.com/sea350/ustart_go/post"

//IsFollowing ...
//Determines if specific doc id is being followed

//GP COMMENT cannot use int AS STRING (followID)
func IsFollowing(eclient *elastic.Client, userID string, followID int, followType string) (bool, error) {
	_, follows, err := ByID(eclient, userID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return false, err
	}

	var exists bool
	switch strings.ToLower(followType) {
	case "user":
		_, exists = follows.UserFollowing[followID]
	case "project":
		_, exists = follows.ProjectFollowing[followID]
	case "event":
		_, exists = follows.EventFollowing[followID]
	default:
		return false, errors.New("Invalid field")
	}

	return exists, err
}
