package uses

import (
	bcrypt "golang.org/x/crypto/bcrypt"

	get "github.com/sea350/ustart_go/get/user"
	post "github.com/sea350/ustart_go/post/user"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChangePassword ...
func ChangePassword(eclient *elastic.Client, userID string, oldPass []byte, newPass []byte) error {

	usr, err := get.UserByID(eclient, userID)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(usr.Password, oldPass)
	if err != nil {
		return err
	}
	/*if !bytes.Equal(usr.Password, oldPass) {
		return errors.New("Invalid old password")
	}*/
	newHashedPass, err := bcrypt.GenerateFromPassword(newPass, bcrypt.DefaultCost)
	// err = post.UpdateUser(eclient, userID, "Password", newHashedPass)
	usr.Password = newHashedPass

	clearWarnings := make(map[string]types.LoginWarning)
	usr.LoginWarnings = clearWarnings
	//Clear login lockout counter
	if err == nil {
		err = post.ReindexUser(eclient, userID, usr)
	}

	return err

}
