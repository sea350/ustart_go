package uses

import (
	"errors"
	"math"
	"time"

	get "github.com/sea350/ustart_go/get/user"
	post "github.com/sea350/ustart_go/post/user"
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
	//We make this condition integer first for later
	var ipExists bool = false
	var recordWarning int

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

	//Email is valid, we need to check if we are in a login lockout or not based on IP address
	for i := 0; i < len(usr.LoginWarnings); i++ {
		if usr.LoginWarnings[i].IPAddress == addressIP {
			if usr.LoginWarnings[i].LockoutUntil.After(time.Now()) {
				err := errors.New("Account in Lockout")
				return loginSucessful, userSession, err
			}
			ipExists = true
			recordWarning = i
			break
		}
	}

	passErr := bcrypt.CompareHashAndPassword(usr.Password, password)

	if passErr != nil {
		//If password incorrect, the following evaluation on login lockout procedure is done
		//If IP doesn't exist, this a New Warning for a New IP Address. Else we are simply updating an existing one our usr
		if !(ipExists) {
			var newWarning types.LoginWarning
			newWarning.IPAddress = addressIP
			newWarning.NumberAttempts = newWarning.NumberAttempts + 1
			newWarning.LastAttempt = time.Now()
			usr.LoginWarnings = append(usr.LoginWarnings, newWarning)
		}
		if ipExists {
			if time.Since(usr.LoginWarnings[recordWarning].LastAttempt) >= (time.Minute * 10080) {
				usr.LoginWarnings[recordWarning].NumberAttempts = 0
				usr.LoginWarnings[recordWarning].LockoutCounter = 0
			}
			usr.LoginWarnings[recordWarning].NumberAttempts = usr.LoginWarnings[recordWarning].NumberAttempts + 1
			usr.LoginWarnings[recordWarning].LastAttempt = time.Now()
			if usr.LoginWarnings[recordWarning].NumberAttempts > 5 {
				usr.LoginWarnings[recordWarning].LockoutCounter = usr.LoginWarnings[recordWarning].LockoutCounter + 1
				usr.LoginWarnings[recordWarning].LockoutUntil = usr.LoginWarnings[recordWarning].LastAttempt.Add(time.Minute * 5 * time.Duration(lockoutOP(usr.LoginWarnings[recordWarning].LockoutCounter)))
				usr.LoginWarnings[recordWarning].NumberAttempts = 0
			}

		}
		//Update in Elastic Search Client all of our Login Warning information
		usrID, err1 := get.UserIDByEmail(eclient, userEmail)
		if err1 != nil {
			return false, userSession, err1
		}
		err2 := post.UpdateUser(eclient, usrID, "LoginWarnings", usr.LoginWarnings)
		if err != nil {
			return false, userSession, err2
		}
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

	if ipExists {
		usrID, err1 := get.UserIDByEmail(eclient, userEmail)
		if err1 != nil {
			return false, userSession, err1
		}
		err2 := post.UpdateUser(eclient, usrID, "LoginWarnings", []types.LoginWarning{})
		if err2 != nil {
			return false, userSession, err2
		}
	}

	return loginSucessful, userSession, err

}

func lockoutOP(LockoutCounter int) int {
	timeOP := int(math.Exp2(float64(LockoutCounter) - 1))
	return timeOP
}
