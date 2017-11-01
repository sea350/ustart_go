package types
//WARNING: NOT FOR DATABASE USE


type SessionUser struct {
	
	FirstName			string			`json:FirstName`
	LastName			string			`json:LastName`
	Username			string			`json:Username`
	Email				string			`json:Email`
	DocID				string			`json:ID`
	Interface			interface{}		`json:Interface`

}

type FloatingHead struct{

	UserID				string			`json:FirstName`
	FirstName			string			`json:FirstName`
	LastName			string			`json:LastName`
	Image				string			`json:Image`
	Followed			bool			`json:Followed`

}

type JournalEntry struct{

	ElementID			string			`json:ElementID`
	FirstName			string			`json:FirstName`
	LastName			string			`json:LastName`
	Image				string			`json:Image`
	Element				Entry			`json:Element`
	NumReplies			int				`json:NumReplies`
	NumLikes			int				`json:NumLikes`
	NumShares			int				`json:NumShares`
	ReferenceElement	interface{}		`json:ReferenceElement`

}