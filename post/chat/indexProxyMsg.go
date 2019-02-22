package post

import (
	"context"
	"log"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
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
		log.Println(err)
	}
	if !exists {
		_, err := eclient.CreateIndex(globals.ProxyMsgIndex).Do(ctx)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return proxyMsgID, err
		}
	}
	idx, err := eclient.Index().
		Index(globals.ProxyMsgIndex).
		Type(globals.ProxyMsgType).
		BodyJson(newProxyMsg).
		Do(ctx)

	return idx.Id, err
}
