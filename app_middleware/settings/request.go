package settings

type formBio struct {
	// Fields related to the tag
	Username  string `json:"Username"`
	NewUName  string `json:"NewUname"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Intent    string `json:"Intent"` // Can be "ccp" or "gcp" or "del" for change color palette, get color palette and delete

	// Fields related to the requestor
	Token string `json:"Token"`
	//User  string `json:"User"`
}

type formAccount struct {
	// Fields related to the tag
	Username    string `json:"Username"`
	Email       string `json:"Email"`
	Password    string `json:"Password"`
	NewPassword string `json:"NewPassword"`
	Intent      string `json:"Intent"` // Can be "ccp" or "gcp" or "del" for change color palette, get color palette and delete

	// Fields related to the requestor
	Token string `json:"Token"`
	//User  string `json:"User"`
}

type formPic struct {
	// Fields related to the tag
	Username string `json:"Username"`
	Avatar   string `json:"Avatar"`
	Banner   string `json:"Banner"`
	Intent   string `json:"Intent"` // Can be "ccp" or "gcp" or "del" for change color palette, get color palette and delete

	// Fields related to the requestor
	Token string `json:"Token"`
	//User  string `json:"User"`
}
