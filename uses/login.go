package uses

import (
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	types "github.com/sea350/ustart_go/types"
	"golang.org/x/crypto/bcrypt"
	elastic "gopkg.in/olivere/elastic.v5"
)

//Login ... Checks username/password combo
//Requires username and password
//Returns whether or not username and password match, a type SessionUser, and an error
func Login(eclient *elastic.Client, userEmail string, password []byte) (bool, types.SessionUser, error) {

	var loginSucessful = false
	var userSession types.SessionUser

	inUse, err := get.EmailInUse(eclient, userEmail)
	if err != nil {
		return loginSucessful, userSession, err
	}
	if !inUse {
		err := errors.New("Invalid Email")
		return loginSucessful, userSession, err
	}

	usr, err := get.UserByEmail(eclient, userEmail)
	if err != nil {
		return loginSucessful, userSession, err
	}

	passErr := bcrypt.CompareHashAndPassword(usr.Password, password)

	if passErr != nil {
		return false, userSession, passErr
	}

	uID, err := get.UserIDByEmail(eclient, userEmail)
	if err != nil {
		return loginSucessful, userSession, err
	}

	loginSucessful = true
	userSession.FirstName = usr.FirstName
	userSession.LastName = usr.LastName
	userSession.Email = userEmail
	userSession.DocID = uID
	userSession.Username = usr.Username
	userSession.Avatar = usr.Avatar

	return loginSucessful, userSession, err

}
