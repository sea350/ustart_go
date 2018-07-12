package uses

import (
	getEvnt "github.com/sea350/ustart_go/get/event"
	getUser "github.com/sea350/ustart_go/get/user"
	postEvnt "github.com/sea350/ustart_go/post/event"
	postUser "github.com/sea350/ustart_go/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UserFollowEventToggle ... if a user wants to follow or unfollow an event, use this function to toggle
//Note, once error is taken care of function should match follow status on both sides automatically
func UserFollowEventToggle(eclient *elastic.Client, userID string, eventID string) error {

	//first check if user is already followed
	isFollowed := false

	usr, err := getUser.UserByID(eclient, userID)
	if err != nil {
		return err
	}

	postEvnt.EventFollowerLock.Lock() //make sure no one else changes followers array while we read/write it
	defer postEvnt.EventFollowerLock.Unlock()

	evnt, err := getEvnt.EventByID(eclient, eventID)
	if err != nil {
		return err
	}

	//since we already locate the followed event, remove it so we dont have to search again
	for i, id := range usr.FollowingEvent {
		if id == eventID {
			isFollowed = true

			//check if we are removing the last item
			if i != len(usr.FollowingEvent)-1 { //if no, popout
				usr.FollowingEvent = append(usr.FollowingEvent[:i], usr.FollowingEvent[i+1:]...)
			} else { //if yes the popback
				usr.FollowingEvent = usr.FollowingEvent[:i]
			}
			break
		}
	}

	//if the user is followed, unfollow them
	if isFollowed {
		//update user
		err := postUser.UpdateUser(eclient, userID, "FollowingEvent", usr.FollowingEvent)
		if err != nil {
			return err
		}

		//update event

		for i, id := range evnt.FollowedUsers {
			if id == userID {
				//check if we are removing the last item
				if i != len(evnt.FollowedUsers)-1 { //if no, popout
					evnt.FollowedUsers = append(evnt.FollowedUsers[:i], evnt.FollowedUsers[i+1:]...)
				} else { //if yes the popback
					evnt.FollowedUsers = evnt.FollowedUsers[:i]
				}
				break
			}
		}

		err = postEvnt.UpdateEvent(eclient, eventID, "FollowedUsers", evnt.FollowedUsers)
		if err != nil {
			return err
		}
	} else { //if the user is not followed, follow them
		//update user with new event id
		err = postUser.UpdateUser(eclient, userID, "FollowingEvent", append(usr.FollowingEvent, eventID))
		if err != nil {
			return err
		}

		//update event with new follower id
		err = postEvnt.UpdateEvent(eclient, eventID, "FollowedUsers", append(evnt.FollowedUsers, userID))
		if err != nil {
			return err
		}

	}

	return nil
}
