package uses

import (
	"errors"
	"time"

	get "github.com/sea350/ustart_go/get/user"
	types "github.com/sea350/ustart_go/types"
	"golang.org/x/crypto/bcrypt"
	elastic "gopkg.in/olivere/elastic.v5"
)

//Login ... Checks username/password combo
//Requires username and password
//Returns whether or not username and password match, a type SessionUser, and an error
func Login(eclient *elastic.Client, userEmail string, password []byte, addressIP string) (bool, types.SessionUser, error) {
	//addressIP = the user IP, Steven knows how to do this

	var loginSucessful = false
	var userSession types.SessionUser

	//We want to keep track of login attempts and lock out users who have too many failed attempts for a period of time
	var loginWarnings types.LoginWarning

	inUse, err := get.EmailInUse(eclient, userEmail)
	if err != nil {
		return loginSucessful, userSession, err
	}
	if !inUse {
		err := errors.New("Invalid Email")
		return loginSucessful, userSession, err
	}

	//Email is valid, we need to check if we are in a login lockout or not from multiple attempts
	if loginWarnings.LockoutUntil.After(time.Now()) {
		err := errors.New("Account in Lockout")
		return loginSucessful, userSession, err
	}

	usr, err := get.UserByEmail(eclient, userEmail)
	if err != nil {
		return loginSucessful, userSession, err
	}

	passErr := bcrypt.CompareHashAndPassword(usr.Password, password)

	if passErr != nil {
		loginWarnings.NumberAttempts++
		return false, userSession, passErr
	}

	uID, err := get.UserIDByEmail(eclient, userEmail)
	if err != nil {
		//If password incorrect, the following evaluation on login lockout procedure is followed
		loginWarnings.NumberAttempts = loginWarnings.NumberAttempts + 1
		loginWarnings.LastAttempt = time.Now()
		if loginWarnings.NumberAttempts > 5 {
			loginWarnings.LockoutCounter = loginWarnings.LockoutCounter + 1
			loginWarnings.LockoutUntil = loginWarnings.LastAttempt.Add(time.Minute * (5 + 5*time.Duration(loginWarnings.LockoutCounter-1)))
			loginWarnings.NumberAttempts = 0
		}
		usr.LoginWarningsIP = append(usr.LoginWarningsIP, addressIP)
		return loginSucessful, userSession, err
	}

	loginSucessful = true
	userSession.FirstName = usr.FirstName
	userSession.LastName = usr.LastName
	userSession.Email = userEmail
	userSession.DocID = uID
	userSession.Username = usr.Username
	userSession.Avatar = usr.Avatar

	//Clear Login Warning struct and the IP array, which I still don't know where to put
	loginWarnings = types.LoginWarning{}

	return loginSucessful, userSession, err

}
