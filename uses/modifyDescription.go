package uses

import (
	get "github.com/sea350/ustart_go/get/user"
	post "github.com/sea350/ustart_go/post/user"
	elastic "github.com/olivere/elastic"
)

//ModifyDescription ... CHANGES A SPECIFIC USER'S DESCRIPTION
//Requires the target user's docID and the new description
//Returns an error
func ModifyDescription(eclient *elastic.Client, userID string, newDescription string) error {

	usr, err := get.UserByID(eclient, userID)

	if err != nil {
		return err
	}

	usr.Description = []rune(newDescription)

	retErr := post.UpdateUser(eclient, userID, "Description", usr)
	return retErr

}
