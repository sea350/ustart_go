package uses

import (
	getFollow "github.com/sea350/ustart_go/get/follow"
	get "github.com/sea350/ustart_go/get/user"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UserPage ... Loads relevant information for the User page
//Requires username and the docID of the person viewing the page
//Returns a typer User, the user's docID, whether or not the viewer is following the person and an error
func UserPage(eclient *elastic.Client, username string, viewerID string) (types.User, string, bool, []types.BadgeProxy, error) {

	var usr types.User

	var isFollowed bool

	userID, err := get.IDByUsername(eclient, username)
	if err != nil {
		return usr, "errors passed 0, username is " + username, isFollowed, nil, err
	}

	usr, err = get.UserByID(eclient, userID)
	if err != nil {
		return usr, "errors passed 1, userId is " + userID, isFollowed, err
	}

	viewer, err := get.UserByID(eclient, viewerID)
	if err != nil {
		return usr, "errors passed 2, viewId is " + viewerID, isFollowed, err
	}

	usrBadges, err := LoadBadgeProxies(eclient, usr.BadgeIDs)
	isFollowed, err = getFollow.IsFollowing(eclient, userID, viewerID, "user")

	for _, element := range viewer.Following {
		if element == userID {
			isFollowed = true
			break
		}
	}

	return usr, "errors passed 3", isFollowed, usrBadges, err
}
