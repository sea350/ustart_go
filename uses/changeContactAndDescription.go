package uses

import (
	post "github.com/sea350/ustart_go/post/user"
	elastic "github.com/olivere/elastic"
)

//ChangeContactAndDescription ...
func ChangeContactAndDescription(eclient *elastic.Client, userID string, phone string, phoneVis bool, gender string, genderVis bool, emailVis bool, description []rune) error {

	err := post.UpdateUser(eclient, userID, "Phone", phone)
	if err != nil {
		return err
	}
	err = post.UpdateUser(eclient, userID, "PhoneVis", phoneVis)
	if err != nil {
		return err
	}
	err = post.UpdateUser(eclient, userID, "Gender", gender)
	if err != nil {
		return err
	}
	err = post.UpdateUser(eclient, userID, "GenderVis", genderVis)
	if err != nil {
		return err
	}
	err = post.UpdateUser(eclient, userID, "EmailVis", emailVis)
	if err != nil {
		return err
	}
	err = post.UpdateUser(eclient, userID, "Description", description)
	return err

}
