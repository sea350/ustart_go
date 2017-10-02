package types
import(
	"time"
)

type Link struct{
	URL					string		`json:URL`
	Name				string		`json:Name`
}


type ProjectInfo struct{
	ProjectID			string		`ProjectID`
	Visible				bool		`Visible`	 
}


type User struct {
	Password			[]byte		`json:Password` // Maybe we shouldn't keep it in plain text later?
	
	
	FirstName			string		`json:FirstName`
	LastName			string		`json:LastName`
	Email				string		`json:Email`
	EmailVis			bool		`json:EmailVis`
	Gender				string		`json:Gender`
	GenderVis			bool		`json:GenderVis`
	Phone				string		`json:Phone`
	PhoneVis			bool		`json:PhoneVis`
	Description 		[]rune		`json:Description`
	Username			string  	`json:Username`
	Location			LocStruct	`json:Location`
	HighSchool			string		`json:HighSchool`
	HSGradDate			string		`json:GradDate`
	CollegeGradDate		string		`json:GradDate`
	University			string  	`json:UndergradSchool`
	Majors				[]string	`json:Majors`
	Minors				[]string	`json:Minors`
	Class				int			`json:Class` //freshman:0,sophomore:1...
	Dob					time.Time	`json:Dob`
	AccCreation 		time.Time	`json:AccCreation`
	Visible				bool		`json:Visible`
	Status				bool		`json:Status`
	ExpirationDate		time.Time	`json:ExpirationDate`
	Avatar				string		`json:Avatar`
	CroppedAvatar		string		`json:CropAvatar`
	Banner				string		`json:Banner`
	Organization 		string		`json:Organization`
	Category			string		`json:Category`
	Paid				bool		`json:Paid`
	AccType				int			`json:AccType`//highschool, college etc
	QuickLinks			[]Link		`json:QuickLinks`
	Tags				[]string	`json:Tags`
	SearchHist			[]string	`json:SearchHist`
	LikedEntryIDs		[]string	`json:LikedEntryIds`
	Projects			[]ProjectInfo`json:Projects`
	BlockedUsers 		[]string	`json:BlockedUsers`
	BlockedBy			[]string	`json:BlockedBy`
	ConversationIDs 	[]string	`json:ConversationIDs`
	EntryIDs			[]string	`json:EntryIDs`
	Following			[]string	`json:Following`
	Followers			[]string	`json:Followers`
	Colleagues			[]string	`json:Colleagues`
	SentCollReq			[]string	`json:SentCollReq`
	ReceivedCollReq		[]string	`json:ReceivedCollReq`
	SentProjReq			[]string	`json:SentProjReq`
	ReceivedProjReq		[]string	`json:ReceivedProjReq`
	FirstLogin			bool		`json:FirstLogin`
	
}


