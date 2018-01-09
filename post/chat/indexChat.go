package post

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart_go/types"
	globals "github.com/sea350/ustart_go/globals"
	"context"

)



func IndexChat(eclient *elastic.Client, newChat types.Chat)(string, error) {
	//ADDS NEW CHAT TO ES RECORDS (requires an elastic client and a Chat type)
	//RETURNS AN error and the new chat's ID IF SUCESSFUL error = nil
	ctx := context.Background()
	var chatID string

	idx, Err := eclient.Index().
		Index(globals.ChatIndex).
		Type(globals.ChatType).
		BodyJson(newChat).
		Do(ctx)

	
	if (Err!=nil){return chatID,Err}
	chatID = idx.Id

	return chatID, nil
}