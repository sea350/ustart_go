package types

//ProxyNotifications ... an ES indexed array of conversations designed to offload upload demand from user
type ProxyNotifications struct {
	DocID             string   `json:"DocID"`
	NumUnread         int      `json:"NumUnread"`
	NotificationCache []string `json:"NotificationCache"` //must be limited to 5
}

//NotificationSettings ... settings for which notifications a user wants to opt in or out of
type NotificationSettings struct {
	EntryLiked   bool `json:"EntryLiked"`   //Class 1
	EntryReplied bool `json:"EntryReplied"` //Class 2
	EntryShared  bool `json:"EntryShared"`  //Class 3

	NewFollower bool `json:"NewFollower"` //Class 4

	FollowedUserNews    bool `json:"FollowedUserNews"`    //Class 5
	FollowedProjectNews bool `json:"FollowedProjectNews"` //Class 6
	FollowedEventNews   bool `json:"FollowedEventNews"`   //Class 7

	BellUserNews    bool `json:"BellUserNews"`    //Class 8
	BellProjectNews bool `json:"BellProjectNews"` //Class 9
	BellEventNews   bool `json:"BellEventNews"`   //Class 10

	ProjectJoinRequestReceived bool `json:"ProjectJoinRequestReceived"` //Class 11: meant for project managing side
	ProjectJoinRequestAccepted bool `json:"ProjectJoinRequestAccepted"` //Class 12
	ProjectJoinRequestDeclined bool `json:"ProjectJoinRequestDeclined"` //Class 13

	EventRSVPReceived bool `json:"ProjectRSVPReceived"` //Class 14: meant for project managing side
	EventRSVPAccepted bool `json:"ProjectRSVPAccepted"` //Class 15
	EventRSVPDeclined bool `json:"ProjectRSVPDeclined"` //Class 16

}

//Default ... default notification settings configuration
func (notif NotificationSettings) Default() {
	notif.EntryLiked = true
	notif.EntryReplied = true
	notif.EntryShared = true

	notif.NewFollower = true

	notif.FollowedUserNews = false
	notif.FollowedProjectNews = false
	notif.FollowedEventNews = false

	notif.BellUserNews = true
	notif.BellProjectNews = true
	notif.BellEventNews = true

	notif.ProjectJoinRequestReceived = true
	notif.ProjectJoinRequestAccepted = true
	notif.ProjectJoinRequestDeclined = true

	notif.EventRSVPReceived = true
	notif.EventRSVPAccepted = true
	notif.EventRSVPDeclined = true
}
