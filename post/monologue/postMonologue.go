package post

import(
	elastic "gopkg.in/olivere/elastic.v5"
	"github.com/sea350/ustart_go/types"
	get "github.com/sea350/ustart_go/get"
	"context"
	"errors"
	//"fmt"

)

const globals.MonologueIndex = "test-mono_data"
const globals.MonologueType  = "MONOLOGUE"

func IndexMonologue(eclient *elastic.Client, newMonologue types.Monologue)(string, error) {
	//ADDS NEW MOLOGUE TO ES RECORDS (requires an elastic client and a Monologue type)
	//RETURNS AN error and the new mono's ID IF SUCESSFUL error = nil
	ctx := context.Background()
	var monoID string

	idx, Err := eclient.Index().
		Index(globals.MonologueIndex).
		Type(globals.MonologueType).
		BodyJson(newMonologue).
		Do(ctx)

	
	if (Err!=nil){return monoID,Err}
	monoID = idx.Id

	return monoID, nil
}

func ReindexMonologue(eclient *elastic.Client, newMonologue types.Monologue, monoID string)error{
	ctx:=context.Background()

	exists, err := eclient.IndexExists(globals.MonologueIndex).Do(ctx)

	if err != nil {return err}

	if !exists {return errors.New("Index does not exist")}

	_, err = eclient.Index().
		Index(globals.MonologueIndex).
		Type(globals.MonologueType).
		Id(monoID).
		BodyJson(newMonologue).
		Do(ctx)

	if err != nil {return err}

	return nil
}

func UpdateMonologue(eclient *elastic.Client, monoID string, field string, newContent interface{}) error{
	ctx:=context.Background()

	exists, err := eclient.IndexExists(globals.MonologueIndex).Do(ctx)
	if err != nil {return err}
	if !exists {return errors.New("Index does not exist")}

	_, err=get.GetEntryByID(eclient, monoID)
	if (err!=nil){return err}

	_, err = eclient.Update().Index(globals.MonologueIndex).
		Type(globals.MonologueType).
		Id(monoID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)

	if err != nil {return err}
	
	return nil
}