package uses

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart_go/types"
	post "github.com/sea350/ustart_go/post"
	get "github.com/sea350/ustart_go/get"
	//"net/http"
	//"io"
	//"fmt"
	//"bytes"
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
	newUsr.UndergradSchool = school
	newUsr.Majors = major
	newUsr.Dob = bday

	newLoc := types.LocStruct{}
	newLoc.Country = country
	newLoc.State = state
	newLoc.City = city
	newLoc.Zip = zip
	newUsr.Location = newLoc

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

	return loginSucessful, userSession, err

}


func UserPage(eclient *elastic.Client, userID string)(types.User, []types.WallEntry, error){
	//Returns all information relevant to the user page given an userid
	var entries []types.WallEntry
	maxPull := 20
	maxReplyPull := 5
	counter := 0
	replyCounter := 0


	usr, err := get.GetUserById(eclient, userID)
	if (err!=nil){return usr, entries, err}


	for _, i := range usr.EntryIDs {
		entry, err := get.GetEntryById(eclient, i)
		if (err!=nil){return usr, entries, err}
		var newEntry types.WallEntry
		newEntry.Element = entry
		if(len(entry.ReplyIDs) > 0){
			for _, j := range entry.ReplyIDs{
				reply, err := get.GetEntryById(eclient, j)
				if (err!=nil){return usr, entries, err}
				newEntry.RepliesArray = append(newEntry.RepliesArray, reply)
				replyCounter += 1
				if (replyCounter > maxReplyPull){break}
			}
		}

		entries = append(entries, newEntry)
		counter += 1
		if (counter > maxPull){break}
	}

	return usr, entries, err
}



/*func ModifyDescription(eclient *elastic.Client, userID string, newDescription string)error{


	usr, err:= get.GetUserById(eclient, userID)

	if(err!=nil){}

	usr.Description = newDescription
    
	return post.UpdateUser(eclient,userID,usr)
}
*/


func userCreatesEntry(eclient *elastic.Client, userID string, newContent []rune) error{
	createdEntry:= types.Entry{}
	createdEntry.PosterId := userID
	createdEntry.Classification:= 0
	createdEntry.Content := newContent
	createdEntry.TimeStamp := time.Now()
	createdEntryVisible := true

	err:=post.IndexEntry(eclient, createdEntry)

	return err

}
