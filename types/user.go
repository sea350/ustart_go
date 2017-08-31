package types
import(
	"time"
)

type User struct {
	Password			[]byte		`json:Password` // Maybe we shouldn't keep it in plain text later?
	
	
	FirstName			string		`json:FirstName`
	LastName			string		`json:LastName`
	Email				string		`json:Email`
	Username			string  	`json:Username`
	Location			LocStruct	`json:Location`
	HighSchool			[]string	`json:HighSchool`
	GradDate			time.Time	`json:GradDate`
	UndergradSchool		string  	`json:UndergradSchool`
	Majors				[]string	`json:Majors`
	Minors				[]string	`json:Minors`
	Class				int8		`json:Class`
	Dob					time.Time	`json:Dob`
	AccCreation 		time.Time	`json:AccCreation`
	Visible				bool		`json:Visible`
	Status				bool		`json:Status`
	ExpirationDate		time.Time	`json:ExpirationDate`
	Avatar				string		`json:Avatar`
	CroppedAvatar 		string		`json:CropAvatar`
	Banner				string		`json:Banner`
	Organization 		string		`json:Organization`
	Category			string		`json:Category`
	Phone				string		`json:Phone`
	Paid				bool		`json:Paid`
	AccType				int8		`json:AccType`
	Description 		[]rune		`json:Description`
	QuickLinks			[]string	`json:QuickLinks`
	Tags				[]string	`json:Tags`
	LikedPostIds		[]string	`json:LikedPostIds`
	Projects			[]string	`json:Projects`
	BlockedUsers 		[]string	`json:BlockedUsers`
	ConversationIDs 	[]string	`json:ConversationIDs`
	EntryIDs			[]string	`json:EntryIDs`
	
}


