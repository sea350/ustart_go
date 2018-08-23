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
