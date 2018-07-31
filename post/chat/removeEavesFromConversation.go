package post

import (
	"context"
	"log"

	get "github.com/sea350/ustart_go/get/chat"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//RemoveEavesFromConversation ...removes eaves from convo and also syncronizes proxies
func RemoveEavesFromConversation(eclient *elastic.Client, conversationID string, eavesID string) error {

	AppendToProxyLock.Lock()
	defer AppendToProxyLock.Unlock()

	ctx := context.Background()

	convo, err := get.ConvoByID(eclient, conversationID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return err
	}

	for i := range convo.Eavesdroppers {
		if convo.Eavesdroppers[i].DocID == eavesID {
			if i < len(convo.Eavesdroppers)-1 {
				convo.Eavesdroppers = append(convo.Eavesdroppers[:i], convo.Eavesdroppers[i+1:]...)
			} else {
				convo.Eavesdroppers = convo.Eavesdroppers[:i]
			}
		}
	}

	proxyID, err := get.ProxyIDByUserID(eclient, eavesID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return err
	}

	proxy, err := get.ProxyMsgByID(eclient, proxyID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return err
	}

	for i := range proxy.Conversations {
		if proxy.Conversations[i].ConvoID == conversationID {
			proxy.Conversations = append(proxy.Conversations[:i], proxy.Conversations[i+1:]...)
			break
		}
	}

	_, err = eclient.Update().
		Index(globals.ProxyMsgIndex).
		Type(globals.ProxyMsgType).
		Id(proxyID).
		Doc(map[string]interface{}{"Conversations": proxy.Conversations}).
		Do(ctx)

	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return err
	}

	_, err = eclient.Update().
		Index(globals.ConvoIndex).
		Type(globals.ConvoType).
		Id(conversationID).
		Doc(map[string]interface{}{"Eavesdroppers": convo.Eavesdroppers}).
		Do(ctx)

	return err
}
