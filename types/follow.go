package types

//Follow struct ...
//Stores maps for followers and following
type Follow struct {
	DocID            string          `json:"DocID"`
	UserFollowers    map[string]bool `json:"UserFollowers"` //List of users who are following the entity
	UserFollowing    map[string]bool `json:"UserFollowing"`
	ProjectFollowers map[string]bool `json:"ProjectFollowers"`
	ProjectFollowing map[string]bool `json:"ProjectFollowing"`
	EventFollowers   map[string]bool `json:"EventFollowers"`
	EventFollowing   map[string]bool `json:"EventFollowing"`
	UserBell         map[string]bool `json:"UserBell"`
	ProjectBell      map[string]bool `json:"ProjectBell"`
	EventBell        map[string]bool `json:"EventBell"`
}

//TODO: store ids in lowercase
//Changes must be made in the following ( ;) ):
//signup
//view profile
//user/project(?) type
//app middleware
