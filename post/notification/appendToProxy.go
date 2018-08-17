package post

import (
	"context"
	"errors"
	"log"

	get "github.com/sea350/ustart_go/get/notification"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendToProxy ... appends a new conversation state OR brings a certain conversation state to the back of the list
//needs its own lock for concurrency control
func AppendToProxy(eclient *elastic.Client, proxyID string, notifID string, seen bool) error {

	ctx := context.Background()

	AppendToProxyLock.Lock()
	defer AppendToProxyLock.Unlock()

	exists, err := eclient.IndexExists(globals.ProxyNotifIndex).Do(ctx)
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

	proxy, err := get.ProxyNotificationByID(eclient, proxyID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return err
	}

	var temp string
	for i := len(proxy.NotificationCache) - 1; i >= 0; i-- {
		if proxy.NotificationCache[i] == notifID {
			temp = proxy.NotificationCache[i]
			proxy.NotificationCache = append(proxy.NotificationCache[:i], proxy.NotificationCache[i+1:]...)
			break
		}
	}

	if temp == `` { //adding a new convo
		if len(proxy.NotificationCache) >= 10 {
			proxy.NotificationCache = append(proxy.ConverNotificationCachesations[1:], notifID)
		} else {
			proxy.NotificationCache = append(proxy.NotificationCache, notifID)
		}
	}

	if !seen {
		proxy.NumUnread++
	}

	err = ReindexProxyNotification(eclient, proxyID, proxy)

	return err
}
