package uses

import (
	post "github.com/sea350/ustart_go/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChangeEducation ...
func ChangeEducation(eclient *elastic.Client, userID string, uni string, uniGrad string, major []string, minor []string, class string) error {

	// err = post.UpdateUser(eclient, userID, "HighSchool", hs)
	// if err != nil {
	// 	return err
	// }
	// err = post.UpdateUser(eclient, userID, "HSGradDate", hsGrad)
	// if err != nil {
	// 	return err
	// }
	err := post.UpdateUser(eclient, userID, "UndergradSchool", uni)
	if err != nil {
		return err
	}
	// err = post.UpdateUser(eclient, userID, "CollegeGradDate", hsGrad)
	// if err != nil {
	// 	return err
	// }
	err = post.UpdateUser(eclient, userID, "Majors", major)
	if err != nil {
		return err
	}
	err = post.UpdateUser(eclient, userID, "Minors", minor)
	return err

}
