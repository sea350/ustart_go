package types

import (
	"time"
)

type Like struct {
	UserID    string    `json:UserID`
	TimeStamp time.Time `json:TimeStamp`
}

type Entry struct {
	PosterID       string `json:PosterID`
	Classification int    `json:Classification`
	//class 0 = user original post
	//class 1 = user reply post
	//class 2 = user share post
	Content        []rune    `json:Content`
	ReferenceEntry string    `json:ReferenceEntry`
	MediaRef       string    `json:MediaRef`
	TimeStamp      time.Time `json:TimeStamp`
	Likes          []Like    `json:Likes`
	ShareIDs       []string  `json:Shares`
	ReplyIDs       []string  `json:ReplyIDs`
	Visible        bool      `json:Visible`
}
