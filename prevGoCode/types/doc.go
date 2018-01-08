package types

type Doc struct {
	DocID            string    `json:DocID`
	DocName          string    `json:DocName`
	Timestamp        string    `json:Timestamp`
	UploaderID       string    `json:UploaderID`
	UploaderUsername string    `json:UploaderUsername`
	Text             string    `json:Text`
	File             string    `json:File` // Maybe we should store the actual File here instead of just the filename? We can store it as a []byte or figure out a better way to do so
}
