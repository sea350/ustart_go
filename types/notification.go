package types

import (
	"time"
)

//Notification ... a generalized struct designed to store all information paticular to a certain notification
type Notification struct {
	Class        int       `json:"Class"`
	DocID        string    `json:"DocID"`
	Hyperlink    string    `json:"Hyperlink"`
	ReferenceIDs []string  `json:"ReferenceIDs"`
	Seen         bool      `json:"Seen"`
	Timestamp    time.Time `json:"Timestamp"`
	Visible      bool      `json:"Visible"`
}
