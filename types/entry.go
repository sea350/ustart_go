package types

import(
	"time"


)

type Like struct{

	UserId				string				`json:UserId`
	TimeStamp			time.Time			`json:TimeStamp`

}

type Share struct{

	UserId				string				`json:UserId`
	TimeStamp			time.Time			`json:TimeStamp`

}

type Entry struct{

	PosterId			string				`json:PosterId`
	Classification		int8				`json:Classification`
	Content				[]rune				`json:Content`
	ReferenceEntry		string				`json:RefrenceEntry`
	MediaRef			string				`json:MediaRef`
	TimeStamp			time.Time			`json:TimeStamp`
	Likes				[]Like				`json:Likes`
	Shares				[]Share				`json:Shares`
	ReplyIDs			[]string			`json:ReplyIDs`
	Visible				bool				`json:Visible`


}