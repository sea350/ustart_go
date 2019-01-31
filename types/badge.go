package types

//Img ... image, class lets you know wether the image is randomly accessible or needs permissions
//ie, ustart used images vs user used images in public domain
type Badge struct {
	Id          string   `json:Id`
	Type        string   `json:"Type"`
	Roster      []string `json:"Roster"`
	ImageLink   string   `json:"ImageLink"`
	Description string   `json:"Description"`
	Tags        []string `json:"Tags"`
}
