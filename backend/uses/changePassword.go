package uses

import (
	"errors"

	bcrypt "golang.org/x/crypto/bcrypt"

	get "github.com/sea350/ustart_go/backend/get/user"
	post "github.com/sea350/ustart_go/backend/post/user"
	types "github.com/sea350/ustart_go/backend/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChangePassword ...
func ChangePassword(eclient *elastic.Client, userID string, oldPass []byte, newPass []byte) error {

	usr, err := get.UserByID(eclient, userID)
	if err != nil {
		return err
	}

	if bcrypt.CompareHashAndPassword(usr.Password, oldPass) != nil {
		return errors.New("Invalid old password")
	}
	/*if !bytes.Equal(usr.Password, oldPass) {
		return errors.New("Invalid old password")
	}*/
	newHashedPass, err := bcrypt.GenerateFromPassword(newPass, 10)
	err = post.UpdateUser(eclient, userID, "Password", newHashedPass)

	//Clear login lockout counter
	if err == nil {
		err = post.UpdateUser(eclient, userID, "LoginWarnings", []types.LoginWarning{})
	}

	return err

}
