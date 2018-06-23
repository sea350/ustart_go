package types

//ProxyMessages ... an ES indexed array of conversations designed to offload upload demand from user
type ProxyMessages struct {
	DocID         string              `json:"DocID"`
	Class         int                 `json:"Class"`
	Conversations []ConversationState `json:"Conversations"` //LIMIT 8, must be ordered by most recent interaction
	ConvoArchive  []string            `json:"ConvoArchive"`  //no limit, also must be organized by most recent
}

//ConversationState ... In addition to providing the ID of the conversation, this caches the last message and how many unread messages there are
type ConversationState struct {
	ConvoDocID  string  `json:"ConvoDocID"`
	NumUnread   int     `json:"NumUnread"`
	LastMessage Message `json:"LastMessage"`
}
