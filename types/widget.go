package types

type Widget struct {
	UserID         string `json:UserID`
	Link           string `json:Link`
	Position       int    `json:Position`
	Classification string `json:Classification`
}
