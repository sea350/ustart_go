package types

import (
	"time"
)

//Notification ... a generalized struct designed to store all information paticular to a certain notification
type Notification struct {
	Class        int       `json:"Class"`
	DocID        string    `json:"DocID"`
	RedirectToID string    `json:"RedirectToID"` //if the notification is in reference to something specific ie.post id or project
	ReferenceIDs []string  `json:"ReferenceIDs"`
	Seen         bool      `json:"Seen"`
	Timestamp    time.Time `json:"Timestamp"`
	Invisible    bool      `json:"Invisible"`
}

//LikedEntry ... creates a notification for a new follower
func (notif *Notification) LikedEntry(posterID string, entryID string, likerID string) {
	notif.Class = 1
	notif.DocID = posterID
	notif.RedirectToID = entryID
	notif.ReferenceIDs = append(notif.ReferenceIDs, likerID)
	notif.Timestamp = time.Now()
}

//CommentedEntry ... creates a notification for a new follower
func (notif *Notification) CommentedEntry(posterID string, entryID string, commenterID string) {
	notif.Class = 2
	notif.DocID = posterID
	notif.RedirectToID = entryID
	notif.ReferenceIDs = append(notif.ReferenceIDs, commenterID)
	notif.Timestamp = time.Now()
}

//SharedEntry ... creates a notification for a new follower
func (notif *Notification) SharedEntry(posterID string, entryID string, sharerID string) {
	notif.Class = 3
	notif.DocID = posterID
	notif.RedirectToID = entryID
	notif.ReferenceIDs = append(notif.ReferenceIDs, sharerID)
	notif.Timestamp = time.Now()
}

//NewFollower ... creates a notification for a new follower
func (notif *Notification) NewFollower(followerID string, followingID string) {
	notif.Class = 4
	notif.DocID = followerID
	notif.ReferenceIDs = append(notif.ReferenceIDs, followingID)
	notif.Timestamp = time.Now()
}

//ProjectJoinRequestAccepted ... creates a notification for when a new user is accepted into a project
func (notif *Notification) ProjectJoinRequestAccepted(userID string, projectID string) {
	notif.Class = 12
	notif.DocID = userID
	notif.RedirectToID = projectID
	notif.Timestamp = time.Now()
}

//ProjectJoinRequestDeclined ... creates a notification for when a new user is rejected from a project
func (notif *Notification) ProjectJoinRequestDeclined(userID string, projectID string) {
	notif.Class = 13
	notif.DocID = userID
	notif.RedirectToID = projectID
	notif.Timestamp = time.Now()
}
