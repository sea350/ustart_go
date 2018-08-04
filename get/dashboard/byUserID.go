package get

import (
	get "github.com/sea350/ustart_go/get/user"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DashboardByUserID ...
func DashboardByUserID(eclient *elastic.Client, usrID string) (types.Dashboard, error) {

	var newDash types.Dashboard

	usr, err := get.UserByID(eclient, usrID)
	if err != nil {
		return newDash, err
	}

	newDash.ID = usrID
	newDash.Followers = usr.Followers
	newDash.FollowingProject = usr.FollowingProject
	newDash.FollowingEvent = usr.FollowingEvent
	newDash.EntryIDs = usr.EntryIDs

	return newDash, err

}
