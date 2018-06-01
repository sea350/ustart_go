package uses

import (
	"errors"
	"math"
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
	//We make this condition integer first for later
	var condition int
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
			condition = 1
			recordWarning = i
		}
	}

	passErr := bcrypt.CompareHashAndPassword(usr.Password, password)

	if passErr != nil {
		//If password incorrect, the following evaluation on login lockout procedure is done
		//If condition is 0, this a New Warning for a New IP Address. Else we are simply updating an existing one in usr
		if condition == 0 {
			var newWarning types.LoginWarning
			newWarning.IPAddress = addressIP
			newWarning.NumberAttempts = newWarning.NumberAttempts + 1
			newWarning.LastAttempt = time.Now()
			usr.LoginWarnings = append(usr.LoginWarnings, newWarning)
		}
		if condition == 1 {
			usr.LoginWarnings[recordWarning].NumberAttempts = usr.LoginWarnings[recordWarning].NumberAttempts + 1
			usr.LoginWarnings[recordWarning].LastAttempt = time.Now()
			if usr.LoginWarnings[recordWarning].NumberAttempts > 5 {
				usr.LoginWarnings[recordWarning].LockoutCounter = usr.LoginWarnings[recordWarning].LockoutCounter + 1
				usr.LoginWarnings[recordWarning].LockoutUntil = usr.LoginWarnings[recordWarning].LastAttempt.Add(time.Minute * (5 + time.Duration(lockoutOP(usr.LoginWarnings[recordWarning].LockoutCounter))))
				usr.LoginWarnings[recordWarning].NumberAttempts = 0
			}

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

	return loginSucessful, userSession, err

}

func lockoutOP(LockoutCounter int) int {
	timeOP := int(math.Exp2(float64(LockoutCounter) - 1))
	return timeOP
}
