package types

import (
	"time"
)

//Chat ... its a cat, but french
type Chat struct {
	Name         string    `json:"Name"`
	Image        string    `json:"Image"`
	MonologueIDs []string  `json:"MonologueIDs"`
	TimeCreated  time.Time `json:"TimeCreated"`
}
