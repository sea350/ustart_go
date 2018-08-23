package types

import (
	"time"
)

//Notification ... a generalized struct designed to store all information paticular to a certain notification
type Notification struct {
	Class        int       `json:"Class"`
	DocID        string    `json:"DocID"`
	ReferenceIDs []string  `json:"ReferenceIDs"`
	Seen         bool      `json:"Seen"`
	Timestamp    time.Time `json:"Timestamp"`
	Invisible    bool      `json:"Invisible"`
}
