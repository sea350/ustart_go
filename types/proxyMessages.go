package types

import "time"

//ProxyMessages ... an ES indexed array of conversations designed to offload upload demand from user
type ProxyMessages struct {
	DocID         string                       `json:"DocID"`
	Class         int                          `json:"Class"`
	Conversations map[string]ConversationState `json:"Conversations"` //must be ordered by most recent interaction
	//Class 1 = User inbox
	//Class 2 = Project Inbox
}

//ConversationState ... In addition to providing the ID of the conversation, this caches the last message and how many unread messages there are
type ConversationState struct {
	// NumUnread   int       `json:"NumUnread"`
	// LastMessage Message   `json:"LastMessage"`
	Read        bool      `json:"Read"`
	Muted       bool      `json:"Muted"`
	MuteTimeout time.Time `json:"MuteTimeout"`
}
