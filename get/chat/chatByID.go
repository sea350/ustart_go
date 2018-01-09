package get

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart_go/types"
	globals "github.com/sea350/ustart_go/globals"
	"context"
	"encoding/json"
	
)


func ChatByID(eclient *elastic.Client, chatID string)(types.Chat, error){

	var chat types.Chat //initialize type chat

	ctx:=context.Background() //intialize context background
	searchResult, err := eclient.Get(). //Get returns doc type, index, etc.
		Index(globals.ChatIndex).
		Type(globals.ChatType).
		Id(chatID).
		Do(ctx)
	if (err!=nil){return chat,err}


	Err:= json.Unmarshal(*searchResult.Source, &chat) //unmarshal type RawMessage into user struct
	if (Err!=nil){return chat,Err} //forward error

	return chat, Err

}