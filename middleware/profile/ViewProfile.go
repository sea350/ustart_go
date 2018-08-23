package profile

import (
	"log"
	"net/http"
	"strings"

	"github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"

	getFollow "github.com/sea350/ustart_go/get/follow"
	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
)

//ViewProfile ... Loads data relevant to profile page and displays it
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	pageUserName := strings.ToLower(r.URL.Path[9:])

	userstruct, errMessage, _, err := uses.UserPage(client.Eclient, pageUserName, docID.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		log.Println("User Error: " + errMessage)
		http.Redirect(w, r, "/404/", http.StatusFound)
		return
	}

	widgets, errs := uses.LoadWidgets(client.Eclient, userstruct.UserWidgets)
	if len(errs) != 0 {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("one or more errors have occured in loading widgets")
		log.Println(errs)
	}

	jEntries, err := uses.LoadEntries(client.Eclient, userstruct.EntryIDs, docID.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	followingState := false
	_, follDoc, err := getFollow.ByID(client.Eclient, session.Values["DocID"].(string))
	exist1, _ := follDoc.UserFollowers[session.Values["DocID"].(string)]
	exist2, _ := follDoc.ProjectFollowers[session.Values["DocID"].(string)]
	exist3, _ := follDoc.EventFollowers[session.Values["DocID"].(string)]
	if exist1 || exist2 || exist3 {
		followingState = true
	}

	var ClassYear string
	if userstruct.Class == 1 {
		ClassYear = "Freshman"
	}
	if userstruct.Class == 2 {
		ClassYear = "Sophomore"
	}
	if userstruct.Class == 3 {
		ClassYear = "Junior"
	}
	if userstruct.Class == 4 {
		ClassYear = "Senior"
	}
	if userstruct.Class == 5 {
		ClassYear = "Graduate"
	}
	if userstruct.Class == 6 {
		ClassYear = "Post-Graduate"
	}
	bday := userstruct.Dob.String()
	month := bday[5:7]
	day := bday[8:10]
	year := bday[0:4]
	birthdayline := month + "/" + day + "/" + year

	viewingDOC, err := get.IDByUsername(client.Eclient, strings.ToLower(pageUserName))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	temp := string(userstruct.Description)

	numberFollowers := len(follDoc.UserFollowers) + len(follDoc.ProjectFollowers) + len(follDoc.EventFollowers)
	numberFollowing := len(follDoc.UserFollowing) + len(follDoc.ProjectFollowing) + len(follDoc.EventFollowing)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	var projHeads []types.FloatingHead
	for _, projID := range userstruct.Projects {
		head, err := uses.ConvertProjectToFloatingHead(client.Eclient, projID.ProjectID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			continue
		}
		head.Followed = projID.Visible
		projHeads = append(projHeads, head)
	}

	cs := client.ClientSide{UserInfo: userstruct, Wall: jEntries, Birthday: birthdayline, Class: ClassYear, Description: temp, Followers: numberFollowers, Following: numberFollowing, Page: viewingDOC, FollowingStatus: followingState, Widgets: widgets, ListOfHeads: projHeads}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "profile-nil", cs)
}
