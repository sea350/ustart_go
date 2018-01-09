package post

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart_go/types"
	globals "github.com/sea350/ustart_go/globals"
	"context"
	"errors"

)

func ReindexChat(eclient *elastic.Client, newChat types.Chat, chatID string)error{
	ctx:=context.Background()

	exists, err := eclient.IndexExists(globals.ChatIndex).Do(ctx)

	if err != nil {return err}

	if !exists {return errors.New("Index does not exist")}

	_, err = eclient.Index().
		Index(globals.ChatIndex).
		Type(globals.ChatType).
		Id(chatID).
		BodyJson(newChat).
		Do(ctx)

	if err != nil {return err}

	return nil
}