package post

import (
	"context"
	"log"
	"os"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//IndexProxyMsg ...
//Indexes a new proxy message
func IndexProxyMsg(eclient *elastic.Client, newProxyMsg types.ProxyMessages) (string, error) {
	//ADDS NEW CHAT TO ES RECORDS (requires an elastic client and a Chat type)
	//RETURNS AN error and the new chat's ID IF SUCESSFUL error = nil
	ctx := context.Background()
	var proxyMsgID string
	exists, err := eclient.IndexExists(globals.ProxyMsgIndex).Do(ctx)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	if !exists {
		_, err := eclient.CreateIndex(globals.ProxyMsgIndex).Do(ctx)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}

	}
	idx, Err := eclient.Index().
		Index(globals.ProxyMsgIndex).
		Type(globals.ProxyMsgType).
		BodyJson(newProxyMsg).
		Do(ctx)

	if Err != nil {
		return proxyMsgID, Err
	}
	proxyMsgID = idx.Id

	return proxyMsgID, nil
}
