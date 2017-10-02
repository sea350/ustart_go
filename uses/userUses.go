package uses

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart_go/types"
	post "github.com/sea350/ustart_go/post"
	get "github.com/sea350/ustart_go/get"
	//uses "github.com/sea350/ustart_go/uses"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)



func SignUpBasic(eclient *elastic.Client, email string, password []byte, fname string, lname string, country string, state string, city string, zip string, school string, major []string, bday time.Time, currYear string) error{
	//A basic user signup process
	//Returns an error if there was a problem with database submission
	//Or if email is in use

	inUse, err := get.EmailInUse(eclient, email)
	if(err!=nil){return err}
	if(inUse){return errors.New("Email is in use.")}

	newUsr:=types.User{}
	newUsr.FirstName = fname
	newUsr.LastName = lname
	newUsr.Email = email
	newUsr.Username = get.EmailToUsername(email)
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
	if currYear == "Freshman"{
		newUsr.Class = 0
	}else if currYear == "Sophomore"{
		newUsr.Class = 1
	}else if currYear == "Junior"{
		newUsr.Class = 2
	}else if currYear == "Senior"{
		newUsr.Class = 3
	}else if currYear == "Graduate"{
		newUsr.Class = 4
	}else{
		newUsr.Class = 5
	}


	retErr:=post.IndexUser(eclient,newUsr) 

	return retErr
}





func Login(eclient *elastic.Client, userEmail string, password []byte)(bool, types.SessionUser, error) {
	//Contains all logic for a user login including email and password check
	//Returns a bool if the login was sucessful, 
	//the user information required by a session
	//and an error


	var loginSucessful bool = false
	var userSession types.SessionUser

	inUse, err := get.EmailInUse(eclient, userEmail)
	if(err!=nil){return loginSucessful, userSession, err}
	if(!inUse){
		err := errors.New("Invalid Email")
		return loginSucessful, userSession, err}

	usr,err := get.GetUserByEmail(eclient,userEmail)
	if(err!=nil){return loginSucessful, userSession, err}
	//if(!bytes.Equal(password, usr.Password)){return loginSucessful, userSession, errors.New(
	//	"Email And Password Do Not Match")}

	passErr:=bcrypt.CompareHashAndPassword(usr.Password,password)

	if(passErr!=nil){return false, userSession, passErr}

	uID,err := get.GetUserIDByEmail(eclient, userEmail)
	if(err!=nil){return loginSucessful, userSession, err}


	loginSucessful = true
	userSession.FirstName = usr.FirstName
	userSession.LastName = usr.LastName
	userSession.Email = userEmail
	userSession.DocID = uID
	userSession.Username = usr.Username

	return loginSucessful, userSession, err

}


func UserPage(eclient *elastic.Client, username string, viewerID string)(types.User, []types.JournalEntry, string, bool, error){
	//Returns all information relevant to the user page given a username
	maxPull := 20
	//maxReplyPull := 5
	counter := 0
	//replyCounter := 0
	var usr types.User

	var entries []types.JournalEntry
	var isFollowed bool
	userID, err := get.GetIDByUsername(eclient, username)
	if (err!=nil){return usr, entries, userID, isFollowed, err}


	userID, err = get.GetIDByUsername(eclient, username)
	if (err!=nil){return usr, entries, userID, isFollowed, err}
	usr, err = get.GetUserByID(eclient, userID)
	if (err!=nil){return usr, entries, userID, isFollowed, err}

	viewer, err := get.GetUserByID(eclient, viewerID)
	if (err!=nil){return usr, entries, userID, isFollowed, err}

	for _, element := range viewer.Following{
		if (element == userID) {
			isFollowed = true
			break
		}
	}


	for _, i := range usr.EntryIDs {
		//goes through the user's entries
		entry, err := get.GetEntryByID(eclient, i)
		if (err!=nil){return usr, entries, userID, isFollowed, err}
		if (!entry.Visible){continue}//checks if entry is visible
		//if invisible, then skip

		var newEntry types.JournalEntry
		newEntry.Element = entry
		newEntry.FirstName = usr.FirstName
		newEntry.LastName = usr.LastName
		newEntry.NumReplies = len(entry.ReplyIDs)
		newEntry.NumLikes = len(entry.Likes)
		newEntry.NumShares = len(entry.ShareIDs)


		/*if(len(entry.ReplyIDs) > 0){
			for _, j := range entry.ReplyIDs{
				//loops through each entries comment entries
				reply, err := get.GetEntryByID(eclient, j)
				if (err!=nil){return usr, entries, userID, isFollowed, err}
				if (!reply.Visible){continue}//checks if entry is visible
				//if invisible, then skip
				replyUsr, err := get.GetUserByID(eclient, reply.PosterID)
				if (err!=nil){return usr, entries, userID, isFollowed, err}

				var newReply types.JournalEntry
				newReply.Element = reply
				newReply.FirstName = replyUsr.FirstName
				newReply.LastName = replyUsr.LastName
				newReply.NumReplies = len(entry.ReplyIDs)
				newReply.NumLikes = len(entry.Likes)
				newReply.NumShares = len(entry.ShareIDs)

				newEntry.RepliesArray = append(newEntry.RepliesArray, newReply)
				replyCounter += 1
				if (replyCounter > maxReplyPull){
					replyCounter = 0
					break
				}
			}
		}
		*/
		//check if invis
		entries = append(entries, newEntry)
		counter += 1
		if (counter > maxPull){break}
	}

	return usr, entries, userID, isFollowed, err
}



func ModifyDescription(eclient *elastic.Client, userID string, newDescription string)error{


	usr, err := get.GetUserByID(eclient, userID)

	if(err!=nil){return err}

	usr.Description = []rune(newDescription)
    
    retErr:=post.UpdateUser(eclient,userID,"Description",usr)
	return retErr

}

func UserCreatesEntry(eclient *elastic.Client, userID string, newContent []rune) error{
	createdEntry:= types.Entry{}
	createdEntry.PosterID = userID
	createdEntry.Classification= 0
	createdEntry.Content = newContent
	createdEntry.TimeStamp = time.Now()
	createdEntry.Visible = true


	//usr, err := get.GetUserByID(eclient,userID)


	entryID,err:=post.IndexEntry(eclient,createdEntry)
	if (err != nil) {return err}
	err = post.AppendEntryID(eclient, userID, entryID)

	return err

}

func UserLike(eclient *elastic.Client, entryID string,userID string)error{
	_, err:= get.GetEntryByID(eclient,entryID)

	if (err!=nil){return errors.New("Post does not exist")}

	//var usrLike types.Like
	//usrLike.UserID = userID
	//usrLike.TimeStamp = time.Now()

	return post.AppendLike(eclient,entryID, userID)  

}




/*func UserFollow(eclient *elastic.Client, usrID string, usrFollowID string) error{

	usr, err := get.GetUserByID(eclient,usrID)
	if (err!=nil){return err}


	for _, element := range usr.Following{
		if (element == usrID) {return errors.New("You are already following this user")}
	}

	usr.Following = append(usr.Following, usrFollowID)
	//append new follow to user
	//append new follow by user


	return err
}

func UserUnfollow(eclient *elastic.Client, usrID string, followID string) error{

	usr, err := get.GetUserByID(eclient,usrID)
	if (err!=nil){return err}

	var newFollows []string

	for _, element := range usr.Following{
		if (element != followID) {
			newFollows = append(newFollows, element)
		}
		if (isFollowing) {return errors.New("You are already following this user")}
	}

	usr.Following = append(usr.Following, usrFollowID)
	post.UpdateUser(eclient, usrID, "Following", usr.Following)//CHANGE TO APPEND!!!!!

	return err
}
*/
func RequestColleague(eclient *elastic.Client, usrID string, requestedUserID string) error{
	usr, err := get.GetUserByID(eclient,usrID) 
	if (err!=nil){return err}
	//_, errRequest := get.GetUserByID(eclient,requestedUserID)

	alreadyRequested := false
	for _, element := range usr.ReceivedCollReq{
		if (element == usrID) {alreadyRequested=true}
		if (alreadyRequested) {return errors.New("You have already requested this user")}
	}

	usr.Colleagues = append(usr.Colleagues, requestedUserID)


	return post.ReindexUser(eclient, usrID,usr)
}

func ReplyToColleagueRequest(eclient *elastic.Client, usrID string, requestedUserID string, reply bool) error{
	if (reply == true){
		usr, err := get.GetUserByID(eclient, usrID)
		if(err!=nil){return err}
		//usrRequest, _ := get.GetUserByID(eclient, requestedUserID)
		usr.Colleagues = append(usr.Colleagues, requestedUserID)

		return post.ReindexUser(eclient, usrID,usr)
	}

	return errors.New("Could not reply to colleague request")

}

func UserNewTextEntry (eclient *elastic.Client, userID string, content []rune) error{

	var newEntry types.Entry
	newEntry.PosterID = userID
	newEntry.Content = content
	newEntry.TimeStamp = time.Now()
	newEntry.Visible = true
	newEntry.Classification = 0

	//remove this after append is made!!!!!
	usr, err :=  get.GetUserByID(eclient, userID)
	if (err != nil){return err}
	//-----------

	id, err := post.IndexEntry(eclient, newEntry)
	if (err != nil){return err}

	//modify this after append is made
	err = post.UpdateUser(eclient, userID, "EntryIDs", append(usr.EntryIDs, id))
	return err

}

func UserNewReplyEntry (eclient *elastic.Client, userID string, content []rune, hostEntryID string) error{

	var newEntry types.Entry
	newEntry.PosterID = userID
	newEntry.Content = content
	newEntry.TimeStamp = time.Now()
	newEntry.Visible = true
	newEntry.Classification = 1
	newEntry.ReferenceEntry = hostEntryID

	//remove this after append is made!!!!!
	usr, err :=  get.GetUserByID(eclient, userID)
	if (err != nil){return err}
	entry, err := get. GetEntryByID(eclient, hostEntryID)
	if (err != nil){return err}
	//-----------

	id, err := post.IndexEntry(eclient, newEntry)
	if (err != nil){return err}

	//modify this after append is made
	err = post.UpdateUser(eclient, userID, "EntryIDs", append(usr.EntryIDs, id))
	if (err != nil){return err}
	err = post.UpdateEntry(eclient, hostEntryID, "ReplyIDs", append(entry.ReplyIDs, id))
	return err

}

func UserNewShareEntry (eclient *elastic.Client, userID string, content []rune, hostEntryID string) error{

	var newEntry types.Entry
	newEntry.PosterID = userID
	newEntry.Content = content
	newEntry.TimeStamp = time.Now()
	newEntry.Visible = true
	newEntry.Classification = 2
	newEntry.ReferenceEntry = hostEntryID

	//remove this after append is made!!!!!
	usr, err :=  get.GetUserByID(eclient, userID)
	if (err != nil){return err}
	entry, err := get. GetEntryByID(eclient, hostEntryID)
	if (err != nil){return err}
	//-----------

	id, err := post.IndexEntry(eclient, newEntry)
	if (err != nil){return err}

	//modify this after append is made
	err = post.UpdateUser(eclient, userID, "EntryIDs", append(usr.EntryIDs, id))
	if (err != nil){return err}
	err = post.UpdateEntry(eclient, hostEntryID, "ShareIDs", append(entry.ShareIDs, id))
	return err

}

func UpdateUserLinks (eclient *elastic.Client, userID string, lynx []types.Link) error{
	err := post.UpdateUser(eclient, userID, "QuickLinks", lynx)
	return err
}

func UpdateUserTags (eclient *elastic.Client, userID string, tags []string) error{
	err := post.UpdateUser(eclient, userID, "Tags", tags)
	return err
}



func ToggleUserFollow (eclient *elastic.Client, followerID string, targetID string) error{
	follower, err := get.GetUserByID(eclient, followerID)
	if (err != nil) {return err}

	for element := range follower.Following{
		if (follower.Following[element] == targetID){
			//remove target id from follower
			//remove follower id from target
			
			return err
		}
	}

	//add target id to follower's followed
	//add follower id to target's followers

	return err

}


func ToggleUserLike (eclient *elastic.Client, likerID string, entryID string) error {

	liker, err := get.GetUserByID(eclient, likerID)
	if (err != nil) {return err}

	for _,element := range liker.LikedEntryIDs{
		if (element == entryID){
			//remove post id from liked list
			//remove liker id from entry's likes list
			
			return err
		}
	}

	//add post id to liker's liked list
	//add liker id to posts likes list

	return err

}

