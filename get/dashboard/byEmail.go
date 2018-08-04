package get

import (
	get "github.com/sea350/ustart_go/get/user"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DashboardByEmail ...
func DashboardByEmail(eclient *elastic.Client, email string) (types.Dashboard, error) {

	var newDash types.Dashboard

	usr, err := get.UserByEmail(eclient, email)
	if err != nil {
		return newDash, err
	}

	id, err := get.UserIDByEmail(eclient, email)
	if err != nil {
		return newDash, err
	}
	newDash.ID = id
	newDash.Followers = usr.Followers
	newDash.FollowingProject = usr.FollowingProject
	newDash.FollowingEvent = usr.FollowingEvent
	newDash.EntryIDs = usr.EntryIDs

	return newDash, err

}
