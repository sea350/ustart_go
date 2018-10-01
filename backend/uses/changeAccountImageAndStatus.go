package uses

import (
	post "github.com/sea350/ustart_go/backend/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
	//"golang.org/x/crypto/bcrypt"
	//"time"
)

//ChangeAccountImagesAndStatus ...
func ChangeAccountImagesAndStatus(eclient *elastic.Client, userID string, image string, status bool, banner string, action string) error {

	if action == "Avatar" {
		err := post.UpdateUser(eclient, userID, "Avatar", image)
		return err
	} else if action == "Status" {
		err := post.UpdateUser(eclient, userID, "Status", status)
		return err
	} else {
		err := post.UpdateUser(eclient, userID, "Banner", banner)
		return err
	}

}
