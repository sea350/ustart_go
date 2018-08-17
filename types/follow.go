package types

//Follow struct ...
//Stores maps for followers and following
type Follow struct {
	DocID     string          `json:"DocID"`
	Followers map[string]bool `json:"Followers"`
	Following map[string]bool `json:"Following"`
	Bell      map[string]bool `json:"Bell"`
}

//TODO: store ids in lowercase
//Changes must be made in the following ( ;) ):
//signup
//view profile
//user/project(?) type
//app middleware
