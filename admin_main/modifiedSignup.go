package main

import (
	"fmt"
	"log"

	elastic "github.com/olivere/elastic"
	getUser "github.com/sea350/ustart_go/get/user"
	getWarning "github.com/sea350/ustart_go/get/warning"
	"github.com/sea350/ustart_go/globals"
	postChat "github.com/sea350/ustart_go/post/chat"
	postFollow "github.com/sea350/ustart_go/post/follow"
	postNotif "github.com/sea350/ustart_go/post/notification"
	postUser "github.com/sea350/ustart_go/post/user"
	postWarning "github.com/sea350/ustart_go/post/warning"
	types "github.com/sea350/ustart_go/types"

	"errors"
	"time"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

//SignUpBasic ... A basic user signup process
//Requires all basic signup feilds (email, password ...)
//Returns an error if there was a problem with database submission
func s(eclient *elastic.Client, username string, email string, password []byte, fname string, lname string, school string, major []string, bday time.Time, currYear string, addressIP string) error { //, country string, state string, city string, zip string) error {

	newSignWarning, err := getWarning.SingupWarningByIP(eclient, addressIP)
	if err != nil {
		return err
	}
	if newSignWarning.SignLockoutUntil.After(time.Now()) {
		err := errors.New("IP Address in Lockout")
		return err
	}

	inUse, err := getUser.EmailInUse(eclient, email)
	if err != nil {
		return err
	}
	if inUse { //We start keeping track here of signup warnings
		newSignWarning.SignIPAddress = addressIP
		newSignWarning.SignNumberofAttempts = newSignWarning.SignNumberofAttempts + 1
		if newSignWarning.SignLastAttempt.IsZero() {
			newSignWarning.SignLastAttempt = time.Now()
		} else {
			if time.Since(newSignWarning.SignLastAttempt) >= (time.Hour * 168) {
				newSignWarning.SignNumberofAttempts = 0
				newSignWarning.SignLockoutCounter = 0
			}
			newSignWarning.SignLastAttempt = time.Now()
		}

		if newSignWarning.SignNumberofAttempts > 10 {
			newSignWarning.SignLockoutCounter = newSignWarning.SignLockoutCounter + 1
			newSignWarning.SignLockoutUntil = newSignWarning.SignLastAttempt.Add(time.Minute * 30 * time.Duration(lockoutOP2(newSignWarning.SignLockoutCounter)))
			newSignWarning.SignNumberofAttempts = 0
		}
		if newSignWarning.SignDiscovered == false {
			newSignWarning.SignDiscovered = true
		}
		postWarning.ReIndexSignupWarning(eclient, newSignWarning, addressIP)
		return errors.New("email is in use ")
	}

	validEmail := ValidEmail(email)
	if !validEmail {
		if newSignWarning.SignDiscovered == true {
			newSignWarning.SignIPAddress = addressIP
			newSignWarning.SignNumberofAttempts = newSignWarning.SignNumberofAttempts + 1
			if newSignWarning.SignLastAttempt.IsZero() {
				newSignWarning.SignLastAttempt = time.Now()
			} else {
				if time.Since(newSignWarning.SignLastAttempt) >= (time.Hour * 168) {
					newSignWarning.SignNumberofAttempts = 0
					newSignWarning.SignLockoutCounter = 0
				}
				newSignWarning.SignLastAttempt = time.Now()
			}

			if newSignWarning.SignNumberofAttempts > 10 {
				newSignWarning.SignLockoutCounter = newSignWarning.SignLockoutCounter + 1
				newSignWarning.SignLockoutUntil = newSignWarning.SignLastAttempt.Add(time.Minute * 30 * time.Duration(lockoutOP2(newSignWarning.SignLockoutCounter)))
				newSignWarning.SignNumberofAttempts = 0
			}
			postWarning.ReIndexSignupWarning(eclient, newSignWarning, addressIP)
		}
		return errors.New("invalid email")
	}

	inUse, err = getUser.UsernameInUse(eclient, username)
	if err != nil {
		return err
	}
	if inUse {
		return errors.New("username is in use")
	}

	newUsr := types.User{}
	newUsr.Avatar = "https://i.imgur.com/8BnkFLO.png"
	newUsr.Banner = "https://i.imgur.com/XTj1t1J.png"

	newUsr.FirstName = fname
	newUsr.LastName = lname
	newUsr.Email = email
	newUsr.Username = username
	//New user verification process
	newUsr.Verified = false
	// SendVerificationEmail(email)
	// token, err := GenerateRandomString(32)
	// if err != nil {
	// 	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 	log.Println(err)
	// }
	// newUsr.AuthenticationCode = token
	// subject := "Your verification link"
	// //link := globals.SiteURL + ":" + globals.Port + "/Activation/?email=" + email + "&verifCode=" + token
	// link := globals.SiteURL + "/Activation/?email=" + email + "&verifCode=" + token
	// r := NewRequest([]string{email}, subject)
	// r.Send(globals.HTMLPATH+"email_template.html", map[string]string{"username": username, "link": link,
	// 	"contentjuan":   "We received a request to activate your Ustart Account. We would love to assist you!",
	// 	"contentdos":    "Simply click the button below to verify your account",
	// 	"contenttres":   "VERIFY ACCOUNT",
	// 	"contentquatro": "a new account"})

	newUsr.Password = password
	newUsr.University = school
	newUsr.Majors = major
	newUsr.Dob = bday
	newLoc := types.LocStruct{}
	// newLoc.Country = country
	// newLoc.State = state
	// newLoc.City = city
	// newLoc.Zip = zip
	newUsr.Location = newLoc
	newUsr.Visible = true
	newUsr.Status = true

	badgeIDs, badgeTags, err := BadgeSetup(eclient, email)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	newUsr.Tags = badgeTags
	newUsr.BadgeIDs = badgeIDs
	newUsr.AccCreation = time.Now()
	if currYear == "Freshman" {
		newUsr.Class = 0
	} else if currYear == "Sophomore" {
		newUsr.Class = 1
	} else if currYear == "Junior" {
		newUsr.Class = 2
	} else if currYear == "Senior" {
		newUsr.Class = 3
	} else if currYear == "Graduate" {
		newUsr.Class = 4
	} else if currYear == "Alumni" {
		newUsr.Class = 5
	} else if currYear == "Faculty" {
		newUsr.Class = 6
	} else if currYear == "Other" {
		newUsr.Class = 7
	} else {
		newUsr.Class = -1
	}

	id, retErr := postUser.ReIndexUser(eclient, "6_7VeGkBN3Vvtvdi-6q6", newUsr)
	if retErr != nil {
		return retErr
	}

	errFollow := postFollow.IndexFollow(eclient, id)
	if errFollow != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(errFollow)
	}
	newProxy := types.ProxyMessages{DocID: id, Class: 1}
	proxyID, err := postChat.IndexProxyMsg(eclient, newProxy)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	newProxyNotif := types.ProxyNotifications{DocID: id}
	newProxyNotif.Settings.Default()
	_, err = postNotif.IndexProxyNotification(eclient, newProxyNotif)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	err = postUser.UpdateUser(eclient, id, "ProxyMessages", proxyID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	go SendVerificationEmail(eclient, email)

	return err
}

func main() {
	pwd := []byte("Ilikedogs1")
	err := s(eclient, "testuser", "minhazur.bhuiyan@nyu.edu", pwd, "Test", "User", "NYU", []string{"Quality Assurance"}, time.Now(), "Other", " ")
	if err != nil {
		fmt.Println(err)
	}
}
