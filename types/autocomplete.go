package types

//AutoCompleteObject ... object for autocomplete
type AutoCompleteObject struct {
	Value string `json:"value"`
	Label string `json:"label"`
	Desc  string `json:"desc"`
	Icon  string `json:"icon"`
}
