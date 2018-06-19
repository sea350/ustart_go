package settings

import (
	"fmt"
	"net/http"

	"github.com/sea350/ustart_go/uses"

	get "github.com/sea350/ustart_go/get/project"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/project"
)

//LeaveProject ... lets a user leave a project
//If Rol
func LeaveProject(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	docID := test1.(string)

	leavingUser := r.FormValue("leaverID")
	projID := r.FormValue("projectID")
	newCreator := r.FormValue("newCreator")

	proj, err := get.ProjectByID(eclient, projID)
	if err != nil {
		fmt.Println("critical err middleware/settings/leaveproject line 34")
		fmt.Println(err)
		return
	}
	defer http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)

	var usrIndex int
	usrMembershipExists := false
	var leaverIndex int
	leaverExists := false
	var newCreatorIndex int
	newCreatorMemberShipExists := false
	numCreators := 0

	//getting acting user's membership credentials
	//finding new creator's index
	//also counting how many creators are set
	for i, mem := range proj.Members {
		if mem.MemberID == docID {
			usrIndex = i
			usrMembershipExists = true
		}
		if mem.MemberID == newCreator {
			newCreatorIndex = i
			newCreatorMemberShipExists = true
		}
		if mem.MemberID == leavingUser {
			leaverIndex = i
			leaverExists = true
		}
		if mem.Role == 0 {
			numCreators++
		}
	}

	//if membership not found do nothing
	if !usrMembershipExists {
		return
	}

	//if a user quits, ie removes themself
	if leavingUser == docID {
		//if the usr is the creator
		if proj.Members[usrIndex].Role == 0 {
			//if no new creator is appointed and there arent any creators to spare you cant do anything
			if newCreator == `` && numCreators < 2 {
				return
			} else if newCreator != `` { //if there is a new creator specified, appoint new creator and delete old one
				//if we couldnt find new creator as a member
				if !newCreatorMemberShipExists {
					return
				}
				proj.Members[newCreatorIndex].Role = 0
				err = post.UpdateProject(client.Eclient, projID, "Members", proj.Members)
				if err != nil {
					fmt.Println("err middleware/settings/leaveproject line 81")
					fmt.Println(err)
				}
				err = post.DeleteMember(client.Eclient, projID, docID)
				if err != nil {
					fmt.Println("err middleware/settings/leaveproject line 86")
					fmt.Println(err)
				}
			}
		}
		//if none of the previous conditions then you should be clear to remove them w/o issue
		//and if the user is not a leader no one cares, just delete them
		err = post.DeleteMember(client.Eclient, projID, docID)
		if err != nil {
			fmt.Println("err middleware/settings/leaveproject line 93")
			fmt.Println(err)
		}
	} else { //user is trying to remove someone else WARNING order matters
		//if there is no leaver to remove you cant do anything
		if !leaverExists {
			return
		}
		//if acting user doesnt have priveledge fuck em
		if !uses.HasPrivilege("member", proj.PrivilegeProfiles, proj.Members[usrIndex]) {
			return
		}
		//if a non creator is attempting to remove a creator fuck em
		if proj.Members[usrIndex].Role != 0 && proj.Members[leaverIndex].Role == 0 {
			return
		}

		//if none of the previous conditions you should be clear to remove the member?
		err = post.DeleteMember(client.Eclient, projID, leavingUser)
		if err != nil {
			fmt.Println("err middleware/settings/leaveproject line 122")
			fmt.Println(err)
		}
	}

}
