package uses

import (
	post "github.com/sea350/ustart_go/backend/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChangeEducation ...
func ChangeEducation(eclient *elastic.Client, userID string, accType int, hs string, hsGrad string, uni string, uniGrad string, major []string, minor []string) error {

	err := post.UpdateUser(eclient, userID, "AccType", accType)
	if err != nil {
		return err
	}
	err = post.UpdateUser(eclient, userID, "HighSchool", hs)
	if err != nil {
		return err
	}
	err = post.UpdateUser(eclient, userID, "HSGradDate", hsGrad)
	if err != nil {
		return err
	}
	err = post.UpdateUser(eclient, userID, "UndergradSchool", uni)
	if err != nil {
		return err
	}
	err = post.UpdateUser(eclient, userID, "CollegeGradDate", hsGrad)
	if err != nil {
		return err
	}
	err = post.UpdateUser(eclient, userID, "Majors", major)
	if err != nil {
		return err
	}
	err = post.UpdateUser(eclient, userID, "Minors", minor)
	return err
}
