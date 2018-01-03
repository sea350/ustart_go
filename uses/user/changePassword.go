package uses

import (
	"bytes"
	"errors"

	post "github.com/sea350/ustart_go/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChangePassword ...
func ChangePassword(eclient *elastic.Client, userID string, oldPass []byte, newPass []byte) error {

	usr, err := get.GetUserByID(eclient, userID)
	if err != nil {
		return err
	}
	if !bytes.Equal(usr.Password, oldPass) {
		return errors.New("Invalid old password")
	}
	err = post.UpdateUser(eclient, userID, "Password", newPass)
	return err

}
