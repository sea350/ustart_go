package types

//Follow struct ...
//Stores maps for followers and following
type Follow struct {
	DocID     string          `json:"DocID"`
	Followers map[string]bool `json:"Followers"`
	Following map[string]bool `json:"Following"`
	Bell      []string        `json:"Bell"`
}
