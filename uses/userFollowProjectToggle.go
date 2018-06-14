package uses

import (
	getProj "github.com/sea350/ustart_go/get/project"
	getUser "github.com/sea350/ustart_go/get/user"
	postProj "github.com/sea350/ustart_go/post/project"
	postUser "github.com/sea350/ustart_go/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UserFollowProjectToggle ... if a user wants to follow or unfollow a project, use this function to toggle
//Note, once error is taken care of function should match follow status on both sides automatically
func UserFollowProjectToggle(eclient *elastic.Client, userID string, projectID string) error {

	//first check if user is already followed
	isFollowed := false

	usr, err := getUser.UserByID(eclient, userID)
	if err != nil {
		return err
	}

	postProj.FollowerLock.Lock() //make sure no one else changes followers array while we read/write it
	defer postProj.FollowerLock.Unlock()

	proj, err := getProj.ProjectByID(eclient, projectID)
	if err != nil {
		return err
	}

	//since we already locate the followed project, remove it so we dont have to search again
	for i, id := range usr.FollowingProject {
		if id == projectID {
			isFollowed = true

			//check if we are removing the last item
			if i != len(usr.FollowingProject)-1 { //if no, popout
				usr.FollowingProject = append(usr.FollowingProject[:i], usr.FollowingProject[i+1:]...)
			} else { //if yes the popback
				usr.FollowingProject = usr.FollowingProject[:i]
			}
			break
		}
	}

	//if the user is followed, unfollow them
	if isFollowed {
		//update user
		err := postUser.UpdateUser(eclient, userID, "FollowingProject", usr.FollowingProject)
		if err != nil {
			return err
		}

		//update project

		for i, id := range proj.FollowedUsers {
			if id == userID {
				//check if we are removing the last item
				if i != len(proj.FollowedUsers)-1 { //if no, popout
					proj.FollowedUsers = append(proj.FollowedUsers[:i], proj.FollowedUsers[i+1:]...)
				} else { //if yes the popback
					proj.FollowedUsers = proj.FollowedUsers[:i]
				}
				break
			}
		}

		err = postProj.UpdateProject(eclient, projectID, "FollowedUsers", proj.FollowedUsers)
		if err != nil {
			return err
		}
	} else { //if the user is not followed, follow them
		//update user with new project id
		err = postUser.UpdateUser(eclient, userID, "FollowingProject", append(usr.FollowingProject, projectID))
		if err != nil {
			return err
		}

		//update project with new follower id
		err = postProj.UpdateProject(eclient, projectID, "FollowedUsers", append(proj.FollowedUsers, userID))
		if err != nil {
			return err
		}

	}

	return nil
}
