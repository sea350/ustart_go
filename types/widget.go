package types

import "html/template"

type Widget struct {
	ID             string          `json:UserID`
	UserID         string          `json:UserID`
	Data           []template.HTML `json:Data`
	Position       int             `json:Position`
	Classification int             `json:Classification`
	//PLS add data configuration for each classification
	//class 0 = text widget
	//	Data[0] = title
	// 	Data[1] = description
	//class 1 = gallery
	//	Data = full array of images
	//class 2 = calendar
	//	Data = WIP
	//class 3 = poll
	//	Data = WIP
	//class 4 = instagram
	//	Data[] = all urls
	//class 5 = soundcloud
	//	Data[0] = url
	//class 6 = youtube
	//	data[0] = url
	//class 7 = codepen
	//	Data[0] = codepenID
	//class 8 = pinterest
	//	Data[0] = url
	//class 9 = tumblr
	//	Data[0] = tumblrEmbedCode
	//class 10 = spotify
	//	Data[0] = spotifyEmbedCode
	//class 11 = anchor
	//	Data[0] = anchorCode
	//class 12 = medium
	//	Data[0] = username (if nil pay attention to  Data[1:2])
	//	Data[1] = publication (if nil pay attention to  Data[0])
	//	Data[2] = publicationTag (if nil pay attention to  Data[0])
	//	Data[3] = article count
	//class 13 = devianart
	//	Data[0] = username
	//	Data[1] = articlecount
	//class 14 = twitch
	//	Data[0] = username
}
