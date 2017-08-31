package types
//WARNING: NOT FOR DATABASE USE


type SessionUser struct {
	
	FirstName			string			`json:FirstName`
	LastName			string			`json:LastName`
	Email				string			`json:Email`
	DocID				string			`json:ID`
	Interface			interface{}		`json:Interface`

}

type WallEntry struct{

	Element				Entry			`json:Element`
	RepliesArray		[]Entry			`json:RepliesArray`

}