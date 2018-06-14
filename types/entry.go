package types

import (
	"time"
)

//Like ... like and like accessories
type Like struct {
	UserID    string    `json:"UserID"`
	TimeStamp time.Time `json:"TimeStamp"`
}

//Entry ... generic entry struct
type Entry struct {
	PosterID       string `json:"PosterID"`
	Classification int    `json:"Classification"`
	//class 0 = user original post
	//class 1 = user reply post
	//class 2 = user share post
	//class 3 = project page post
	//class 4 = project comment
	//class 5 = etc.
	Content        []rune    `json:"Content"`
	ReferenceEntry string    `json:"ReferenceEntry"`
	MediaRef       string    `json:"MediaRef"`
	TimeStamp      time.Time `json:"TimeStamp"`
	Likes          []Like    `json:"Likes"`
	ShareIDs       []string  `json:"ShareIDs"`
	ReplyIDs       []string  `json:"ReplyIDs"`
	Visible        bool      `json:"Visible"`
}
