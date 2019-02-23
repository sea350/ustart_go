package post

import (
	"context"
	"errors"
	"log"

	get "github.com/sea350/ustart_go/get/chat"
	globals "github.com/sea350/ustart_go/globals"
	elastic "github.com/olivere/elastic"
)

//MarkAsRead ... marks a convo as read in proxies
//uses append to proxy lock as it modifies proxies
func MarkAsRead(eclient *elastic.Client, usrID string, conversationID string) error {

	ctx := context.Background()

	AppendToProxyLock.Lock()
	defer AppendToProxyLock.Unlock()

	proxyID, err := get.ProxyIDByUserID(eclient, usrID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return err
	}

	exists, err := eclient.IndexExists(globals.ProxyMsgIndex).Do(ctx)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return err
	}
	if !exists {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(errors.New("Index does not exist"))
		return errors.New("Index does not exist")
	}

	proxy, err := get.ProxyMsgByID(eclient, proxyID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return err
	}

	for i := len(proxy.Conversations) - 1; i >= 0; i-- {
		if proxy.Conversations[i].ConvoID == conversationID {
			if !proxy.Conversations[i].Read {
				proxy.NumUnread--
				proxy.Conversations[i].Read = true
			}
			break
		}
	}

	err = ReindexProxyMsg(eclient, proxyID, proxy)
	return err
}
