package get

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart_go/types"
	//"errors"
	"context"
	"encoding/json"
	
)

const ENTRY_INDEX="test-entry_data"
const ENTRY_TYPE="ENTRY"

func GetEntryByID(eclient *elastic.Client, entryID string)(types.Entry, error){

	var entry types.Entry //initialize type entry

	ctx:=context.Background() //intialize context background
	searchResult, err := eclient.Get(). //Get returns doc type, index, etc.
		Index(ENTRY_INDEX).
		Type(ENTRY_TYPE).
		Id(entryID).
		Do(ctx)
	if (err!=nil){return entry,err}


	Err:= json.Unmarshal(*searchResult.Source, &entry) //unmarshal type RawMessage into user struct
	if (Err!=nil){return entry,Err} //forward error

	return entry, Err

}