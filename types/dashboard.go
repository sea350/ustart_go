package types

//Dashboard ...
type Dashboard struct {
	ID               string   `json:"ID"`
	Followers        []string `json:"Followers"`
	FollowingProject []string `json:"FollowingProject"`
	FollowingEvent   []string `json:"FollowingEvent"`
}
