package uses

import (
	get "github.com/sea350/ustart_go/get"
	post "github.com/sea350/ustart_go/post"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"

	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//SignUpBasic ... A basic user signup process
//Requires all basic signup feilds (email, password ...)
//Returns an error if there was a problem with database submission
func SignUpBasic(eclient *elastic.Client, email string, password []byte, fname string, lname string, country string, state string, city string, zip string, school string, major []string, bday time.Time, currYear string) error {

	inUse, err := get.EmailInUse(eclient, email)
	if err != nil {
		return err
	}
	if inUse {
		return errors.New("email is in use")
	}

	newUsr := types.User{}
	newUsr.FirstName = fname
	newUsr.LastName = lname
	newUsr.Email = email
	newUsr.Username = get.EmailToUsername(email)
	fmt.Println(newUsr.Username)
	fmt.Println("HELLO BUDDY psych")
	//hashPass := bcrypt.GenerateFromPassword(password,10)
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
	} else {
		newUsr.Class = 5
	}

	retErr := post.IndexUser(eclient, newUsr)

	return retErr
}

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

	usr, err := get.GetUserByEmail(eclient, userEmail)
	if err != nil {
		return loginSucessful, userSession, err
	}

	passErr := bcrypt.CompareHashAndPassword(usr.Password, password)

	if passErr != nil {
		return false, userSession, passErr
	}

	uID, err := get.GetUserIDByEmail(eclient, userEmail)
	if err != nil {
		return loginSucessful, userSession, err
	}

	loginSucessful = true
	userSession.FirstName = usr.FirstName
	userSession.LastName = usr.LastName
	userSession.Email = userEmail
	userSession.DocID = uID
	userSession.Username = usr.Username
	fmt.Println(userSession.Username)

	fmt.Println("_____________________________________-")
	fmt.Println(usr.Username)
	return loginSucessful, userSession, err

}

//UserPage ... Loads relevant information for the User page
//Requires username and the docID of the person viewing the page
//Returns a typer User, the user's docID, whether or not the viewer is following the person and an error
func UserPage(eclient *elastic.Client, username string, viewerID string) (types.User, string, bool, error) {

	var usr types.User

	var isFollowed bool

	userID, err := get.GetIDByUsername(eclient, username)
	if err != nil {
		return usr, userID, isFollowed, err
	}

	usr, err = get.GetUserByID(eclient, userID)
	if err != nil {
		return usr, userID, isFollowed, err
	}

	viewer, err := get.GetUserByID(eclient, viewerID)
	if err != nil {
		return usr, userID, isFollowed, err
	}

	for _, element := range viewer.Following {
		if element == userID {
			isFollowed = true
			break
		}
	}

	return usr, userID, isFollowed, err
}

//LoadEntries ... Loads a list of entries as journal entries, if an entry is invisible it is skipped
//Requires an array of entry ids
//Returns an of the data for those ids as journal entries, and an error
func LoadEntries(eclient *elastic.Client, loadList []string) ([]types.JournalEntry, error) {

	var entries []types.JournalEntry

	for _, entryID := range loadList {
		jEntry, err := ConvertEntryToJournalEntry(eclient, entryID)
		if err != nil {
			return entries, err
		}

		if !jEntry.Element.Visible {
			continue
		}

		entries = append(entries, jEntry)
	}

	return entries, nil
}

//LoadComments ... Loads the replies to a specific entry limited by limits
//Requires the parent entry id, the position of the first comment desired to load and the last comment wanted to load
//NOTE set uppper bound to -1 to pull to the end of the array
//Returns the parent entry as a JournalEntry, an array of replies, and an error
//NOTE, if the entry is set to invisible it is skipped
func LoadComments(eclient *elastic.Client, entryID string, lowerBound int, upperBound int) (types.JournalEntry, []types.JournalEntry, error) {
	var entries []types.JournalEntry
	var parent types.JournalEntry
	var start int
	var finish int

	if lowerBound < 0 {
		return parent, entries, errors.New("Lower Bound limit is out of bounds")
	}

	parent, err := ConvertEntryToJournalEntry(eclient, entryID)
	if err != nil {
		return parent, entries, err
	}
	if upperBound == -1 {
		finish = 0
	} else if len(parent.Element.ReplyIDs)-upperBound < 0 {
		finish = 0
	} else {
		finish = len(parent.Element.ReplyIDs) - upperBound
	}

	start = (len(parent.Element.ReplyIDs) - 1) - lowerBound
	for i := start; i > finish; i-- {
		jEntry, err := ConvertEntryToJournalEntry(eclient, parent.Element.ReplyIDs[i])
		if err != nil {
			return parent, entries, err
		}

		if !jEntry.Element.Visible && finish > 0 {
			finish--
			continue
		}

		entries = append(entries, jEntry)
	}

	return parent, entries, err
}

//ConvertEntryToJournalEntry ... load all relevant data of a single entry into a journal entry struct
//Requires entry docID
//Returns entry as a type JournalEntry and an error
func ConvertEntryToJournalEntry(eclient *elastic.Client, entryID string) (types.JournalEntry, error) {
	var newJournalEntry types.JournalEntry

	newJournalEntry.ElementID = entryID

	entry, err := get.GetEntryByID(eclient, entryID)
	if err != nil {
		return newJournalEntry, err
	}
	newJournalEntry.Element = entry
	newJournalEntry.NumShares = len(entry.ShareIDs)
	newJournalEntry.NumLikes = len(entry.Likes)
	newJournalEntry.NumReplies = len(entry.ReplyIDs)

	usr, err := get.GetUserByID(eclient, entry.PosterID)
	if err != nil {
		return newJournalEntry, err
	}
	newJournalEntry.FirstName = usr.FirstName
	newJournalEntry.LastName = usr.LastName
	newJournalEntry.Image = usr.Avatar

	if entry.Classification == 2 {
		newJournalEntry.ReferenceElement, err = ConvertEntryToJournalEntry(eclient, entry.ReferenceEntry)
	}

	return newJournalEntry, err
}

//ModifyDescription ... CHANGES A SPECIFIC USER'S DESCRIPTION
//Requires the target user's docID and the new description
//Returns an error
func ModifyDescription(eclient *elastic.Client, userID string, newDescription string) error {

	usr, err := get.GetUserByID(eclient, userID)

	if err != nil {
		return err
	}

	usr.Description = []rune(newDescription)

	retErr := post.UpdateUser(eclient, userID, "Description", usr)
	return retErr

}

//UserNewEntry ... CREATES AN ORIGINAL POST FROM A USER
//Requires the user's docID and the content of the post
//Returns a the new entry's docID, and an error
func UserNewEntry(eclient *elastic.Client, userID string, newContent []rune) (string, error) {
	var entryID string

	createdEntry := types.Entry{}
	createdEntry.PosterID = userID
	createdEntry.Classification = 0
	createdEntry.Content = newContent
	createdEntry.TimeStamp = time.Now()
	createdEntry.Visible = true

	entryID, err := post.IndexEntry(eclient, createdEntry)
	if err != nil {
		return entryID, err
	}
	err = post.AppendEntryID(eclient, userID, entryID)

	return entryID, err

}

//UserReplyEntry ... CREATES A REPLY ENTRY FROM A USER
//Requires the user's docID, the parent entry docID and the content of the post
//Returns an error
func UserReplyEntry(eclient *elastic.Client, userID string, entryID string, content []rune) error {

	var newReply types.Entry
	newReply.PosterID = userID
	newReply.Content = content
	newReply.ReferenceEntry = entryID
	newReply.TimeStamp = time.Now()
	newReply.Classification = 1

	replyID, err := post.IndexEntry(eclient, newReply)
	if err != nil {
		return err
	}

	err = post.AppendEntryID(eclient, userID, replyID)
	if err != nil {
		return err
	}

	err = post.AppendReplyID(eclient, entryID, replyID)
	return err
}

//UserShareEntry ... CREATES A SHARED ENTRY FROM A USER
//Requires the user's docID, the parent entry docID and the content of the post
//Returns an error
func UserShareEntry(eclient *elastic.Client, userID string, entryID string, content []rune) error {

	var newReply types.Entry
	newReply.PosterID = userID
	newReply.Content = content
	newReply.ReferenceEntry = entryID
	newReply.TimeStamp = time.Now()
	newReply.Classification = 2

	replyID, err := post.IndexEntry(eclient, newReply)
	if err != nil {
		return err
	}

	err = post.AppendEntryID(eclient, userID, replyID)
	if err != nil {
		return err
	}

	err = post.AppendShareID(eclient, entryID, replyID)
	return err
}

//UserLikeEntry ... ALLOWS A USER TO LIKE AN ENTRY
//Requires the entry's docID, and docID of the person who is liking the entry
//Returns an error
func UserLikeEntry(eclient *elastic.Client, entryID string, likerID string) error {

	err := post.AppendLike(eclient, entryID, likerID)
	if err != nil {
		return err
	}

	err = post.AppendLikedEntryID(eclient, likerID, entryID)
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
	err := post.DeleteLike(eclient, entryID, likerID)
	if err != nil {
		return err
	}

	//DeleteLikedEntryID deletes from usr
	err = post.DeleteLikedEntryID(eclient, likerID, entryID)
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
	entry, err := get.GetEntryByID(eclient, entryID)
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
	usr, err := get.GetUserByID(eclient, usrID)
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
	err = post.AppendCollReq(eclient, usrID, requestedUserID, true)
	if err != nil {
		return err
	}

	err = post.AppendCollReq(eclient, requestedUserID, requestedUserID, false)
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
	err := post.UpdateUser(eclient, userID, "QuickLinks", lynx)
	return err
}

//UpdateUserTags ... REPLACES THE ENTIRETY OF A USER'S TAGS WITH AN UPDATED LIST
//Requires the target user's docID and an updated array of strings
//Returns an error
func UpdateUserTags(eclient *elastic.Client, userID string, tags []string) error {
	err := post.UpdateUser(eclient, userID, "Tags", tags)
	return err
}

//UserFollow ... ALLOWS A USER TO FOLLOW SOMEONE ELSE
//Requires the follower's docID and the followed docID
//Returns an error
func UserFollow(eclient *elastic.Client, usrID string, followID string) error {
	//true = append to following
	followErr := post.AppendFollow(eclient, usrID, followID, true)
	if followErr != nil {
		return followErr
	}
	//false = append to followers
	followingErr := post.AppendFollow(eclient, followID, usrID, false)
	if followingErr != nil {
		return followingErr
	}

	return nil
}

//UserUnfollow ... ALLOWS A USER TO UNFOLLOW SOMEONE ELSE
//Requires the follower's docID and the followed docID
//Returns an error
func UserUnfollow(eclient *elastic.Client, usrID string, followID string) error {
	err := post.DeleteFollow(eclient, usrID, followID, true)
	if err != nil {
		return err
	}
	err = post.DeleteFollow(eclient, followID, usrID, false)
	return err
}

//IsFollowed ... CHECKS IF A USER IS FOLLWING SOMEONE ELSE
//Requires the follower's docID and the the potential followed docID
//Returns an error
func IsFollowed(eclient *elastic.Client, usrID string, viewerID string) (bool, error) {
	isFollowed := false
	user, err := get.GetUserByID(eclient, usrID)
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

	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return -1, err
	}
	if whichOne {
		return len(usr.Following), nil
	}

	return len(usr.Followers), nil

}
