package types

import "time"

//GuestCode ... Special codes used for signup
type GuestCode struct {
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
