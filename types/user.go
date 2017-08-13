package types

import(
	//"fmt"
	//"reflect"
	//"errors"
)

type User struct {
	//_id					string		`json:_id`
	Password			string   	`json:Password` // Maybe we shouldn't keep it in plain text later?
	//Privileges []Privilege 	`json:Privileges`
	
	
	FirstName			string      `json:FirstName`
	LastName			string      `json:LastName`
	Email				string	    `json:Email`
	Location			[]string    `json:Location`
	HighSchool			[]string	`json:HighSchool`
	GradDate			string		`json:GradDate`
	UndergradSchool		string  	`json:UndergradSchool`
	Majors				[]string	`json:Majors`
	Minors				[]string	`json:Minors`
	Class				int8		`json:Class`
	Dob					string		`json:Dob`
	AccCreation 		string	    `json:AccCreation`
	Visible				bool		`json:Visible`
	Status				bool		`json:Status`
	ExpirationDate		string		`json:ExpirationDate`
	Avatar				string		`json:Avatar`
	CroppedAvatar 		string		`json:CropAvatar`
	Banner				string		`json:Banner`
	Organization 		string		`json:Organization`
	Category			string		`json:Category`
	Phone				string		`json:Phone`
	Paid				bool		`json:Paid`
	AccType				int8		`json:AccType`
	Description 		string		`json:Description`
	QuickLinks			string		`json:QuickLinks`
	Tags				string		`json:Tags`
	Projects			[]string	`json:Projects`
	BlockedUsers 		[]string	`json:BlockedUsers`
	ConversationIDs 	[]string	`json:ConversationIDs`
}


