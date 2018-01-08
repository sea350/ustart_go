package uses

import (
	"bytes"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	post "github.com/sea350/ustart_go/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChangePassword ...
func ChangePassword(eclient *elastic.Client, userID string, oldPass []byte, newPass []byte) error {

	usr, err := get.UserByID(eclient, userID)
	if err != nil {
		return err
	}
	if !bytes.Equal(usr.Password, oldPass) {
		return errors.New("Invalid old password")
	}
	err = post.UpdateUser(eclient, userID, "Password", newPass)
	return err

}
