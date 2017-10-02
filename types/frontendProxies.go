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

type JournalEntry struct{

	FirstName			string			`json:FirstName`
	LastName			string			`json:LastName`
	Element				Entry			`json:Element`
	RepliesArray		[]JournalEntry	`json:RepliesArray`
	NumReplies			int				`json:NumReplies`
	NumLikes			int				`json:NumLikes`
	NumShares			int				`json:NumShares`

}