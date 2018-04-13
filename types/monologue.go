package types

import (
	"time"
)

//Message ... we gonna skate to one song and one song only
type Message struct {
	Timestamp time.Time `json:"Timestamp"`
	Content   []rune    `json:"Content"`
}

//Monologue ... a single person's contribution in a chat
type Monologue struct {
	SenderID    string    `json:"UserID"`
	SenderClass int       `json:"SenderClass"`
	Messages    []Message `json:"Messages"`
	TimeAdded   time.Time `json:"TimeAdded"`
}
