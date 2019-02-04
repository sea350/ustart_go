package types

//Badge ... something that was so easy it could've been written without warnings by a 5th grader but since min wrote it, it had 3 leaving steven to fix it
type Badge struct {
	ID          string   `json:"ID"`
	Type        string   `json:"Type"`
	Roster      []string `json:"Roster"`
	ImageLink   string   `json:"ImageLink"`
	Description string   `json:"Description"`
	Tags        []string `json:"Tags"`
}
