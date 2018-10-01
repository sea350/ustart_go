package types

//Img ... image, class lets you know wether the image is randomly accessable or needs permissions
//ie, ustart used images vs user used images in public domain
type Img struct {
	Image     string `json:"Image"`
	Invisible bool   `json:"Invisible"`
	Class     int    `json:"Class"`
}
