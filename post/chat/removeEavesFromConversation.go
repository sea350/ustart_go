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

	ctx := context.Background()

	convo, err := get.ConvoByID(eclient, conversationID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return err
	}

	delete(convo.Eavesdroppers, eavesID)

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

	delete(proxy.Conversations, conversationID)

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
