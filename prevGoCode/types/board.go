package types

type Board struct {
	BoardID   string   `json:BoardID`
	BoardName string   `json:BoardName`
	ThreadIDs []string `json:ThreadIDs`
	DocIDs    []string `json:DocIDs`
}