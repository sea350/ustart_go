package types

type Widget struct {
	UserID         string `json:UserID`
	Title          string `json:Title`
	Description    string `json:Description`
	Position       int    `json:Position`
	Classification int    `json:Classification`
}
