package main

import (
	"flag"
	"net/http"

	chat "github.com/sea350/ustart_go/middleware/chat"
	dash "github.com/sea350/ustart_go/middleware/dashboard"
	event "github.com/sea350/ustart_go/middleware/event"
	fail "github.com/sea350/ustart_go/middleware/fail"
	img "github.com/sea350/ustart_go/middleware/img"
	inbox "github.com/sea350/ustart_go/middleware/inbox"
	login "github.com/sea350/ustart_go/middleware/login"
	"github.com/sea350/ustart_go/middleware/notification"
	profile "github.com/sea350/ustart_go/middleware/profile"
	project "github.com/sea350/ustart_go/middleware/project"
	registration "github.com/sea350/ustart_go/middleware/registration"
	search "github.com/sea350/ustart_go/middleware/search"
	settings "github.com/sea350/ustart_go/middleware/settings"
	widget "github.com/sea350/ustart_go/middleware/widget"
)

var currentPort = "5002"

func main() {
	flag.Parse()
	/*
		Lines 18-19 handle the static file locating
		If we wanted to reorganize file/folder locations, this is one of 3 things that would have to change
		In executeTemplates you will need to make the same changes
		The other being the relative link on the actual html pages
	*/
	// fs := http.FileServer(http.Dir("/home/rr2396/www/"))
	_, _ = http.Get("http://ustart.today:" + currentPort + "/KillUstartPlsNoUserinoCappucinoDeniro")
	fs := http.FileServer(http.Dir("/ustart/ustart_front/"))
	// http.Handle("/www/", http.StripPrefix("/www/", fs))
	http.Handle("/ustart_front/", http.StripPrefix("/ustart_front/", fs))
	/*
		The following are all the handlers we have so fart.
	*/

	http.HandleFunc("/404/", fail.Fail)
	http.HandleFunc("/KillUstartPlsNoUserinoCappucinoDeniro", fail.KillSwitch)

	http.HandleFunc("/Inbox/", inbox.Inbox)

	//LOGIN & LOGOUT
	http.HandleFunc("/loginerror/", login.Error)
	http.HandleFunc("/", login.Home)
	http.HandleFunc("/profilelogin/", login.Login)
	http.HandleFunc("/logout/", login.Logout)
	http.HandleFunc("/unverified/", login.Unverified)

	// USER PROFILE AND INTERACTIONS
	http.HandleFunc("/profile/", profile.ViewProfile)
	http.HandleFunc("/callme/", profile.Follow)
	http.HandleFunc("/Like/", profile.Like)
	http.HandleFunc("/getComments/", profile.GetComments)
	http.HandleFunc("/AddComment/", profile.AddComment)
	http.HandleFunc("/AddComment2/", profile.AddComment2)
	http.HandleFunc("/addPost/", profile.WallAdd)
	http.HandleFunc("/addSkill/", profile.AddTag)
	http.HandleFunc("/deleteSkill/", profile.DeleteTag)
	http.HandleFunc("/addLink/", profile.AddQuickLink)
	http.HandleFunc("/deleteLink/", profile.DeleteQuickLink)
	http.HandleFunc("/deletePost/", profile.DeletePost)
	http.HandleFunc("/editPost/", profile.EditPost)
	http.HandleFunc("/shareEntry/", profile.ShareEntry)
	http.HandleFunc("/deleteModal/", profile.GenerateDeleteModal)
	http.HandleFunc("/getPostComments/", profile.PostComments)
	http.HandleFunc("/followers/", profile.FollowersPage)
	http.HandleFunc("/following/", profile.FollowersPage)
	http.HandleFunc("/toggleProjectInvis/", profile.AjaxChangeProjVisibility)
	http.HandleFunc("/toggleEventInvis/", profile.AjaxChangeEventVisibility)

	http.HandleFunc("/testWall/", profile.TestWallPage)
	http.HandleFunc("/ajaxUserEntries/", profile.AjaxLoadUserEntries)

	//WIDGET INTERACTIONS
	http.HandleFunc("/addWidget/", widget.AddWidget)
	http.HandleFunc("/addProjectWidget/", widget.AddProjectWidget)
	http.HandleFunc("/addEventWidget/", widget.AddEventWidget)
	http.HandleFunc("/deleteWidget/", widget.DeleteWidgetProfile)
	http.HandleFunc("/deleteProjectWidget/", widget.DeleteWidgetProject)
	http.HandleFunc("/deleteEventWidget/", widget.DeleteWidgetEvent)
	http.HandleFunc("/deleteLinkFromWidget/", widget.EditWidgetDataDelete)
	http.HandleFunc("/sortUserWidgets/", widget.SortUserWidgets)

	//PROJECT INTERACTIONS
	http.HandleFunc("/Projects/", project.ProjectsPage)
	http.HandleFunc("/MyProjects/", project.MyProjects)
	http.HandleFunc("/CreateProjectPage/", project.CreateProjectPage)
	http.HandleFunc("/UpdateProjectTags/", project.UpdateTags)
	http.HandleFunc("/AddProjectLink/", project.AddQuickLink)
	http.HandleFunc("/DeleteProjectLink/", project.DeleteQuickLink)
	http.HandleFunc("/NewMembers/", project.ManageProjects)
	http.HandleFunc("/LoadJoinRequests/", project.LoadJoinRequests)
	http.HandleFunc("/RequestToJoin/", project.RequestToJoin)
	http.HandleFunc("/AcceptJoinRequest/", project.AcceptJoinRequest)
	http.HandleFunc("/RejectJoinRequest/", project.RejectJoinRequest)
	http.HandleFunc("/ProjectMakeEntry/", project.MakeEntry)
	http.HandleFunc("/AjaxLoadProjectEntries/", project.AjaxLoadProjectEntries)
	http.HandleFunc("/AjaxDeleteProjectEntry/", project.AjaxDeleteEntry)
	http.HandleFunc("/AjaxUserFollowProjectToggle/", project.AjaxToggleFollow)

	//SETTINGS CHANGES
	http.HandleFunc("/Settings/", settings.Settings)
	http.HandleFunc("/ImageUpload/", settings.ImageUpload)
	http.HandleFunc("/changeName/", settings.ChangeName)
	http.HandleFunc("/changePass/", settings.ChangePassword)
	http.HandleFunc("/changeLoc/", settings.ChangeLocation)
	http.HandleFunc("/changeEDU/", settings.ChangeEDU)
	http.HandleFunc("/UpdateDescription/", settings.ChangeContactAndDescription)
	http.HandleFunc("/BannerUpload/", settings.BannerUpload)
	http.HandleFunc("/EventSettings/", settings.Event)
	http.HandleFunc("/eventBannerUpload/", settings.EventBannerUpload)
	http.HandleFunc("/eventName/", settings.EventChangeNameAndDescription)
	http.HandleFunc("/eventLocation/", settings.EventLocation)
	http.HandleFunc("/eventCategory/", settings.EventCategory)
	http.HandleFunc("/eventCustomURL/", settings.EventCustomURL)
	http.HandleFunc("/leaveEvent/", settings.LeaveEvent)
	http.HandleFunc("/leaveEventGuest/", settings.LeaveEventGuest)
	http.HandleFunc("/leaveEventMember/", settings.LeaveEventMember)
	http.HandleFunc("/eventLogo/", settings.EventLogo)
	http.HandleFunc("/changeEventMemberClass/", settings.ChangeEventMemberClass)
	http.HandleFunc("/eventHost/", settings.EventHost)
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
	http.HandleFunc("/Signup/", registration.Signup)
	http.HandleFunc("/Registration/Type/", registration.RegisterType)
	http.HandleFunc("/registrationcomplete/", registration.Complete)
	http.HandleFunc("/welcome/", registration.Registration)
	http.HandleFunc("/Activation/", registration.EmailVerification)
	http.HandleFunc("/ResetPassword/", registration.ResetPassword)
	http.HandleFunc("/SendPasswordResetEmail/", registration.SendPasswordResetEmail)
	http.HandleFunc("/ResendVerificationEmail/", registration.ResendVerificationEmail)

	//SEARCH
	http.HandleFunc("/search", search.Page)
	http.HandleFunc("/AjaxLoadNext/", search.AjaxLoadNext)

	//GENERIC LOAD COMMENTS
	http.HandleFunc("/AjaxLoadComments/", profile.AjaxLoadComments)
	http.HandleFunc("/AjaxLoadEntryArr/", profile.AjaxLoadEntries)

	//EVENT
	http.HandleFunc("/Event/", event.ViewEvent)
	http.HandleFunc("/AddEvent/", event.AddEvent)
	http.HandleFunc("/StartEvent/", event.StartEvent)
	http.HandleFunc("/ManageEvents/", event.ManageEvents)
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

	//CHAT
	http.HandleFunc("/ch/", chat.Page)
	http.HandleFunc("/ws/", chat.HandleConnections) //weebsocket
	http.HandleFunc("/cN/", chat.HandleChatClients) //websocket
	http.HandleFunc("/AjaxLoadChat/", chat.InitialChat)
	http.HandleFunc("/AjaxLoadMoreChat/", chat.AjaxLoadMoreChat)
	http.HandleFunc("/AjaxChatNotifications/", chat.AjaxNotificationLoad)

	//DASHBOARD
	http.HandleFunc("/dash/", dash.ViewDashboard)
	//http.HandleFunc("/ch/", chat.Page)
	// http.HandleFunc("/ws", chat.Run)

	//Notifications
	http.HandleFunc("/AjaxNotifications/", notification.AjaxNotificationLoad)

	//IMG
	http.HandleFunc("/img/", img.Display)

	http.ListenAndServe(":"+currentPort, nil)
}
