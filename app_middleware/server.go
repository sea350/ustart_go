package main

import (
	"flag"
	"net/http"

	//"github.com/sea350/ustart_go/middleware/fail"
	//inbox "github.com/mrb721/ustart_app/ustartapp/middleware/inbox"
	email "github.com/sea350/ustart_go/app_middleware/email"
	login "github.com/sea350/ustart_go/app_middleware/login"
	profile "github.com/sea350/ustart_go/app_middleware/profile"
	project "github.com/sea350/ustart_go/app_middleware/project"
	search "github.com/sea350/ustart_go/app_middleware/search"
	settings "github.com/sea350/ustart_go/app_middleware/settings"
	signup "github.com/sea350/ustart_go/app_middleware/signup"
	fail "github.com/sea350/ustart_go/middleware/fail"
	/*profile "github.com/sea350/ustart_go/middleware/profile"
	project "github.com/sea350/ustart_go/middleware/project"
	registration "github.com/sea350/ustart_go/middleware/registration"
	settings "github.com/sea350/ustart_go/middleware/settings"
	widget "github.com/sea350/ustart_go/middleware/widget"*/)

var currentPort = "5003"

func main() {
	flag.Parse()

	/*
		Lines 18-19 handle the static file locating
		If we wanted to reorganize file/folder locations, this is one of 3 things that would have to change
		In executeTemplates you will need to make the same changes
		The other being the relative link on the actual html pages
	*/
	//fs := http.FileServer(http.Dir("/home/rr2396/www/"))
	_, _ = http.Get("http://ustart.today:" + currentPort + "/KillUstartPlsNoUserinoCappucinoDeniro")
	fs := http.FileServer(http.Dir("/ustart/aws_start/"))
	http.Handle("/www/", http.StripPrefix("/www/", fs))
	http.Handle("/aws_start/", http.StripPrefix("/aws_start/", fs))

	http.HandleFunc("/404/", fail.Fail)
	http.HandleFunc("/KillUstartPlsNoUserinoCappucinoDeniro", fail.KillSwitch)

	/*
		The following are all the handlers we have so far.
	*/

	//http.HandleFunc("/404/", fail.Fail)

	//LOGIN & LOGOUT
	//http.HandleFunc("/Inbox/", inbox.Inbox)
	//http.HandleFunc("/loginerror/", login.Error)
	//http.HandleFunc("/", login.Home)
	http.HandleFunc("/login/", login.Handler)
	http.HandleFunc("/signup/", signup.Handler)
	http.HandleFunc("/settings/", settings.Handler)
	http.HandleFunc("/profile/", profile.Handler)
	http.HandleFunc("/project/", project.Handler)
	http.HandleFunc("/search/", search.Handler)
	http.HandleFunc("/email/", email.Handler)

	//http.HandleFunc("/logout/", login.Logout)

	// USER PROFILE AND INTERACTIONS
	/*http.HandleFunc("/callme/", profile.Follow)
	http.HandleFunc("/Like", profile.Like)
	http.HandleFunc("/getComments/", profile.GetComments)
	http.HandleFunc("/shareComments/", profile.ShareComments)
	http.HandleFunc("/ShareComment", profile.ShareComment2)
	http.HandleFunc("/AddComment", profile.AddComment)
	http.HandleFunc("/AddComment2", profile.AddComment2)
	http.HandleFunc("/loadWall/", profile.WallLoad)
	http.HandleFunc("/addPost/", profile.WallAdd)
	http.HandleFunc("/profile/", profile.ViewProfile)
	http.HandleFunc("/addSkill/", profile.AddTag)
	http.HandleFunc("/deleteSkill/", profile.DeleteTag)
	http.HandleFunc("/addLink/", profile.AddQuickLink)
	http.HandleFunc("/deleteLink/", profile.DeleteQuickLink)
	http.HandleFunc("/deletePost/", profile.DeleteWallPost)
	http.HandleFunc("/deleteModal/", profile.GenerateDeleteModal)
	http.HandleFunc("/getPostComments/", profile.PostComments)
	http.HandleFunc("/followers/", profile.FollowersPage)
	http.HandleFunc("/following/", profile.FollowersPage)

	//WIDGET INTERACTIONS
	http.HandleFunc("/addWidget/", widget.AddWidget)
	http.HandleFunc("/addProjectWidget/", widget.AddProjectWidget)
	http.HandleFunc("/deleteWidget/", widget.DeleteWidgetProfile)
	http.HandleFunc("/deleteProjectWidget/", widget.DeleteWidgetProject)
	http.HandleFunc("/deleteLinkFromWidget/", widget.EditWidgetDataDelete)
	http.HandleFunc("/sortUserWidgets/", widget.SortUserWidgets)

	//PROJECT INTERACTIONS
	http.HandleFunc("/Projects/", project.ProjectsPage)
	http.HandleFunc("/MyProjects/", project.MyProjects)
	http.HandleFunc("/CreateProjectPage/", project.CreateProjectPage)
	http.HandleFunc("/CreateProject/", project.CreateProject)
	http.HandleFunc("/UpdateProjectTags/", project.UpdateTags)
	http.HandleFunc("/AddProjectLink/", project.AddQuickLink)
	http.HandleFunc("/DeleteProjectLink/", project.DeleteQuickLink)
	http.HandleFunc("/NewMembers/", project.ManageProjects)
	http.HandleFunc("/LoadJoinRequests/", project.LoadJoinRequests)
	http.HandleFunc("/RequestToJoin/", project.RequestToJoin)
	http.HandleFunc("/AcceptJoinRequest/", project.AcceptJoinRequest)
	http.HandleFunc("/RejectJoinRequest/", project.RejectJoinRequest)

	//SETTINGS CHANGES
	http.HandleFunc("/Settings/", settings.Settings)
	http.HandleFunc("/ImageUpload/", settings.ImageUpload)
	http.HandleFunc("/changeName/", settings.ChangeName)
	http.HandleFunc("/changePass/", settings.ChangePassword)
	http.HandleFunc("/changeLoc/", settings.ChangeLocation)
	http.HandleFunc("/changeEDU/", settings.ChangeEDU)
	http.HandleFunc("/UpdateDescription/", settings.ChangeContactAndDescription)
	http.HandleFunc("/BannerUpload/", settings.BannerUpload)
	http.HandleFunc("/ProjectSettings/", settings.Project)
	http.HandleFunc("/projectBannerUpload/", settings.ProjectBannerUpload)
	http.HandleFunc("/projectName/", settings.ChangeNameAndDescription)
	http.HandleFunc("/projectLocation/", settings.ProjectLocation)
	http.HandleFunc("/projectCategory/", settings.ProjectCategory)
	http.HandleFunc("/projectCustomURL/", settings.ProjectCustomURL)
	http.HandleFunc("/leaveProject/", settings.LeaveProject)
	http.HandleFunc("/projectLogo/", settings.ProjectLogo)
	http.HandleFunc("/changeMemberClass/", settings.ChangeMemberClass)

	//REGISTRATIONS

	http.HandleFunc("/Registration/Type/", registration.RegisterType)
	http.HandleFunc("/registrationcomplete/", registration.Complete)
	http.HandleFunc("/welcome/", registration.Registration)

	//EVENT
	http.HandleFunc("/Event/", event.ViewEvent)
	http.HandleFunc("/MyEvents/", event.MyEvents)
	http.HandleFunc("/CreateEventPage/", event.CreateEventPage)
	http.HandleFunc("/UpdateEventTags/", event.UpdateEventTags)
	http.HandleFunc("/AddEventLink/", event.AddEventQuickLink)
	http.HandleFunc("/DeleteEventLink/", event.DeleteEventQuickLink)
	http.HandleFunc("/NewUsers/", event.ManageEvents)
	http.HandleFunc("/LoadGuestJoinRequests/", event.LoadGuestJoinRequests)
	http.HandleFunc("/LoadMemberJoinRequests/", event.LoadMemberJoinRequests)
	http.HandleFunc("/GuestRequestToJoin/", event.GuestRequestToJoin)
	http.HandleFunc("/MemberRequestToJoin/", event.MemberRequestToJoin)
	http.HandleFunc("/AcceptGuestJoinRequest/", event.AcceptGuestJoinRequest)
	http.HandleFunc("/AcceptMemberJoinRequest/", event.AcceptMemberJoinRequest)
	http.HandleFunc("/RejectGuestJoinRequest/", event.RejectEventGuestJoinRequest)
	http.HandleFunc("/RejectMemberJoinRequest/", event.RejectEventMemberJoinRequest)
	http.HandleFunc("/EventMakeEntry/", event.MakeEventEntry)
	http.HandleFunc("/AjaxLoadEventEntries/", event.AjaxLoadEventEntries)
	http.HandleFunc("/AjaxDeleteEventEntry/", event.AjaxDeleteEventEntry)
	http.HandleFunc("/StartEvent/", event.StartEvent)
	http.HandleFunc("/ManageEvents/", event.ManageEvents)
	http.HandleFunc("/AddEvent/", event.AddEvent)
	http.HandleFunc("/EventsPage/", event.EventsPage)*/
	http.ListenAndServe(":5003", nil)
}
