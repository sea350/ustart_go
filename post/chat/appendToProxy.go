package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/chat"
	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendToProxy ... appends a new conversation state OR brings a certain conversation state to the back of the list
func AppendToProxy(eclient *elastic.Client, proxyID string, conversationID string) error {

	ctx := context.Background()

	proxy, err := get.ProxyMsgByID(eclient, proxyID)
	if err != nil {
		return err
	}

	temp, exists := proxy.Conversations[conversationID]
	if !exists {
		proxy.Conversations[conversationID] = types.ConversationState{}
	} else {
		delete(proxy.Conversations, conversationID)
		proxy.Conversations[conversationID] = temp
	}

	exists, err = eclient.IndexExists(globals.ProxyMsgIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = get.ProxyMsgByID(eclient, proxyID)
	if err != nil {
		return err
	}

	_, err = eclient.Update().
		Index(globals.ProxyMsgIndex).
		Type(globals.ProxyMsgType).
		Id(proxyID).
		Doc(map[string]interface{}{"Conversations": proxy.Conversations}).
		Do(ctx)

	return err
}
