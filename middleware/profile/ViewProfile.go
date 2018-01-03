package profile

import (
	"net/http"
	uses "github.com/sea350/ustart_go/uses"
	elastic "gopkg.in/olivere/elastic.v5"
	// types "github.com/sea350/ustart_go/types"
	get "github.com/sea350/ustart_go/get"
	"fmt"
	"github.com/gorilla/sessions"
	client "github.com/sea350/ustart_go/middleware/clientstruct"
)

var eclient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))
var store = sessions.NewCookieStore([]byte("RIU3389D1")) // code 




func ViewProfile (w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if (test1 == nil){
		http.Redirect(w, r, "/~", http.StatusFound)
	}

	userstruct,_, followbool,_ := uses.UserPage(eclient,r.URL.Path[9:],session.Values["DocID"].(string))
	jEntries, err5 := uses.LoadEntries(eclient,userstruct.EntryIDs)
	if (err5 != nil){
		fmt.Println(err5);
	}
	followingState := "no"
	if (followbool == true){
		followingState = "yes"
	}
	if (followbool == false){
	}

	var ClassYear string 
	if (userstruct.Class == 1){
		ClassYear = "Freshman"
	}
	if (userstruct.Class == 2){
		ClassYear = "Sophomore"
	}
	if (userstruct.Class == 3){
		ClassYear = "Junior"
	}
	if (userstruct.Class == 4){
		ClassYear = "Senior"
	}
	if (userstruct.Class == 5){
		ClassYear = "Graduate"
	}
	if (userstruct.Class == 6){
		ClassYear = "Post-Graduate"
	}
	bday := userstruct.Dob.String()
	month := bday[5:7]
	day := bday[8:10]
	year := bday[0:4]
	birthdayline := month+"/"+day+"/"+year
	cs := client.ClientSide{UserInfo:userstruct, DOCID: session.Values["DocID"].(string),Birthday: birthdayline,Class:ClassYear} //Class:ClassYear}
	viewingDOC, errID := get.GetIDByUsername(eclient, r.URL.Path[9:])
	if (errID != nil){
		fmt.Println(errID);
	}

	temp := string(cs.UserInfo.Description) 

	numberFollowing,errnF := uses.NumFollow(eclient, session.Values["DocID"].(string),true)
	if (errnF != nil){
		fmt.Println(errnF);
	}
	numberFollowers,errnF2 := uses.NumFollow(eclient, session.Values["DocID"].(string),false)
	if (errnF2 != nil){
		fmt.Println(errnF2);
	}
	cs = client.ClientSide{UserInfo:userstruct, Wall: jEntries, DOCID: session.Values["DocID"].(string),Birthday: birthdayline,Class:ClassYear, Description:temp,Followers:numberFollowers,Following:numberFollowing, Page:viewingDOC,FollowingStatus:followingState}


	client.RenderTemplate(w,"template2-nil",cs)
	client.RenderTemplate(w,"profile-nil",cs)
}


