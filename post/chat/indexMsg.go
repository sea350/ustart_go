package post

import (
	"context"
	"log"
	"os"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//IndexMsg ...
//Indexes a new message
func IndexMsg(eclient *elastic.Client, newMsg types.Message) (string, error) {
	//ADDS NEW CHAT TO ES RECORDS (requires an elastic client and a Chat type)
	//RETURNS AN error and the new chat's ID IF SUCESSFUL error = nil
	ctx := context.Background()
	var msgID string
	exists, err := eclient.IndexExists(globals.MsgType).Do(ctx)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	if !exists {
		_, err := eclient.CreateIndex(globals.MsgIndex).Do(ctx)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}

	}
	idx, Err := eclient.Index().
		Index(globals.MsgIndex).
		Type(globals.MsgType).
		BodyJson(newMsg).
		Do(ctx)

	if Err != nil {
		return msgID, Err
	}
	msgID = idx.Id

	return msgID, nil
}
