package settings

type form struct {
	//Bio fields:
	Username  string `json:"Username"`
	NewUName  string `json:"NewUName"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Intent    string `json:"Intent"` //see below

	//account fields:
	Email       string `json:"Email"`
	Password    string `json:"Password"`
	NewPassword string `json:"NewPassword"`

	//pic fields:
	Avatar string `json:"Avatar"`
	Banner string `json:"Banner"`

	// Fields related to the requestor
	Token string `json:"Token"`
	//User  string `json:"User"`
}

/*Intents:
Change:
	cu = change username
	cn = change name (first/last)
	ce = change email
	cp = change password
	ca = change avatar
	cb = change banner
Get:
	gu = get username
	gn = get name (first/last)
	ge = get email

	ga = get avatar
	gb = get banner
*/