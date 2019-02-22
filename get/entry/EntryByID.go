package get

import (
	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
	//"errors"
	"context"
	"encoding/json"

	globals "github.com/sea350/ustart_go/globals"
)

//EntryByID ...
func EntryByID(eclient *elastic.Client, entryID string) (types.Entry, error) {

	var entry types.Entry //initialize type entry

	ctx := context.Background()         //intialize context background
	searchResult, err := eclient.Get(). //Get returns doc type, index, etc.
						Index(globals.EntryIndex).
						Type(globals.EntryType).
						Id(entryID).
						Do(ctx)
	if err != nil {
		return entry, err
	}

	Err := json.Unmarshal(*searchResult.Source, &entry) //unmarshal type RawMessage into user struct
	if Err != nil {
		return entry, Err
	} //forward error

	return entry, Err

}
