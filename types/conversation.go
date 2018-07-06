package types

import (
	"time"
)

//Message ... we gonna skate to one song and one song only
type Message struct {
	SenderID  string    `json:"SenderID"`
	Timestamp time.Time `json:"Timestamp"`
	Content   string    `json:"Content"`
	Hidden    bool      `json:"Hidden"`
}

//Eavesdropper ... Information about a single person in the conversation
type Eavesdropper struct {
	DocID    string `json:"DocID"`
	Class    int    `json:"Class"`
	Bookmark int    `json:"Bookmark"` //index of the last message the person saw in the ARCHIVE
	Color    string `json:"Color"`    //hex format color
}

//Conversation ... an ES indexed structure that is a full record of the entire conversation including a cache of the most recent
type Conversation struct {
	Eavesdroppers  map[string]Eavesdropper `json:"Eavesdroppers"`
	MessageArchive []Message               `json:"MessageArchive"` //no limit but must be ordered by most recent interaction
	MessageCache   []Message               `json:"MessageCache"`   //LIMIT 100, must be ordered by most recent interaction
}
