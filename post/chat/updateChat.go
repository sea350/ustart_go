package post

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart_go/types"
	get "github.com/sea350/ustart_go/get"
	globals "github.com/sea350/ustart_go/globals"
	"context"
	"errors"

)

func UpdateChat(eclient *elastic.Client, chatID string, field string, newContent interface{}) error{
	ctx:=context.Background()

	exists, err := eclient.IndexExists(globals.ChatIndex).Do(ctx)
	if err != nil {return err}
	if !exists {return errors.New("Index does not exist")}

	_, err=get.ChatByID(eclient, chatID)
	if (err!=nil){return err}

	_, err = eclient.Update().
		Index(globals.ChatIndex).
		Type(globals.ChatType).
		Id(chatID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)

	if err != nil {return err}
	
	return nil
}