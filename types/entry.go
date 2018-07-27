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
	//class 5 = project share
	//class 6 = event page post
	//class 7 = event comment
	//class 8 = event share
	Content        []rune `json:"Content"`
	ReferenceEntry string `json:"ReferenceEntry"`
	ReferenceID    string `json:"ReferenceID"`

	MediaRef  string    `json:"MediaRef"`
	TimeStamp time.Time `json:"TimeStamp"`
	Likes     []Like    `json:"Likes"`
	ShareIDs  []string  `json:"ShareIDs"`
	ReplyIDs  []string  `json:"ReplyIDs"`
	Visible   bool      `json:"Visible"`
}
