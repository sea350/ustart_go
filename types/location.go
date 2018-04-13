package types

//LocStruct .. who knows tbh
type LocStruct struct {
	Country    string `json:"Country"`
	CountryVis bool   `json:"CountryVis"`
	State      string `json:"State"`
	StateVis   bool   `json:"StateVis"`
	City       string `json:"City"`
	CityVis    bool   `json:"CityVis"`
	Zip        string `json:"Zip"`
	ZipVis     bool   `json:"ZipVis"`
}
