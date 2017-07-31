package types

type Member struct{
	MemberID 		string		`json:MemberID`
	Role			int8		`json:Role`
	JoinDate		string		`json:JoinDate`
	Title			string		`json:Title`
	Visible			bool		`json:Visible`

}


type Project struct {

	Name				string      `json:Name`
	Members				[]Member	`json:Members`
	Location			[]string    `json:Location`
	ListNeeded			[]string	`json:ListNeeded`
	CreationDate		string		`json:CreationDate`
	Visible				bool		`json:Visible`
	Status				bool		`json:Status`
	QuickLinks			[]string	`json:QuickLinks`
	Avatar				string		`json:Avatar`
	CroppedAvatar 		string		`json:CropAvatar`
	Banner				string		`json:Banner`
	Description 		string		`json:Description`
	Tags				string		`json:Tags`
	BlockedUsers 		[]string	`json:BlockedUsers`
	ConversationIDs 	[]string	`json:ConversationIDs`
}


