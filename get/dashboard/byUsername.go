package get

import (
	get "github.com/sea350/ustart_go/get/user"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DashboardByUsername ...
func DashboardByUsername(eclient *elastic.Client, username string) (types.Dashboard, error) {

	var newDash types.Dashboard

	usr, err := get.UserByUsername(eclient, username)
	if err != nil {
		return newDash, err
	}

	id, err := get.IDByUsername(eclient, username)
	if err != nil {
		return newDash, err
	}
	newDash.ID = id
	newDash.Followers = usr.Followers
	newDash.FollowingProject = usr.FollowingProject
	newDash.FollowingEvent = usr.FollowingEvent

	return newDash, err

}
