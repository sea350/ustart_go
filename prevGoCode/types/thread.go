package types

type Thread struct {
	ThreadName     string    `json:ThreadName`
	ThreadID       string    `json:ThreadID`
	PosterID       string    `json:PosterID`
	PosterUsername string    `json:PosterUsername`
	Content        string    `json:Content`
	Comments       []Comment `json:Comments`
	Timestamp      string    `json:Timestamp`
}
