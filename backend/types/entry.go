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
	ReferenceEntry string `json:"ReferenceEntry"` //Refers to parent entry's ID
	ReferenceID    string `json:"ReferenceID"`    //ID of project or event

	MediaRef  string    `json:"MediaRef"`
	TimeStamp time.Time `json:"TimeStamp"`
	Likes     []Like    `json:"Likes"`
	ShareIDs  []string  `json:"ShareIDs"`
	ReplyIDs  []string  `json:"ReplyIDs"`
	Visible   bool      `json:"Visible"`
}

//UserOriginalEntry ... converts an entry into one configured for an original user post
func (entry *Entry) UserOriginalEntry(posterID string, content string) {

	entry.PosterID = posterID
	entry.Classification = 0
	entry.Content = []rune(content)
	entry.TimeStamp = time.Now()
	entry.Visible = true
}

//UserShareEntry ... converts an entry into one configured for an original share post
func (entry *Entry) UserShareEntry(posterID string, originalEntryID string, content string) {

	entry.PosterID = posterID
	entry.Content = []rune(content)
	entry.ReferenceEntry = originalEntryID
	entry.TimeStamp = time.Now()
	entry.Classification = 2
	entry.Visible = true
}

//ProjectOriginalEntry ... converts an entry into one configured for an original project post
func (entry *Entry) ProjectOriginalEntry(posterID string, projectID string, content string) {

	entry.PosterID = posterID
	entry.Classification = 3
	entry.Content = []rune(content)
	entry.TimeStamp = time.Now()
	entry.Visible = true
	entry.ReferenceID = projectID
}
