package get

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart_go/types"
	globals "github.com/sea350/ustart_go/globals"
	"context"
	"encoding/json"
	
)

func GetMonologueByID(eclient *elastic.Client, monoID string)(types.Monologue, error){

	var mono types.Monologue //initialize type mono

	ctx:=context.Background() //intialize context background
	searchResult, err := eclient.Get(). //Get returns doc type, index, etc.
		Index(globals.MonologueIndex).
		Type(globals.MonologueType).
		Id(monoID).
		Do(ctx)
	if (err!=nil){return mono,err}


	Err:= json.Unmarshal(*searchResult.Source, &mono) //unmarshal type RawMessage into user struct
	if (Err!=nil){return mono,Err} //forward error

	return mono, Err

}