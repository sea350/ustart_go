package types

//WARNING: NOT FOR DATABASE USE

//SessionUser ... All data needed to be stored in session
type SessionUser struct {
	FirstName string      `json:FirstName`
	LastName  string      `json:LastName`
	Username  string      `json:Username`
	Email     string      `json:Email`
	DocID     string      `json:ID`
	Interface interface{} `json:Interface`
}

//FloatingHead ... All data needed for a generic user icon
type FloatingHead struct {
	Username string `json:Username`
	//for projects Username = project URLName
	FirstName string `json:FirstName`
	//for projects Firstname = project Name
	LastName string `json:LastName`
	Image    string `json:Image`
	Followed bool   `json:Followed`
	Bio      []rune `json:Bio`
	//for projects Bio = project Description
	DocID          string `json:DocID`
	Classification int    `json:Classification`
}

//JournalEntry ... All data needed to display an entry
type JournalEntry struct {
	ElementID        string      `json:ElementID`
	FirstName        string      `json:FirstName`
	LastName         string      `json:LastName`
	Image            string      `json:Image`
	Element          Entry       `json:Element`
	NumReplies       int         `json:NumReplies`
	NumLikes         int         `json:NumLikes`
	NumShares        int         `json:NumShares`
	ReferenceElement interface{} `json:ReferenceElement`
}

//ProjectSubWidgets ... data specific to each project being displayed on the projects widget
type ProjectSubWidgets struct {
	Name   string `json:Name`
	Link   string `json:Link`
	Avatar string `json:Avatar`
	ID     string `json:ID`
}

//ProjectAggregate ... a compiled struct of all relevant project data
type ProjectAggregate struct {
	DocID          string         `json:DocID`
	ProjectData    Project        `json:ProjectData`
	MemberData     []FloatingHead `json:MemberData`
	Editable       bool           `json:Editable`
	RequestAllowed bool           `json:RequestAllowed`
}
