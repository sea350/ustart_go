package post

import (
	"context"
	"log"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//IndexProxyNotification ...
//Indexes a new proxy notification
func IndexProxyNotification(eclient *elastic.Client, newProxy types.ProxyNotifications) (string, error) {
	ctx := context.Background()
	var proxyID string
	exists, err := eclient.IndexExists(globals.ProxyNotifIndex).Do(ctx)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	if !exists {
		_, err := eclient.CreateIndex(globals.ProxyNotifIndex).Do(ctx)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return proxyID, err
		}
	}
	idx, err := eclient.Index().
		Index(globals.ProxyNotifIndex).
		Type(globals.ProxyNotifType).
		BodyJson(newProxy).
		Do(ctx)

	if err != nil {
		return proxyID, err
	}
	proxyID = idx.Id

	return proxyID, nil
}
