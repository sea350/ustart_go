package uses

import (
	get "github.com/sea350/ustart_go/get/user"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UserPage ... Loads relevant information for the User page
//Requires username and the docID of the person viewing the page
//Returns a typer User, the user's docID, whether or not the viewer is following the person and an error
func UserPage(eclient *elastic.Client, username string, viewerID string) (types.User, string, bool, error) {

	var usr types.User

	var isFollowed bool

	userID, err := get.IDByUsername(eclient, username)
	if err != nil {
		return usr, userID, isFollowed, err
	}

	usr, err = get.UserByID(eclient, userID)
	if err != nil {
		return usr, userID, isFollowed, err
	}

	viewer, err := get.UserByID(eclient, viewerID)
	if err != nil {
		return usr, userID, isFollowed, err
	}

	for _, element := range viewer.Following {
		if element == userID {
			isFollowed = true
			break
		}
	}

	return usr, userID, isFollowed, err
}
