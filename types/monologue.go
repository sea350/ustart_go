package types

import (
	"time"
)

type Message struct {
	Timestamp time.Time `json:Timestamp`
	Content   []rune    `json:Content`
}

type Monologue struct {
	SenderID    string    `json:UserID`
	SenderClass int       `SenderClass`
	Messages    []Message `json:Messages`
	TimeAdded   time.Time `json:TimeAdded`
}
