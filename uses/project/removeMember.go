package uses

import elastic "gopkg.in/olivere/elastic.v5"

//NEEDS TO BE REPAIRED

//RemoveMember ... CHANGES NECESSARY DATA FROM USER AND PROJECT FOR REMOVING A MEMBER
//Requires
//Returns an error
func RemoveMember(eclient *elastic.Client, projectID string, userID string) error {

	memberModLock.Lock()
	defer memberModLock.Unlock()

	usr, err := get.GetUserByID(eclient, userID)
	if err != nil {
		return err
	}
	proj, projErr := get.GetProjectByID(eclient, projectID)
	if err != nil {
		return projErr
	}

	var projIdx int
	var usrIdx int
	for idx := range usr.Projects {
		if usr.Projects[idx].ProjectID == projectID {
			usrIdx = idx
			break
		}
	}

	err = post.UpdateUser(eclient, userID, "Projects", append(usr.Projects[:usrIdx], usr.Projects[usrIdx+1:]...))
	if err != nil {
		return err
	}

	for index := range proj.Members {
		if proj.Members[index].MemberID == projectID {
			projIdx = index
			break
		}
	}

	projErr = post.DeleteMember(eclient, projectID, proj.Members[projIdx])
	if projErr != nil {
		return projErr
	}

	return nil

}
