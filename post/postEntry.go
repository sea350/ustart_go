package post

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart_go/types"
	get "github.com/sea350/ustart_go/get"
	"context"
	"errors"
)

const ENTRY_INDEX="test-entry_data"
const ENTRY_TYPE="ENTRY"

func IndexEntry(eclient *elastic.Client, newEntry types.Entry)(string, error) {
	//ADDS NEW ENTRY TO ES RECORDS (requires an elastic client and a User type)
	//RETURNS AN error IF SUCESSFUL error = nil
	ctx := context.Background()
	var entryID string

	idx, Err := eclient.Index().
		Index(ENTRY_INDEX).
		Type(ENTRY_TYPE).
		BodyJson(newEntry).
		Do(ctx)

	
	if (Err!=nil){return entryID,Err}
	entryID = idx.Id

	return entryID, nil
}




func ReindexEntry(eclient *elastic.Client, oldEntry types.Entry, entryID string)error{
	ctx:=context.Background()

	exists, err := eclient.IndexExists(ENTRY_INDEX).Do(ctx)

	if err != nil {return err}

	if !exists {return errors.New("Index does not exist")}

	_, err = eclient.Index().
		Index(ENTRY_INDEX).
		Type(ENTRY_TYPE).
		Id(entryID).
		BodyJson(oldEntry).
		Do(ctx)

	if err != nil {return err}

	return nil
}

func UpdateEntryContent(eclient *elastic.Client, entryID string, newContent []rune) error{
	ctx:=context.Background()
	//stringified := string(newContent)

	exists, err := eclient.IndexExists(ENTRY_INDEX).Do(ctx)
	if err != nil {return err}
	if !exists {return errors.New("Index does not exist")}
	




	//script := elastic.NewScript("ctx._source.Content = newCont").Params(map[string]interface{}{"newCont": "I AM NEW"})
	message := newContent
	_, err = eclient.Update().Index(ENTRY_INDEX).
		Type(ENTRY_TYPE).
		Id(entryID).
		Doc(map[string]interface{}{"Content": message}).
		Do(ctx)


	
	if err != nil {return err}
	return nil
}


func UpdateEntry(eclient *elastic.Client, entryID string, newContent interface{}, field string) error{
	ctx:=context.Background()
	//stringified := string(newContent)


	exists, err := eclient.IndexExists(ENTRY_INDEX).Do(ctx)
	if err != nil {return err}
	if !exists {return errors.New("Index does not exist")}

	_, err=get.GetEntryById(eclient, entryID)
	if (err!=nil){return err}

	
	//script := elastic.NewScript("ctx._source.Content = newCont").Params(map[string]interface{}{"newCont": "I AM NEW"})
	message := newContent

	_, err = eclient.Update().Index(ENTRY_INDEX).
		Type(ENTRY_TYPE).
		Id(entryID).
		Doc(map[string]interface{}{field: message}).
		Do(ctx)


	
	if err != nil {return err}
	return nil
}





func AppendToEntry(eclient *elastic.Client, entryID string, newContent interface{}, field string) error{
	 ctx:=context.Background()

    exists, err := eclient.IndexExists(ENTRY_INDEX).Do(ctx)
    if err != nil {return err}
    if !exists {return errors.New("Index does not exist")}

    script := elastic.NewScript("ctx._source."+field+"+= newCont").Params(map[string]interface{}{"newCont": newContent})

    _, err = eclient.Update().
    	Index(ENTRY_INDEX).
    	Type(ENTRY_TYPE).
    	Id(entryID).
        Script(script).
        Do(ctx)

    if err != nil {return err}

    return nil
}









/*
func RemoveFromEntry(eclient *elastic.Client, entryID string, field interface{})(error){
	ctx:=context.Background()
	//stringified := string(newContent)


	exists, err := eclient.IndexExists(ENTRY_INDEX).Do(ctx)
	if err != nil {return err}
	if !exists {return errors.New("Index does not exist")}
	
	var theEntry types.Entry

	theEntry, err=get.GetEntryById(eclient, entryID)
	 
	if (err!=nil){return err}

	cmd:="ctx._source."+field+".remove("+field+")"
	key:=theEntry.field
	script:= elastic.NewScript(cmd).Params(map[string]interface{}{field:key}) 
	//message := "Hello"

	_, err = eclient.Update().Index(ENTRY_INDEX).
		Type(ENTRY_TYPE).
		Id(entryID).
		Script(script).
		Do(ctx)


	
	if err != nil {return err}
	return nil
}	
*/

/*func AddLike(eclient *elastic.Client, entryID string, newLike types.Like)error{

	update, err := client.Update().Index(ENTRY_INDEX).Type(ENTRY_TYPE).Id(entryID).
		Script("ctx._source.Likes += num").
		ScriptParams(map[string]interface{}{"num": 1}).
		Upsert(map[string]interface{}{"retweets": 0}).
		Do()

	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("New version of tweet %q is now %d", update.Id, update.Version)
}*/