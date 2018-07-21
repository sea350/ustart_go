package types

import (
	"time"
)

//Message ... we gonna skate to one song and one song only
type Message struct {
	SenderID       string    `json:"SenderID"`
	ConversationID string    `json:"ConversationID"`
	TimeStamp      time.Time `json:"TimeStamp"`
	Content        string    `json:"Content"`
	Hidden         bool      `json:"Hidden"`
}

//Eavesdropper ... Information about a single person in the conversation
type Eavesdropper struct {
	Class     int    `json:"Class"`
	DocID     string `json:"DocID"`
	Bookmark  int    `json:"Bookmark"`  //index of the last message the person saw in the ARCHIVE
	TextColor string `json:"TextColor"` //hex format color
	Nickname  string `json:"Nickname"`
	//Class 1 = user
	//Class 2 = project
}

//Conversation ... an ES indexed structure that is a full record of the entire conversation including a cache of the most recent
type Conversation struct {
	Title            string         `json:"Title"`
	ReferenceID      string         `json:"ReferenceID"`
	Class            int            `json:"Class"`
	Eavesdroppers    []Eavesdropper `json:"Eavesdroppers"`
	MessageIDArchive []string       `json:"MessageArchive"` //no limit but must be ordered by most recent interaction
	MessageIDCache   []string       `json:"MessageCache"`   //LIMIT 100, must be ordered by most recent interaction
	PinnedMessages   []string       `json:"PinnedMessages"`
	//Class 1 = DM
	//Class 2 = groupchat
	//Class 3 = project (see RefrenceID)
	//Class 4 = event (see RefrenceID)
}
