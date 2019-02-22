package uses

import (
	"errors"
	"math"
	"time"

	get "github.com/sea350/ustart_go/get/user"
	post "github.com/sea350/ustart_go/post/user"
	types "github.com/sea350/ustart_go/types"
	"golang.org/x/crypto/bcrypt"
	elastic "github.com/olivere/elastic"
)

//Login ... Checks username/password combo
//Requires username and password
//Returns whether or not username and password match, a type SessionUser, and an error
func Login(eclient *elastic.Client, userEmail string, password []byte, addressIP string) (bool, types.SessionUser, error) {

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

	uID, err := get.UserIDByEmail(eclient, userEmail)
	if err != nil {
		return loginSucessful, userSession, err
	}

	//Email is valid, we need to check if we are in a login lockout or not based on IP address

	if len(usr.LoginWarnings) == 0 {
		usr.LoginWarnings = make(map[string]types.LoginWarning)
	}
	loginData, ipExists := usr.LoginWarnings[addressIP]
	if ipExists {
		if loginData.LockoutUntil.After(time.Now()) {
			err := errors.New("Account in Lockout until " + loginData.LockoutUntil.String())
			return loginSucessful, userSession, err
		}
	}

	passErr := bcrypt.CompareHashAndPassword(usr.Password, password)

	if passErr != nil {
		return loginSucessful, userSession, errors.New("Invalid Password")
	}

	// if passErr != nil {
	// 	//If password incorrect, the following evaluation on login lockout procedure is done
	// 	//If IP doesn't exist, this a New Warning for a New IP Address. Else we are simply updating an existing one our usr

	// 	if !ipExists {
	// 		var newWarning types.LoginWarning
	// 		newWarning.IPAddress = addressIP
	// 		newWarning.NumberAttempts = newWarning.NumberAttempts + 1
	// 		newWarning.LastAttempt = time.Now()
	// 		loginData = newWarning
	// 		usr.LoginWarnings[addressIP] = loginData
	// 	}

	// 	if ipExists {
	// 		if time.Since(loginData.LastAttempt) >= (time.Hour * 168) {
	// 			loginData.NumberAttempts = 0
	// 			loginData.LockoutCounter = 0
	// 		}
	// 		loginData.NumberAttempts = loginData.NumberAttempts + 1
	// 		loginData.LastAttempt = time.Now()
	// 		if loginData.NumberAttempts > 5 {
	// 			loginData.LockoutCounter = loginData.LockoutCounter + 1
	// 			loginData.LockoutUntil = loginData.LastAttempt.Add(time.Minute * 5)
	// 			loginData.NumberAttempts = 0
	// 		}
	// 		usr.LoginWarnings[addressIP] = loginData
	// 	}

	// 	//Update in Elastic Search Client all of our Login Warning information

	// 	err2 := post.UpdateUser(eclient, uID, "LoginWarnings", usr.LoginWarnings)
	// 	if err != nil {
	// 		return false, userSession, err2
	// 	}
	// 	return false, userSession, passErr
	// }

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
		err2 := post.UpdateUser(eclient, usrID, "LoginWarnings", map[string]types.LoginWarning{})
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
