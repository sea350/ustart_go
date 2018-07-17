package types

import (
	"time"
)

//SignupWarning ...  Security countermeasure for checking amount of signup attempts and locking out IP address for repeated failures (invalid email)
type SignupWarning struct {
	SignLastAttempt      time.Time `json:"SignLastAttempt"`      //Time since the Last Failed Signup Attempt
	SignNumberofAttempts int       `json:"SignNumberofAttempts"` //Number of Failed Signup Attempts
	SignLockoutUntil     time.Time `json:"SignLockoutUntil"`     //Lockout Until User can attempt again
	SignIPAddress        string    `json:"SignIPAddress"`        //IP address of Failed Signup Attempt Offender
	SignLockoutCounter   int       `json:"SignLockoutCounter"`   //Amount of Lockouts the IP address has
	SignDiscovered       bool      `json:"SignDiscovered"`       //Check to see if we have accessed this before or not
}
