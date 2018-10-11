package types

import "time"

//SignUpCode ... Special codes used for signup
type SignUpCode struct {
	Code           string    `json:"Code"`
	Description    string    `json:"Description"`
	NumUses        int       `json:"NumUses"`
	Expiration     time.Time `json:"Expiration"`
	Users          []int     `json:"Users"` //List of user IDs
	Classification int       `json:"Classification"`
	/*
		0: Does not expire
		1: Expires after a certain number of uses
		2: Expires after a certain time
		3: Expires after a certain number of uses and a certain time (whichever comes first)
	*/
}
