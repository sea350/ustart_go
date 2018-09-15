package types

//LocStruct .. who knows tbh
type LocStruct struct {
	Street     string `json:"Street"`
	Country    string `json:"Country"`
	CountryVis bool   `json:"CountryVis"`
	State      string `json:"State"`
	StateVis   bool   `json:"StateVis"`
	Lng        string `json:"Lng"`
	Lat        string `json:"Lat"`
	City       string `json:"City"`
	CityVis    bool   `json:"CityVis"`
	Zip        string `json:"Zip"`
	ZipVis     bool   `json:"ZipVis"`
	// Street     string `json:"Street"`
	StreetVis bool `json:"StreetVis"`
}
