package types

type Widget struct {
	UserID         string   `json:UserID`
	Data           []string `json:Title`
	Position       int      `json:Position`
	Classification int      `json:Classification`
	//PLS add data configuration for each classification
	//class 0 = text widget
	//	Data[0] = title
	// 	Data[1] = description

}
