package types

type Comment struct {
	CommentID      string    `json:CommentID`
	PosterID       string    `json:PosterID`
	PosterUsername string    `json:PosterUsername`
	Timestamp      string	 `json:Timestamp`
	Content        string    `json:Text`
}
