package User

type user struct {
	
	Password   string   	`json:Password` // Maybe we shouldn't keep it in plain text later?
	//Privileges []Privilege 	`json:Privileges`
	
	
	FirstName			string      `json:FirstName`
	Email				string	    `json:Email`
	Location			[]string    `json:Location`
	HighSchool			[]string	`json:HighSchool`
	GradDate			string		`json:GradDate`
	UndergradSchool		string  `json:UndergradSchool`
	Majors				[]string	`json:Majors`
	Minors				[]string	`json:Minors`
	Class				int8		`json:Class`
	dob					string		`json:DOB`
	AccCreation 		string	    `json:AccCreation`
	Visible				bool		`json:Visible`
	Status				bool		`json:Status`
	ExpirationDate		string	`json:ExpirationDate`
	Avatar				string		`json:Avatar`
	CroppedAvatar 		string	`json:CropAvatar`
	Banner				string		`json:Banner`
	Organization 		string		`json:Organization`
	Category			string		`json:Category`
	Phone				string		`json:Phone`
	Paid				bool		`json:Paid`
	AccType				int8		`json:AccType`
	Description 		string		`json:Description`
	QuickLinks			string		`json:QuickLinks`
	Tags				string		`json:Tags`
	//Projects			[]Project	`json:Projects`
	BlockedUsers 		[]User		`json:BlockedUsers`
	//ConversationIDs []ConversationID	`json:ConversationIDs`
	
}


