package post

import (
	"context"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//IndexConvo ...
//Indexes a new conversation
func IndexConvo(eclient *elastic.Client, newConvo types.Conversation) (string, error) {
	//ADDS NEW CHAT TO ES RECORDS (requires an elastic client and a Chat type)
	//RETURNS AN error and the new chat's ID IF SUCESSFUL error = nil
	ctx := context.Background()
	var convoID = ""

	idx, Err := eclient.Index().
		Index(globals.ConvoIndex).
		Type(globals.ConvoType).
		BodyJson(newConvo).
		Do(ctx)

	if Err != nil {

		return convoID, Err
	}
	convoID = idx.Id

	return convoID, Err
}
