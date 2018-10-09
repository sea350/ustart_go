package backend

import (
	"flag"
	"log"
	"net/http"

	chat "github.com/sea350/ustart_go/backend/middleware/chat"
	dash "github.com/sea350/ustart_go/backend/middleware/dashboard"
	"github.com/sea350/ustart_go/backend/middleware/entry"
	event "github.com/sea350/ustart_go/backend/middleware/event"
	fail "github.com/sea350/ustart_go/backend/middleware/fail"
	follow "github.com/sea350/ustart_go/backend/middleware/follow"
	img "github.com/sea350/ustart_go/backend/middleware/img"
	inbox "github.com/sea350/ustart_go/backend/middleware/inbox"
	login "github.com/sea350/ustart_go/backend/middleware/login"
	"github.com/sea350/ustart_go/backend/middleware/notification"
	profile "github.com/sea350/ustart_go/backend/middleware/profile"
	project "github.com/sea350/ustart_go/backend/middleware/project"
	registration "github.com/sea350/ustart_go/backend/middleware/registration"
	search "github.com/sea350/ustart_go/backend/middleware/search"
	settings "github.com/sea350/ustart_go/backend/middleware/settings"
	widget "github.com/sea350/ustart_go/backend/middleware/widget"
)

// Server is a monolithic service providing access to all of UStart's data
type Server struct {
	port string
}

// New returns a new backend server, given the config object
func New(cfg *Config) *Server {
	return &Server{}
}

// Run starts the backend http server
func (srv *Server) Run() error {
	log.SetPrefix("Backend Server:")
	log.Println("Booting...")

	flag.Parse()

	http.HandleFunc("/404/", fail.Fail)
	http.HandleFunc("/KillUstartPlsNoUserinoCappucinoDeniro", fail.KillSwitch)

	http.HandleFunc("/Inbox/", inbox.Inbox)

	// login/out
	http.HandleFunc("/loginerror/", login.Error)
	http.HandleFunc("/", login.Home)
	http.HandleFunc("/profilelogin/", login.Login)
	http.HandleFunc("/logout/", login.Logout)
	http.HandleFunc("/unverified/", login.Unverified)

	// generic user interactions
	http.HandleFunc("/profile/", profile.ViewProfile)
	http.HandleFunc("/callme/", profile.Follow)
	http.HandleFunc("/Like/", profile.Like)
	http.HandleFunc("/addSkill/", profile.AddTag)
	http.HandleFunc("/deleteSkill/", profile.DeleteTag)
	http.HandleFunc("/addLink/", profile.AddQuickLink)
	http.HandleFunc("/deleteLink/", profile.DeleteQuickLink)
	http.HandleFunc("/followers/", profile.FollowersPage)
	http.HandleFunc("/following/", profile.FollowersPage)
	http.HandleFunc("/toggleProjectInvis/", profile.AjaxChangeProjVisibility)
	http.HandleFunc("/toggleEventInvis/", profile.AjaxChangeEventVisibility)
	http.HandleFunc("/AjaxUserFollowProjectToggle/", follow.AjaxUserFollowsProject)
	http.HandleFunc("/testWall/", profile.TestWallPage)
	http.HandleFunc("/AjaxUserFollowsUser/", follow.AjaxUserFollowsUser)
	http.HandleFunc("/AjaxUserFollowsProject/", follow.AjaxUserFollowsProject)
	//http.HandleFunc("/AjaxUserSuggestions/", profile.AjaxLoadSuggestedUsers)
	//http.HandleFunc("/UserSuggestions/", profile.LoadSuggestedUsers)

	// widgets
	http.HandleFunc("/addWidget/", widget.AddWidget)
	http.HandleFunc("/addProjectWidget/", widget.AddProjectWidget)
	http.HandleFunc("/addEventWidget/", widget.AddEventWidget)
	http.HandleFunc("/deleteWidget/", widget.DeleteWidgetProfile)
	http.HandleFunc("/deleteProjectWidget/", widget.DeleteWidgetProject)
	http.HandleFunc("/deleteEventWidget/", widget.DeleteWidgetEvent)
	http.HandleFunc("/deleteLinkFromWidget/", widget.EditWidgetDataDelete)
	http.HandleFunc("/sortUserWidgets/", widget.SortUserWidgets)

	// projects
	http.HandleFunc("/Projects/", project.ProjectsPage)
	http.HandleFunc("/MyProjects/", project.MyProjects)
	http.HandleFunc("/CreateProjectPage/", project.CreateProjectPage)
	http.HandleFunc("/UpdateProjectTags/", project.UpdateTags)
	http.HandleFunc("/UpdateProjectWantedSkills/", project.UpdateSkills)
	http.HandleFunc("/AddProjectLink/", project.AddQuickLink)
	http.HandleFunc("/DeleteProjectLink/", project.DeleteQuickLink)
	http.HandleFunc("/NewMembers/", project.ManageProjects)
	http.HandleFunc("/LoadJoinRequests/", project.LoadJoinRequests)
	http.HandleFunc("/RequestToJoin/", project.RequestToJoin)
	http.HandleFunc("/AcceptJoinRequest/", project.AcceptJoinRequest)
	http.HandleFunc("/RejectJoinRequest/", project.RejectJoinRequest)
	http.HandleFunc("/AjaxLoadProjectFollowers", project.AjaxLoadProjectFollowers)
	http.HandleFunc("/ProjectFollowers/", project.FollowersPage)
	http.HandleFunc("/DeleteProject/", project.Nuke)

	// settings
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
	http.HandleFunc("/eventTime/", settings.EventTime)
	http.HandleFunc("/eventLocationChang/", settings.EventLocation)
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
	http.HandleFunc("/projectLocChange/", settings.ProjectLocation)
	http.HandleFunc("/projectCategory/", settings.ProjectCategory)
	http.HandleFunc("/projectCustomURL/", settings.ProjectCustomURL)
	http.HandleFunc("/leaveProject/", settings.LeaveProject)
	http.HandleFunc("/projectLogo/", settings.ProjectLogo)
	http.HandleFunc("/changeMemberClass/", settings.ChangeMemberClass)

	// user registration
	http.HandleFunc("/Signup/", registration.Signup)
	http.HandleFunc("/Registration/Type/", registration.RegisterType)
	http.HandleFunc("/registrationcomplete/", registration.Complete)
	http.HandleFunc("/welcome/", registration.Registration)
	http.HandleFunc("/Activation/", registration.EmailVerification)
	http.HandleFunc("/ResetPassword/", registration.ResetPassword)
	http.HandleFunc("/SendPasswordResetEmail/", registration.SendPasswordResetEmail)
	http.HandleFunc("/ResendVerificationEmail/", registration.ResendVerificationEmail)

	// search
	http.HandleFunc("/search", search.Page)
	http.HandleFunc("/AjaxLoadNext/", search.AjaxLoadNext)

	// events
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
	http.HandleFunc("/AjaxNewGuest/", event.NewGuest)
	http.HandleFunc("/FindEventGuest/", event.FindEventMember)
	http.HandleFunc("/FindEventMember/", event.FindEventMember)
	http.HandleFunc("/FindEventProject/", event.FindEventProject)
	http.HandleFunc("/AddEventGuest/", event.AddEventGuestRequest)
	http.HandleFunc("/AddEventMember/", event.AddEventMemberRequest)
	http.HandleFunc("/AddEventProject/", event.AddEventProjectRequest)

	// chat
	http.HandleFunc("/ch/", chat.Page)
	http.HandleFunc("/ws/", chat.HandleConnections) //weebsocket
	http.HandleFunc("/cN/", chat.HandleChatClients) //websocket
	http.HandleFunc("/AjaxLoadChat/", chat.InitialChat)
	http.HandleFunc("/AjaxLoadMoreChat/", chat.AjaxLoadMoreChat)
	http.HandleFunc("/AjaxChatNotifications/", chat.AjaxNotificationLoad)

	// entries
	http.HandleFunc("/AjaxLoadComments/", entry.AjaxLoadComments) //general
	http.HandleFunc("/editPost/", entry.EditEntry)
	http.HandleFunc("/deletePost/", entry.DeleteEntry)
	http.HandleFunc("/addPost/", entry.MakeUserEntry) //user
	http.HandleFunc("/shareEntry/", entry.ShareEntry)
	http.HandleFunc("/ajaxUserEntries/", entry.AjaxLoadUserEntries)
	http.HandleFunc("/getComments/", profile.GetComments)
	http.HandleFunc("/AddComment/", profile.AddComment)
	http.HandleFunc("/AddComment2/", profile.AddComment2)
	http.HandleFunc("/ProjectMakeEntry/", entry.MakeProjectEntry) //project
	http.HandleFunc("/AjaxLoadProjectEntries/", entry.AjaxLoadProjectEntries)
	http.HandleFunc("/EventMakeEntry/", event.MakeEventEntry) //event
	http.HandleFunc("/AjaxLoadEventEntries/", event.AjaxLoadEventEntries)

	// dashboard
	http.HandleFunc("/dash/", dash.ViewDashboard)
	http.HandleFunc("/AjaxDash/", dash.AjaxLoadDashEntries)

	// notifs
	http.HandleFunc("/AjaxNotifications/", notification.AjaxNotificationLoad)
	http.HandleFunc("/AjaxRemoveNotification/", notification.RemoveNotification)
	http.HandleFunc("/AjaxMarkAsSeen/", notification.MarkAsSeen)
	http.HandleFunc("/AjaxScrollNotifications/", notification.AjaxScrollNotification)
	http.HandleFunc("/Notifications/", notification.Page)

	// images
	http.HandleFunc("/img/", img.Display)

	log.Printf("Listening on %s\n", srv.port)
	return http.ListenAndServe(":"+srv.port, nil)
}
