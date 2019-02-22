package uses

import (
	"strings"

	post "github.com/sea350/ustart_go/post/user"
	elastic "github.com/olivere/elastic"
)

//ChangeEducation ...
func ChangeEducation(eclient *elastic.Client, userID string, uni string, major []string, minor []string, class string) error {

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
	if err != nil {
		return err
	}

	classInt := -1
	switch strings.ToLower(class) {
	case "freshman":
		classInt = 0
	case "sophomore":
		classInt = 1
	case "junior":
		classInt = 2
	case "senior":
		classInt = 3
	case "graduate":
		classInt = 4
	case "alumni":
		classInt = 5
	}

	err = post.UpdateUser(eclient, userID, "Class", classInt)
	return err

}
