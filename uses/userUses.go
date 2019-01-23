package uses

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"math"

	getEntry "github.com/sea350/ustart_go/get/entry"
	getGuestCode "github.com/sea350/ustart_go/get/guestCode"
	getUser "github.com/sea350/ustart_go/get/user"
	getWarning "github.com/sea350/ustart_go/get/warning"
	postChat "github.com/sea350/ustart_go/post/chat"
	postEntry "github.com/sea350/ustart_go/post/entry"
	postFollow "github.com/sea350/ustart_go/post/follow"
	updateCode "github.com/sea350/ustart_go/post/guestCode"
	postNotif "github.com/sea350/ustart_go/post/notification"
	postUser "github.com/sea350/ustart_go/post/user"
	postWarning "github.com/sea350/ustart_go/post/warning"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"

	"errors"
	"time"
)

//SignUpBasic ... A basic user signup process
//Requires all basic signup feilds (email, password ...)
//Returns an error if there was a problem with database submission
func SignUpBasic(eclient *elastic.Client, username string, email string, password []byte, fname string, lname string, country string, state string, city string, zip string, school string, major []string, bday time.Time, currYear string, addressIP string) error {

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
	newUsr.Avatar = "https://i.imgur.com/TYFKsdi.png"
	newUsr.Banner = "https://scontent-lga3-1.xx.fbcdn.net/v/t1.0-9/20882995_716470841896823_2067565050010385700_n.jpg?_nc_cat=104&_nc_ht=scontent-lga3-1.xx&oh=358772e4047c100765f12c22b72efd68&oe=5C83172C"

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
	newLoc.Country = country
	newLoc.State = state
	newLoc.City = city
	newLoc.Zip = zip
	newUsr.Location = newLoc
	newUsr.Visible = true
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
	} else {
		newUsr.Class = -1
	}

	id, retErr := postUser.IndexUser(eclient, newUsr)
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

	time.Sleep(30 * time.Second)
	SendVerificationEmail(eclient, email)

	return err
}

//GuestSignUpBasic ... Signup function for guests (non-NYU users)
func GuestSignUpBasic(eclient *elastic.Client, username string, email string, password []byte, fname string, lname string, country string, state string, city string, zip string, school string, major []string, bday time.Time, currYear string, guestCode string) error {
	inUse, err := getUser.EmailInUse(eclient, email)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return errors.New("Error checking if email is in Use")
	}
	if inUse {
		return errors.New("Error: Email is in use ")
	}

	//Check for valid guest email address
	if !ValidGuestEmail(email) {
		// return errors.New("Error: Invalid email")
		return err
	}

	inUse, err = getUser.UsernameInUse(eclient, username)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return err
	}
	if inUse {
		// return errors.New("Error: Username is in use")
		return err
	}

	//Gets GuestCode object and also check if guest code is valid
	guestObj, err := getGuestCode.GuestCodeByID(eclient, guestCode)
	if err != nil {
		return errors.New("There was a problem finding your code")
	}

	validGuestcode, err := ValidGuestCode(eclient, guestCode)
	if !validGuestcode {
		//return errors.New("Error: Invalid code")
		return err
	}

	newUsr := types.User{}
	newUsr.Avatar = "https://i.imgur.com/TYFKsdi.png"
	newUsr.Banner = "https://scontent-lga3-1.xx.fbcdn.net/v/t1.0-9/20882995_716470841896823_2067565050010385700_n.jpg?_nc_cat=104&_nc_ht=scontent-lga3-1.xx&oh=358772e4047c100765f12c22b72efd68&oe=5C83172C"
	newUsr.FirstName = fname
	newUsr.LastName = lname
	newUsr.Email = email
	newUsr.Username = username
	newUsr.Verified = false
	// token, err := GenerateRandomString(32)
	// if err != nil {
	// 	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 	log.Println(err)
	// 	return err
	// }
	// newUsr.AuthenticationCode = token

	// subject := "Your verification link"
	// link := globals.SiteURL + ":" + globals.Port + "/Activation/?email=" + email + "&verifCode=" + token
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
	newLoc.Country = country
	newLoc.State = state
	newLoc.City = city
	newLoc.Zip = zip
	newUsr.Location = newLoc
	newUsr.Visible = true
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
	} else {
		newUsr.Class = -1
	}

	id, retErr := postUser.IndexUser(eclient, newUsr)
	if retErr != nil {
		return retErr
	}

	errFollow := postFollow.IndexFollow(eclient, id)
	if errFollow != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(errFollow)
		return errFollow
	}
	newProxy := types.ProxyMessages{DocID: id, Class: 1}
	proxyID, err := postChat.IndexProxyMsg(eclient, newProxy)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return err
	}

	newProxyNotif := types.ProxyNotifications{DocID: id}
	newProxyNotif.Settings.Default()
	_, err = postNotif.IndexProxyNotification(eclient, newProxyNotif)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return err
	}

	//Add user to GuestCode array
	tempUserArray := guestObj.Users
	tempUserArray = append(tempUserArray, id)
	err = updateCode.UpdateGuestCode(eclient, guestCode, "Users", tempUserArray)
	if err != nil {
		return err
	}

	err = postUser.UpdateUser(eclient, id, "ProxyMessages", proxyID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	time.Sleep(30 * time.Second)
	SendVerificationEmail(eclient, email)

	return err
}

//GenerateRandomBytes ... GENERATES RANDOM BYTES FOR ACCOUNT VERIFICATION
//Requires n, the desired length of bytes
//Returns the randomly generated bytes
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

//GenerateRandomString ... GENERATES RANDOM STRING FROM BYTES
//Requires s, the desired length of the string
//Returns the randomly generated string in base64
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

//UserLikeEntry ... ALLOWS A USER TO LIKE AN ENTRY
//Requires the entry's docID, and docID of the person who is liking the entry
//Returns an error
func UserLikeEntry(eclient *elastic.Client, entryID string, likerID string) error {

	err := postEntry.AppendLike(eclient, entryID, likerID)
	if err != nil {
		return err
	}

	err = postUser.AppendLikedEntryID(eclient, likerID, entryID)
	if err != nil {
		return err
	}

	return nil
}

//UserUnlikeEntry ... ALLOWS A USER TO UNLIKE AN ENTRY
//Requires the entry's docID, and docID of the person who is unliking the entry
//Returns an error
func UserUnlikeEntry(eclient *elastic.Client, entryID string, likerID string) error {

	//DeleteLike deletes from post
	err := postEntry.DeleteLike(eclient, entryID, likerID)
	if err != nil {
		return err
	}

	//DeleteLikedEntryID deletes from usr
	err = postUser.DeleteLikedEntryID(eclient, likerID, entryID)
	if err != nil {
		return err
	}

	return nil
}

//IsLiked ... CHECKS IF AN ENTRY IS ALREADY LIKED BY A USER
//Requires the entry's docID, the user's docID
//Returns true if the entry is liked and false if not, and an error
func IsLiked(eclient *elastic.Client, entryID string, viewerID string) (bool, error) {
	isLiked := false
	entry, err := getEntry.EntryByID(eclient, entryID)
	if err != nil {
		return isLiked, err
	}
	for _, element := range entry.Likes {
		if element.UserID == viewerID {
			isLiked = true
			return isLiked, err
		}
	}
	return isLiked, err
}

//RequestColleague ... SENDS A COLLEGUE REQUEST FROM ONE USER TO ANOTHER
//NOTE: This function checks if a request has already been sent and if the users are allready colleagues
//WARNING: needs to be revised
//Requires the sender's docID and the request receiver's docID
//Returns an error
func RequestColleague(eclient *elastic.Client, usrID string, requestedUserID string) error {
	usr, err := getUser.UserByID(eclient, usrID)
	if err != nil {
		return err
	}

	for _, element := range usr.SentCollReq {
		if element == usrID {
			return errors.New("You have already requested this user")
		}

	}

	for _, element := range usr.Colleagues {
		if element == requestedUserID {
			return errors.New("You have already requested this user")
		}
	}

	//CONFUSING, REVISE!!!!!!!!!!!!!!!!!!!!1111
	err = postUser.AppendCollReq(eclient, usrID, requestedUserID, true)
	if err != nil {
		return err
	}

	err = postUser.AppendCollReq(eclient, requestedUserID, requestedUserID, false)
	return err
}

//ReplyToColleagueRequest ... WARNING NEEDS TO BE FIXED
//ALLOWS A USER TO REPLY TO A COLLEAGUE REQUEST
//Requires user's docID, the docID of the user who sent the request, and true if they acept the request/ false if declined
//Returns an error
func ReplyToColleagueRequest(eclient *elastic.Client, usrID string, requestedUserID string, reply bool) error {
	if reply == true {

	}

	return errors.New("Could not reply to colleague request")

}

//UpdateUserLinks ... REPLACES THE ENTIRETY OF A USER'S LINKS WITH AN UPDATED LIST
//Requires the target user's docID and an updated array of type Link
//Returns an error
func UpdateUserLinks(eclient *elastic.Client, userID string, lynx []types.Link) error {
	err := postUser.UpdateUser(eclient, userID, "QuickLinks", lynx)
	return err
}

//UpdateUserTags ... REPLACES THE ENTIRETY OF A USER'S TAGS WITH AN UPDATED LIST
//Requires the target user's docID and an updated array of strings
//Returns an error
func UpdateUserTags(eclient *elastic.Client, userID string, tags []string) error {
	err := postUser.UpdateUser(eclient, userID, "Tags", tags)
	return err
}

/*
//UserFollow ... ALLOWS A USER TO FOLLOW SOMEONE ELSE
//Requires the follower's docID and the followed docID
//Returns an error
func UserFollow(eclient *elastic.Client, hostID string, viewerID string) error {
	//true = append to following
	followingErr := postUser.AppendFollowing(eclient, viewerID, hostID)
	if followingErr != nil {
		return followingErr
	}
	//false = append to followers
	followErr := postUser.AppendFollower(eclient, hostID, viewerID)
	if followErr != nil {
		return followErr
	}

	return nil
}
*/

//UserUnfollow ... ALLOWS A USER TO UNFOLLOW SOMEONE ELSE
//Returns an error
func UserUnfollow(eclient *elastic.Client, hostID string, viewerID string) error {
	err := postUser.DeleteFollow(eclient, hostID, viewerID, false)
	if err != nil {
		fmt.Println("userUses line 252")
		return err
	}
	err = postUser.DeleteFollow(eclient, viewerID, hostID, true)
	if err != nil {
		fmt.Println("userUses line 257")
	}
	return err
}

//IsFollowed ... CHECKS IF A USER IS FOLLWING SOMEONE ELSE
//Returns an error
func IsFollowed(eclient *elastic.Client, usrID string, viewerID string) (bool, error) {
	isFollowed := false
	user, err := getUser.UserByID(eclient, usrID)
	if err != nil {
		return isFollowed, err
	}
	for _, element := range user.Followers {
		if element == viewerID {
			isFollowed = true
			return isFollowed, err
		}
	}
	return isFollowed, err
}

//NumFollow ... FINDS THE NUMBER OF PEOPLE FOLLOWED BY/FOLLOWING SOMEONE
//Requires the user's docID, and true if you want num people person is following/ false if you want number of followers
//Returns an error
func NumFollow(eclient *elastic.Client, usrID string, whichOne bool) (int, error) {

	usr, err := getUser.UserByID(eclient, usrID)
	if err != nil {
		return -1, err
	}
	if whichOne {
		return len(usr.Following), nil
	}

	return len(usr.Followers), nil

}

func lockoutOP2(LockoutCounter int) int {
	timeOP := int(math.Exp2(float64(LockoutCounter) - 1))
	return timeOP
}
